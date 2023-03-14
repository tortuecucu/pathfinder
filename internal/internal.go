package internal

import (
	"fmt"
)

func main() {

	fmt.Println("start")

}

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

//}
