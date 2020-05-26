/*
 */

// Copyright © 2020 Hedzr Yeh.

package tool

import (
	"errors"
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
)

var priorList = []string{
	"192.168.0.", "192.168.1.", "10.", "172.16.", "172.17.", "172.29.", "172.19.",
}

// IsV4 test ip if it's IPv4
func IsV4(ip net.IP) bool {
	if len(ip) == net.IPv4len {
		return true
	}
	if ip4 := ip.To4(); ip4 != nil {
		return true
	}
	return false
}

// IsV6 test ip if it's IPv6
func IsV6(ip net.IP) bool {
	if len(ip) == net.IPv6len {
		return !IsV4(ip)
	}
	return false
}

// HostToIP would parse string and transform host to ip addresses.
//
// The valid formats are:
//   host[:port],
//   scheme://host[:port]/path,
//
func HostToIP(host string) string {
	if strings.Contains(host, "://") {
		a := strings.Split(host, "://")
		b := strings.Split(a[1], "/")
		if len(b) > 1 {
			c := strings.Split(b[0], ":")
			addrs, e := net.LookupHost(c[0])
			if e != nil {
				err := fmt.Errorf("[WARN] Oops [in HostToIP()]: CAN'T LookupHost(): %v", e)
				logrus.Fatal(err)
			}
			for _, addr := range addrs {
				ip2 := net.ParseIP(addr)
				c[0] = ip2.String()
				b[0] = net.JoinHostPort(c[0], c[1])
				host = fmt.Sprintf("%v://%v", a[0], path.Join(b...))
				break
			}
		}
		return host
	}

	if strings.Contains(host, ":") {
		if h, p, err := net.SplitHostPort(host); err != nil {
			logrus.Fatal(err)
		} else {
			addrs, e := net.LookupHost(h)
			if e != nil {
				err := fmt.Errorf("[WARN] Oops [in HostToIP()]: CAN'T LookupHost(): %v", e)
				logrus.Fatal(err)
			}
			for _, addr := range addrs {
				ip2 := net.ParseIP(addr)
				host = net.JoinHostPort(ip2.String(), p)
				break
			}
		}
		return host
	}

	addrs, e := net.LookupHost(host)
	if e != nil {
		err := fmt.Errorf("[WARN] Oops [in HostToIP()]: CAN'T LookupHost(): %v", e)
		logrus.Fatal(err)
	}
	for _, addr := range addrs {
		ip2 := net.ParseIP(addr)
		host = ip2.String()
		break
	}
	return host
}

func getPriorList() (list []string) {
	list = cmdr.GetStringSliceR("prior-list", priorList...)
	if v, ok := os.LookupEnv("PRIOR_LIST"); ok {
		list = strings.Split(v, ",")
	}
	return
}

func isPrior(ip net.IP) (found bool, ix int) {
	str := ip.String()
	// logrus.Info("prior_list = ", os.Getenv("PRIOR_LIST"))
	var x string
	for ix, x = range getPriorList() {
		if strings.HasPrefix(str, x) {
			found = true
			break
		}
	}
	return
}

func theLanIPs(ipv6Prefer bool) (ips []net.IP, err error) {
	var ifaces []net.Interface
	ifaces, err = net.Interfaces()
	if err != nil {
		return
	}

	var s6, s4 []net.IP

	for _, iface := range ifaces {
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
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			if IsV4(ip) {
				s4 = append(s4, ip)
			} else {
				s6 = append(s6, ip)
			}
		}
	}

	ips = append(ips, s6...)
	ips = append(ips, s4...)

	if len(ips) == 0 {
		err = errors.New("are you connected to the network")
	}
	return
}

