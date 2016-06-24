/*
Package iputil implements some useful functions for manipulating IP addresses
*/
package iputil

import (
	"math/rand"
	"net"
)

// SubnetEqualSubnet returns if to IPNets are equal
func SubnetEqualSubnet(net1, net2 *net.IPNet) bool {
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
func SubnetContainsSubnet(supernet, subnet *net.IPNet) bool {
	if supernet.Contains(subnet.IP) {
		n1len, n1bits := supernet.Mask.Size()
		n2len, n2bits := subnet.Mask.Size()
		if n1len <= n2len && n1bits == n2bits {
			return true
		}
	}
	return false
}

// LastAddr returns the last address in an IPNet, usually the broadcast address
func LastAddr(n *net.IPNet) net.IP {
	rip := make([]byte, len(n.IP)) // return ip
	for i := range n.IP {
		rip[i] = n.IP[i] | ^n.Mask[i]
	}
	return rip
}

// FirstAddr returns the first address in an IPNet, usually the network address
func FirstAddr(n *net.IPNet) net.IP {
	rip := make([]byte, len(n.IP)) // return ip
	for i := range n.IP {
		rip[i] = n.IP[i] & n.Mask[i]
	}
	return rip
}

// NetworkID returns an IPNet representing the network, based on an IPNet of any IP in a network
func NetworkID(n *net.IPNet) *net.IPNet {
	return &net.IPNet{IP: FirstAddr(n), Mask: n.Mask}
}

// RandAddr generates a reandom address in an IPNet
func RandAddr(n *net.IPNet) (net.IP, error) {
	randBytes := make([]byte, len(n.IP))
	_, err := rand.Read(randBytes)
	if err != nil {
		return nil, err
	}

	rip := make([]byte, len(n.IP)) // return ip
	for i := range n.IP {
		rip[i] = n.IP[i] | (^n.Mask[i] & randBytes[i])
	}

	return rip, nil
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
