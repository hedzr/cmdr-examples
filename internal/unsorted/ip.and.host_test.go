/*
 */

// Copyright © 2020 Hedzr Yeh.

package tool

import (
	"fmt"
	"github.com/hedzr/logex"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"testing"
)

func TestExternalIPs(t *testing.T) {
	ip, err := externalIP()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ip)
}

func TestLookupHostInfo(t *testing.T) {
	defer logex.CaptureLog(t).Release()
	logrus.SetLevel(logrus.TraceLevel)

	host, _ := os.Hostname()
	{
		addrs, _ := net.LookupIP(host)
		for _, addr := range addrs {
			if ipv4 := addr.To4(); ipv4 != nil {
				fmt.Println("IPv4: ", ipv4)
			}
		}
	}

	fmt.Println("test more...")

	// {
	// 	addrs, _ := net.InterfaceAddrs()
	// 	fmt.Printf("%v\n", addrs)
	// 	for _, addr := range addrs {
	// 		fmt.Printf("IPv4: %v | %v / %v\n", addr, addr.String(), addr.Network())
	// 	}
	// }

	fmt.Println("test more...")

	{
		ifaces, err := net.Interfaces()
		if err != nil {
			return
		}

		// fmt.Printf("%v\n", ifaces)
		for _, iface := range ifaces {
			t.Logf("<i>: %v", iface)
			if iface.Flags&net.FlagUp == 0 {
				continue // interface down
			}
			if iface.Flags&net.FlagLoopback != 0 {
				continue // loopback interface
			}

			var addrs []net.Addr
			addrs, err = iface.Addrs()
			if err != nil {
				return
			}
			for _, addr := range addrs {
				t.Logf("   : %v | %v / %v", addr, addr.String(), addr.Network())
			}
		}
	}

	ip, port, err := LookupHostInfo(false, false)
	t.Logf("%v, %v, %v", ip, port, err)
}