func theLanInterface(ipv6Prefer, ipv4Prefer bool) (ifx net.Interface, addrIndex int, ip net.IP, err error) {
	var ifaces []net.Interface
	ifaces, err = net.Interfaces()
	if err != nil {
		return
	}

	list := getPriorList()
	ipsSave := make([]net.IP, len(list)+1)
	ifaceIndices := make([]int, len(list)+1)
	addrIndices := make([]int, len(list)+1)
	for ii := 0; ii < len(ifaceIndices); ii++ {
		ifaceIndices[ii] = -1
		addrIndices[ii] = -1
	}
	logrus.Tracef("    prior-list: %v", list)

	for ii, iface := range ifaces {
		// logrus.Tracef("    # %v - %v", ii, iface)
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

		for iz, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}

			// str := ip.String()
			// logrus.Tracef("      -> %v", ip)

			found, ix := isPrior(ip)
			if found == false {
				ipsSave[len(list)] = ip
				ifaceIndices[len(list)] = ii
				addrIndices[len(list)] = iz
				logrus.Tracef("    - %v - %v | not found for prior list", ip, ii)
			} else {
				if ipsSave[ix] == nil {
					ipsSave[ix] = ip
					ifaceIndices[ix] = ii
					addrIndices[ix] = iz
					logrus.Tracef("    X %v - ii:%v -iz:%v - %v", ip, ii, iz, iface)
				}
				// savedIP = addr
			}
		}
	}

	var addrs []net.Addr
	var savedIP net.IP
	for j, x := range ipsSave {
		if !x.IsUnspecified() && j < len(ipsSave)-1 {
			ii := ifaceIndices[j]
			if ii >= 0 {
				ifx = ifaces[ii]
				addrIndex = addrIndices[j]
				// logrus.Tracef("LAN interface FOUND: j=%v, ii=%v, iz=%v, ifaceIndices=%v, addrIndices=%v, %v", j, ii, addrIndex, ifaceIndices, addrIndices, ifx)
				// return

				addrs, err = ifx.Addrs()
				if err != nil {
					return
				}

				addr := addrs[addrIndex]
				var ipx net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ipx = v.IP
				case *net.IPAddr:
					ipx = v.IP
				}
				if ipx == nil {
					continue
				}

				logrus.Tracef("     . ipx = %v", ipx)
				if savedIP == nil {
					savedIP = ipx
				}
				if IsV4(ipx) {
					if ipv6Prefer {
						continue
					}
					ip = ipx
					if ipv4Prefer {
						logrus.Tracef("    hit: %v", ip)
						return
					}
				} else {
					if ipv4Prefer {
						continue
					}
					ip = ipx
					if ipv6Prefer {
						logrus.Tracef("    hit: %v", ip)
						return
					}
				}

			}
		}
	}

	ip = savedIP
	if ip == nil {
		// logrus.Warnf("  A internet ip address found: '%s'; but we need LAN address; keep searching with findExternalIP().", savedIP.String())
		err = errors.New("not found")
	} else {
		logrus.Tracef("    hit: %v", ip)
	}
	return
}

// LookupHostInfo scans the network adapters on local machine, finds all available IPs
func LookupHostInfo(ipv6Prefer, ipv4Prefer bool) (ip net.IP, port int, err error) {
	fAddr := cmdr.GetStringR("server.rpc_address", os.Getenv("RPC_ADDR"))
	port = cmdr.GetIntR("server.port", DefaultPort)
	if port <= 0 || port > 65535 {
		port = DefaultPort
	}
	if _, port1, err := net.SplitHostPort(fAddr); err == nil {
		port, err = strconv.Atoi(port1)
	}

	if len(fAddr) == 0 {
		// var addrs []net.Addr
		// var savedIP net.IP
		var iface net.Interface
		var idx int

		iface, idx, ip, err = theLanInterface(ipv6Prefer, ipv4Prefer)
		if err == nil {
			logrus.Tracef("    hit: %v | iface = %v | idx = %v", ip, iface, idx)
		}
		// if err == nil {
		// 	addrs, err = iface.Addrs()
		// 	if err != nil {
		// 		return
		// 	}
		//
		// 	logrus.Tracef("LAN interface found: %v, %v addrs. %v", iface, len(addrs), addrs[idx])
		//
		// 	for ij, addr := range addrs {
		// 		logrus.Tracef("  checking addr %v: %v.", ij, addr)
		// 		var ipx net.IP
		// 		switch v := addr.(type) {
		// 		case *net.IPNet:
		// 			ipx = v.IP
		// 		case *net.IPAddr:
		// 			ipx = v.IP
		// 		}
		// 		if ipx == nil {
		// 			continue
		// 		}
		//
		// 		logrus.Tracef("  checking ip: %v. (4:%t 6:%t)(%v %v)", ipx, ipv4Prefer, ipv6Prefer, IsV4(ipx), IsV6(ipx))
		//
		// 		savedIP = ipx
		// 		if IsV4(ipx) {
		// 			if ipv6Prefer {
		// 				continue
		// 			}
		// 			ip = ipx
		// 			if ipv4Prefer {
		// 				logrus.Tracef("    hit: %v", ip)
		// 				return
		// 			}
		// 		} else {
		// 			if ipv4Prefer {
		// 				continue
		// 			}
		// 			ip = ipx
		// 			if ipv6Prefer {
		// 				logrus.Tracef("    hit: %v", ip)
		// 				return
		// 			}
		// 		}
		// 	}
		// }
		//
		// ip = savedIP
		// if savedIP == nil {
		// 	logrus.Warnf("  A internet ip address found: '%s'; but we need LAN address; keep searching with findExternalIP().", savedIP.String())
		// }
		// logrus.Tracef("    hit: %v", ip)
	} else {
		ip, port, err = hostInfo(fAddr, port)
		logrus.Tracef("    hit: %v", ip)
	}
	return
}

