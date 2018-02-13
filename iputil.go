/*
Package iputil implements some useful functions for manipulating IP addresses
*/
package iputil

import (
	"math/rand"
	"net"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// SubnetEqualSubnet returns true if to IPNets are equal
// nil is considered to be a global supernet "0.0.0.0/0" or "::/0"
func SubnetEqualSubnet(net1, net2 *net.IPNet) bool {
	if net1 == nil && net2 == nil {
		return true
	}

	if net1 == nil {
		net1 = &net.IPNet{IP: make([]byte, len(net2.IP)), Mask: make([]byte, len(net2.Mask))}
	}

	if net2 == nil {
		net2 = &net.IPNet{IP: make([]byte, len(net1.IP)), Mask: make([]byte, len(net1.Mask))}
	}

	if net1.Contains(net2.IP) {
		n1len, n1bits := net1.Mask.Size()
		n2len, n2bits := net2.Mask.Size()
		if n1len == n2len && n1bits == n2bits {
			return true
		}
	}
	return false
}

// SubnetContainsSubnet returns true if the first subnet contains the second
// nil is considered to be a global supernet "0.0.0.0/0" or "::/0"
func SubnetContainsSubnet(supernet, subnet *net.IPNet) bool {
	if supernet == nil {
		return true
	}

	if subnet == nil {
		subnet = &net.IPNet{IP: make([]byte, len(supernet.IP)), Mask: make([]byte, len(supernet.Mask))}
	}

	if supernet.Contains(subnet.IP) {
		n1len, n1bits := supernet.Mask.Size()
		n2len, n2bits := subnet.Mask.Size()
		if n1len <= n2len && n1bits == n2bits {
			return true
		}
	}
	return false
}

func manipulateAddr(n *net.IPNet, f func(n, m byte) byte) net.IP {
	ml := len(n.Mask) - 1
	il := len(n.IP) - 1
	minl := ml
	if len(n.IP)-1 < minl {
		minl = len(n.IP) - 1
	}
	rip := make([]byte, minl+1) // return ip
	for i := 0; i <= minl; i++ {
		rip[minl-i] = f(n.IP[il-i], n.Mask[ml-i])
	}
	return rip
}

// LastAddr returns the last address in an IPNet, usually the broadcast address
func LastAddr(n *net.IPNet) net.IP {
	return manipulateAddr(n, func(n, m byte) byte { return n | ^m })
}

// FirstAddr returns the first address in an IPNet, usually the network address
func FirstAddr(n *net.IPNet) net.IP {
	return manipulateAddr(n, func(n, m byte) byte { return n & m })
}

// NetworkID returns an IPNet representing the network, based on an IPNet of any IP in a network
func NetworkID(n *net.IPNet) *net.IPNet {
	rn := &net.IPNet{IP: FirstAddr(n), Mask: n.Mask}
	return rn
}

// RandAddr generates a reandom address in an IPNet
func RandAddr(n *net.IPNet) net.IP {
	f := func(n, m byte) byte {
		randBytes := make([]byte, 1)
		rand.Read(randBytes) // rand.Read never returns an err.
		return n | (^m & randBytes[0])
	}
	return manipulateAddr(n, f)
}

//RandAddrWithExclude Generates a random address in an IPNet, excluding the first xf and last xl addresses.
//To generate a random address, excluding the network and broadcast addresses use 1 for xf and xl
func RandAddrWithExclude(n *net.IPNet, xf, xl int) net.IP {
	f := IPAdd(FirstAddr(n), xf)
	l := IPAdd(LastAddr(n), -xl)
	d := IPDiff(l, f)
	if d <= 0 {
		return nil
	}
	return IPAdd(f, rand.Intn(d))
}

//IPDiff returns the difference between ip and ip2
//nil is treated as the zero address
func IPDiff(ip, ip2 net.IP) int {
	ip, ip2 = makeNilZero(ip, ip2)
	ip, ip2 = makeSameLength(ip, ip2)
	o := 1
	if IPBefore(ip, ip2) {
		ip, ip2 = ip2, ip
		o = -1
	}
	ri := 0
	il := len(ip) - 1 // last element in ip
	for i := range ip {
		r := il - i // loop in reverse order
		ri = ri + ((int(ip[r] - ip2[r])) << (uint(8*i) & 0xff))
	}
	return ri * o
}

//IPDiff returns true if ip < ip2
func IPBefore(ip, ip2 net.IP) bool {
	ip, ip2 = makeNilZero(ip, ip2)
	ip, ip2 = makeSameLength(ip, ip2)
	for i := range ip {
		if int(ip[i]) < int(ip2[i]) {
			return true
		}
	}
	return false
}

func makeNilZero(ip, ip2 net.IP) (net.IP, net.IP) {
	if ip == nil {
		ip = net.IP{0, 0, 0, 0}
	}
	if ip2 == nil {
		ip2 = net.IP{0, 0, 0, 0}
	}

	return ip, ip2
}

func makeSameLength(ip, ip2 net.IP) (net.IP, net.IP) {
	if len(ip) != len(ip2) {
		return ip.To16(), ip2.To16()
	}
	/*
		if len(ip) < len(ip2) {
			ip = append(append([]byte{}, ip2[:len(ip2)-len(ip)]...), ip...)
		}
		if len(ip2) < len(ip) {
			ip2 = append(append([]byte{}, ip[:len(ip)-len(ip2)]...), ip2...)
		}
	*/

	return ip, ip2
}

//IPAdd adds an offset to an IP
func IPAdd(ip net.IP, offset int) net.IP {
	rip := make([]byte, len(ip)) // return ip
	os := 1                      // offset sign
	if offset < 0 {
		os = -1
	}
	aos := offset * os // absolute offset
	il := len(ip) - 1  // last element in ip
	var c int          // carryover
	for i := range ip {
		r := il - i                           // loop in reverse order
		ofb := (aos >> uint(8*i) & 0xff) * os // byte offset
		rip[r] = byte(int(ip[r]) + ofb + c)
		if int(rip[r]) != int(ip[r])+c && ((os > 0) == (int(rip[r]) < int(ip[r])+c)) {
			// Indicates that we've wrapped the previous operation, carry
			c = os
		} else {
			c = 0
		}
	}
	return rip
}
