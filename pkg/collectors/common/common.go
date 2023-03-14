package common

import (
	"net"
	"net/http"
	"net/url"
	"os/user"
	"regexp"
	"strings"
	"time"

	"github.com/tortuecucu/pathfinder/pkg/commands"
	"github.com/tortuecucu/pathfinder/pkg/core"
	"golang.org/x/sys/windows/registry"
)

func HttpGet(url string) (*http.Response, error) {
	//use the proxy set in env
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		return resp, nil
	}
}

func getRegistryStringValue(key registry.Key, path string, name string) (string, error) {
	k, err := registry.OpenKey(key, path, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue(name)
	if err != nil {
		return "", err
	}
	return s, nil
}

func NetLookup(name string) (string, error) {
	iprecords, err := net.LookupIP(name)
	if err != nil {
		return "", err
	}
	var retVal string
	for _, ip := range iprecords {
		retVal = ip.String()
	}
	return retVal, nil
}

func ScanPort(protocol string, hostname string, port string) bool {
	var address string
	if port != "" {
		address = hostname + ":" + port
	} else {
		address = hostname
	}

	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}

type ProxyPacUrl struct {
}

func (p ProxyPacUrl) Run(exe *core.FactCollection) {
	const factKey string = "proxy.proxypac.autoconfigurl"

	//getting proxy.pac url from windows registry
	pacUrl, err := getRegistryStringValue(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, "AutoConfigURL")
	if err != nil {
		exe.Facts[factKey] = core.NewFact(p, err, factKey)
	} else {
		exe.Facts[factKey] = core.NewFact(p, pacUrl, factKey)
	}

	//port querying the file server
	u, err := url.Parse(pacUrl)
	if err != nil {
		exe.AddFact("proxy.proxypac.server.parseerror", err, p)
	}
	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		exe.AddFact("proxy.proxypac.server.parseerror", err, p)
	}
	exe.AddFact("proxy.proxypac.server.portquery", ScanPort("tcp", host, port), p)

	//getting .pac with an http request
	resp, err := HttpGet(pacUrl)
	if err != nil {
		exe.AddFact("proxy.proxypac.get", err, p)
	} else {
		exe.AddFact("proxy.proxypac.get", resp.Status, p)
	}

}
func (p ProxyPacUrl) Name() string {
	return "ProxyPacUrl"
}

type OutboundIP struct{}

func (o OutboundIP) Run(exe *core.FactCollection) {
	ip, err := GetOutboundIP()
	if err != nil {
		exe.AddFact("net.outboundip", err, o)
	}
	exe.AddFact("net.outboundip", ip.String(), o)
}
func (o OutboundIP) Name() string {
	return "net.outboundip"
}

type User struct{}

func extractGroups(re *regexp.Regexp, data string) map[string]string {
	result := make(map[string]string)
	if re.MatchString(data) {
		match := re.FindStringSubmatch(data)

		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
	}
	return result
}

func (o User) Run(exe *core.FactCollection) {
	currentUser, err := user.Current()
	if err != nil {
		exe.AddFact("user.error", err, o)
	}

	exe.AddFact("user.username", currentUser.Username, o)
	exe.AddFact("user.name", currentUser.Name, o)

	usernameValues := extractGroups(regexp.MustCompile(`(?m)(?P<domain>\S+)\\(?P<id>\S+)`), currentUser.Username)
	exe.AddFact("user.domain", strings.TrimSpace(usernameValues["domain"]), o)
	exe.AddFact("user.id", strings.TrimSpace(usernameValues["id"]), o)

	nameValues := extractGroups(regexp.MustCompile(`(?P<fullname>[^(]+)\((?P<company>[^)]+)\)`), currentUser.Name)
	exe.AddFact("user.fullname", strings.TrimSpace(nameValues["fullname"]), o)
	exe.AddFact("user.company", strings.TrimSpace(nameValues["company"]), o)

}
func (o User) Name() string {
	return "User"
}

func AddCoreCollectors(plan core.Plan) {
	plan.AddCollector(ProxyPacUrl{})
	plan.AddCollector(OutboundIP{})
	plan.AddCollector(User{})

	plan.AddCollector(commands.NewCommand("systeminfo"))

}