// externalIP try to find the internet public ip address.
// it works properly in aliyun network.
//
// TODO only available to get IPv4 address.
//
// externalIP 尝试获得LAN地址.
// 对于aliyun来说，由于eth0是LAN地址，因此此函数能够正确工作；
// 对于本机多网卡的情况，通常这个函数的结果是正确的；
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

// ThisHostname 返回本机名
func ThisHostname() string {
	name, err := os.Hostname()
	if err != nil {
		logrus.Warnf("WARN: %v", err)
		return "host"
	}
	return name
}

// ThisHost 返回当前服务器的LAN ip，通过本机名进行反向解析
func ThisHost() (ip net.IP) {
	ip = net.IPv4zero
	name, err := os.Hostname()
	if err != nil {
		logrus.Warnf("WARN: %v", err)
		return
	}
	logrus.Infof("detected os hostname: %s", name)
	ip, _, _ = hostInfo(name, 0)
	return
}

// LookupHostInfoOld 依据配置文件的 server.rpc_address 尝试解释正确的rpc地址，通常是IPv4的
func LookupHostInfoOld() (net.IP, int, error) {
	fAddr := cmdr.GetStringR("server.rpc_address", "0.0.0.0")
	fPort := cmdr.GetIntR("server.port", 2301)
	if fPort <= 0 || fPort > 65535 {
		fPort = DefaultPort
	}

	if len(fAddr) == 0 {
		name, err := os.Hostname()
		if err != nil {
			logrus.Warnf("WARN: %v", err)
			return net.IPv4zero, 0, err
		}
		logrus.Infof("detected os hostname: %s", name)
		fAddr = name
	}
	return hostInfo(fAddr, fPort)
}

func lookupHostInfoOld2() (ip net.IP, port int, err error) {
	fAddr := cmdr.GetStringR("server.rpc_address", "0.0.0.0")
	port = cmdr.GetIntR("server.port", 2301)
	if port <= 0 || port > 65535 {
		port = DefaultPort
	}

	if len(fAddr) == 0 {
		addrs, _ := net.InterfaceAddrs()
		ips := make([]net.IP, len(getPriorList())+1)
		var savedIP net.IP
		for _, addr := range addrs {
			// fmt.Println("IPv4/v6: ", addr)
			// addr.Network()
			if ipNet, ok := addr.(*net.IPNet); ok {
				if !ipNet.IP.IsUnspecified() {
					// 排除回环地址，广播地址
					// if ipNet.IP.IsGlobalUnicast() {
					// 	continue
					// }
					// // 排除公网地址
					if isLAN(ipNet.IP) {
						found := false
						for ix, x := range getPriorList() {
							if strings.HasPrefix(ipNet.IP.String(), x) {
								if ips[ix] == nil {
									ips[ix] = ipNet.IP
									found = true
									logrus.Debugf("    X %v", ipNet)
									break
								}
							}
						}
						if found == false {
							ips[len(getPriorList())] = ipNet.IP
							// logrus.Debugf("    - %v | not found for prior list", ipNet)
						}
						// return ipNet.IP, port, nil
					} else {
						savedIP = ipNet.IP
						// logrus.Debugf("    . %v | savedIP", ipNet)
					}
				}
			}
		}

		for _, x := range ips {
			if !x.IsUnspecified() {
				ip = x
				return
			}
		}

		if savedIP != nil {
			logrus.Warnf("  A internet ipaddress found: '%s'; but we need LAN address; keep searching with findExternalIP().", savedIP.String())
		}

		// return findExternalIP(ipOrHost, port)

		return
	}
	return hostInfo(fAddr, port)
}

