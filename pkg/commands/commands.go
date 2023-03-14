package commands

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"

	"github.com/go-cmd/cmd"
	"github.com/tortuecucu/pathfinder/pkg/core"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type CommandRun struct {
	Command *Command
	Lines   []string
	Start   time.Time
	End     time.Time
}

type Command struct {
	core.Runnable
	commandName       string
	commandParameters []string
}

func (t Command) Run(facts *core.FactCollection) {
	run := runCommand(&t)
	facts.AddFact(t.commandName, run.Lines, t)
}
func (t Command) Name() string {
	return "command '" + t.commandName + " " + strings.Join(t.commandParameters[:], " ")
}

func NewCommand(name string, parameters ...string) Command {
	return Command{commandName: name, commandParameters: parameters}
}

func runCommand(command *Command) CommandRun {
	var lines []string
	args := []string{"/c", command.commandName}
	args = append(args, command.commandParameters...)
	cmd := cmd.NewCmd("cmd", args...)
	start := time.Now()
	statusChan := cmd.Start()

	go func() { // command timeout
		<-time.After(1 * time.Minute)
		cmd.Stop()
	}()

	finalStatus := <-statusChan // Block waiting for command to exit, be stopped, or be killed

	for _, line := range finalStatus.Stdout {
		//convert line to uft-8 to avoid errors on accented characters
		transformer := transform.NewReader(strings.NewReader(string(line)), charmap.CodePage850.NewDecoder())
		bytes, _ := ioutil.ReadAll(transformer)
		lines = append(lines, string(bytes))
	}
	return CommandRun{Command: command, Lines: lines, Start: start, End: time.Now()}
}

func GetLanguageCode() (string, error) {
	var lang string
	cmd := exec.Command("cmd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	text := strings.ToLower(string(out))
	if strings.Contains(text, "droits") {
		lang = "fr"
	} else if strings.Contains(text, "right") {
		lang = "en"
	} else if strings.Contains(text, "rechte") {
		lang = "de"
	} else if strings.Contains(text, "todos") {
		lang = "es"
	} else {
		lang = ""
	}

	return lang, nil
}

func GetCodepage() (string, error) {
	//TODO: no error handling !
	cmd := cmd.NewCmd("cmd", "/c", "chcp")

	statusChan := cmd.Start()

	go func() { // command timeout
		<-time.After(10 * time.Second)
		cmd.Stop()
	}()

	finalStatus := <-statusChan

	for _, line := range finalStatus.Stdout {
		parts := strings.SplitAfter(line, ":")
		if len(parts) == 2 {
			return strings.TrimSpace(parts[1]), nil
		}
	}
	return "", nil
}

func GetSite() string {
	cmd := NewCommand("nltest", "/dsgetsite")
	result := runCommand(&cmd)
	if len(result.Lines) == 2 {
		return strings.TrimSpace(result.Lines[0])
	} else {
		return ""
	}
}

type LookupResult struct {
	NsServer     string
	NsAddress    string
	MatchName    string
	MatchAddress string
	MatchAliases []string
}

func NsLookup(address string, lang string) (LookupResult, error) {
	cmd := NewCommand("nslookup", address)
	result := runCommand(&cmd)
	match := LookupResult{}
	match.MatchAliases = []string{}

	if len(result.Lines) > 2 {
		match.NsServer = splitLine(result.Lines[0])[1]
		match.NsAddress = splitLine(result.Lines[1])[1]

		if len(result.Lines) > 4 {
			match.MatchName = splitLine(result.Lines[3])[1]
			match.MatchAddress = splitLine(result.Lines[4])[1]
		}

		//add aliases if any
		if len(result.Lines) > 4 {
			for i := 5; i < len(result.Lines); i++ {
				if len(strings.TrimSpace(result.Lines[i])) > 0 {
					match.MatchAliases = append(match.MatchAliases, splitLine(result.Lines[i])[1])
				}
			}
		}

	} else {
		return match, errors.New("nslooup result not parsable")
	}

	return match, nil
}

func normalize(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

func splitLine(line string) []string {
	if strings.Contains(line, ":") {
		parts := strings.Split(line, ":")
		for i, p := range parts {
			parts[i] = normalize(p)
		}
		return parts
	} else {
		return []string{normalize(line)}
	}
}

type PingReply struct {
	Address string
	Bytes   int
	Time    int
	Ttl     int
}

type PingResult struct {
	Name     string
	Start    time.Time
	End      time.Time
	Address  string
	Sent     int
	Received int
	Lost     int
	Minimum  int
	Maximum  int
	Mean     int
	Replies  []PingReply
}

func Ping(address string) PingResult {
	ping := PingResult{}
	//TODO: code it
	return ping
}
