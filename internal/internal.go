package internal

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/go-cmd/cmd"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func main() {

	fmt.Println("start")

	// dirtyCode()

	// iprecords, _ := net.LookupIP("lhsrpimp.corp.ad.aircelle")
	// for _, ip := range iprecords {
	// 	fmt.Println(ip)
	// }

	// //27010@LHSRPLC1.corp.ad.aircelle

	// open := scanPort("tcp", "LHSRPLC1.corp.ad.aircelle", 27010)
	// fmt.Printf("Port Open: %t\n", open)

	// p := proxy.NewProvider("")
	// q := p.GetHTTPProxy("https://rapid7.com")
	// if q != nil {
	// 	fmt.Printf("Found proxy: %s\n", p)
	// } else {
	// 	fmt.Println(q)
	// }

	// //use the proxy set in env
	// resp, err := http.Get("https://insite.collab.group.safran/")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	defer resp.Body.Close()
	// }

	// Start a long-running process, capture stdout and stderr
	command := cmd.NewCmd("cmd", "/c", "systeminfo")
	statusChan := command.Start()

	// command timeout
	go func() {
		<-time.After(1 * time.Minute)
		command.Stop()
	}()

	// Block waiting for command to exit, be stopped, or be killed
	finalStatus := <-statusChan
	var lines []string

	for _, line := range finalStatus.Stdout {

		//cmd chcp display the active code page
		//convert line to uft-8 to avoid errors on accented characters
		transformer := transform.NewReader(strings.NewReader(string(line)), charmap.CodePage850.NewDecoder())
		bytes, _ := ioutil.ReadAll(transformer)
		lines = append(lines, string(bytes))
	}

	fmt.Println(lines)

}

// func scanPort(protocol, hostname string, port int) bool {
// 	address := hostname + ":" + strconv.Itoa(port)
// 	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()
// 	return true
// }

// func dirtyCode() {
// 	fmt.Println("test")
// 	fmt.Println(time.Now())

// 	currentUser, err := user.Current()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	name := currentUser.Name

// 	fmt.Println("username: " + currentUser.Username)

// 	fmt.Println("name is: " + name)

// 	hostname, err := os.Hostname()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("Hostname: " + hostname)
// 	fmt.Println("outbound ip" + GetOutboundIP().String())

// l, err := net.Interfaces()
// if err != nil {
// 	panic(err)

// }
// for _, f := range l {
// 	if f.Flags&net.FlagUp > 0 {
// 		fmt.Printf("%s is up\n", f.Name)
// 	}
// 	fmt.Println(f)
// }

//fmt.Println("=== interfaces ===")

// ifaces, _ := net.Interfaces()
// for _, iface := range ifaces {
// 	fmt.Println("net.Interface:", iface)

// 	addrs, _ := iface.Addrs()
// 	for _, addr := range addrs {
// 		addrStr := addr.String()
// 		fmt.Println("    net.Addr: ", addr.Network(), addrStr)

// 		// Must drop the stuff after the slash in order to convert it to an IP instance
// 		split := strings.Split(addrStr, "/")
// 		addrStr0 := split[0]

// 		// Parse the string to an IP instance
// 		ip := net.ParseIP(addrStr0)
// 		if ip.To4() != nil {
// 			fmt.Println("       ", addrStr0, "is ipv4")
// 		} else {
// 			fmt.Println("       ", addrStr0, "is ipv6")
// 		}
// 		fmt.Println("       ", addrStr0, "is interface-local multicast :", ip.IsInterfaceLocalMulticast())
// 		fmt.Println("       ", addrStr0, "is link-local multicast      :", ip.IsLinkLocalMulticast())
// 		fmt.Println("       ", addrStr0, "is link-local unicast        :", ip.IsLinkLocalUnicast())
// 		fmt.Println("       ", addrStr0, "is global unicast            :", ip.IsGlobalUnicast())
// 		fmt.Println("       ", addrStr0, "is multicast                 :", ip.IsMulticast())
// 		fmt.Println("       ", addrStr0, "is loopback                  :", ip.IsLoopback())
// 	}
// }

// cmd := exec.Command("cmd", "nltest", "/dsgetsite")
// out, err := cmd.CombinedOutput()
// if err != nil {
// 	fmt.Printf("cmd.Run() failed with %s\n", efmt
// }
// fmt.Printf("combined out:\n%s\n", string(out))
//}

// func GetOutboundIP() net.IP {
// 	conn, err := net.Dial("udp", "8.8.8.8:80")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	localAddr := conn.LocalAddr().(*net.UDPAddr)

// 	return localAddr.IP
// }