func hostInfo(ipOrHost string, port int) (net.IP, int, error) {
	// macOS 可能会得到错误的主机名
	if strings.EqualFold("bogon", ipOrHost) {
		return findExternalIP(ipOrHost, port)
	}

	host, port1, err := net.SplitHostPort(ipOrHost)
	if err == nil {
		ipOrHost = host
		port, err = strconv.Atoi(port1)
	}

	ip := net.ParseIP(ipOrHost)
	var savedIP net.IP
	if ip == nil || ip.IsUnspecified() {
		addrs, err := net.LookupHost(ipOrHost)
		if err != nil {
			logrus.Warnf("[WARN] Oops: LookupHost(): %v", err)
			return findExternalIP(ipOrHost, port)
		}

		ips := make([]net.IP, len(getPriorList())+1)
		for _, addr := range addrs {
			ip2 := net.ParseIP(addr)
			if !ip2.IsUnspecified() {
				// 排除回环地址，广播地址
				// if ip2.IsGlobalUnicast() {
				// 	continue
				// }
				// // 排除公网地址
				if isLAN(ip2) {
					found := false
					for ix, x := range getPriorList() {
						if strings.HasPrefix(ip2.String(), x) {
							if ips[ix] == nil {
								ips[ix] = ip2
								found = true
								break
							}
						}
					}
					if found == false {
						ips[len(getPriorList())] = ip2
					}
					// return ip2, port, nil
				} else {
					savedIP = ip2
				}
				logrus.Debugf("      x %v", ip2.String())
			} else {
				// allows ipV4/6 zero
				logrus.Debugf("      x %v", ip2.String())
				return findExternalIP(ip2.String(), port)
			}
		}
		for _, x := range ips {
			if x == nil {
				return hostInfo("localhost", port)
			}
			if !x.IsUnspecified() {
				return x, port, nil
			}
		}
		if savedIP != nil {
			logrus.Warnf("[WARN] an internet ipaddress found: '%s'; but we need LAN address; keep searching with findExternalIP().", savedIP.String())
		}
		return findExternalIP(ipOrHost, port)
	}

	return ip, port, nil
	// return net.IPv4zero, 0, fmt.Errorf("cannot lookup 'server.rpc_address' or 'server.port'. cannot register myself.")
}

func isLAN(ip net.IP) bool {
	if ipv4 := ip.To4(); ipv4 != nil {
		if ipv4[0] == 192 && ipv4[1] == 168 {
			return true
		}
		if ipv4[0] == 172 && ipv4[1] == 16 {
			return true
		}
		if ipv4[0] == 10 {
			return true
		}
	} else {
		// TODO 识别IPv6的LAN地址段
	}
	return false
}

func findExternalIP(ipOrHost string, port int) (net.IP, int, error) {
	ip, err := externalIP()
	if err != nil {
		logrus.Errorf("Oops: findExternalIP(): %v", err)
		return net.IPv4zero, 0, err
	}
	logrus.Infof("      use ip rather than hostname: %s", ip)
	// } else {
	// 	// NOTE 此分支尚未测试，由于macOS得到bogon时LookupHost() 必然失败，因此此分支应该是多余的
	// 	for _, a := range addrs {
	// 		fmt.Println(a)
	// 	}
	// }
	return net.ParseIP(ip), port, nil
}

const (
	// DefaultPort is unused
	DefaultPort = 6666
)
