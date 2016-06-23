package iputil

import (
	"math/rand"
	"net"
)

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

func LastAddr(n *net.IPNet) net.IP {
	ip := n.IP.To4()
	if ip != nil {
		return net.IP{
			ip[0] | ^n.Mask[0],
			ip[1] | ^n.Mask[1],
			ip[2] | ^n.Mask[2],
			ip[3] | ^n.Mask[3],
		}
	}

	return net.IP{
		ip[0] | ^n.Mask[0],
		ip[1] | ^n.Mask[1],
		ip[2] | ^n.Mask[2],
		ip[3] | ^n.Mask[3],
		ip[4] | ^n.Mask[4],
		ip[5] | ^n.Mask[5],
		ip[6] | ^n.Mask[6],
		ip[7] | ^n.Mask[7],
		ip[8] | ^n.Mask[8],
		ip[9] | ^n.Mask[9],
		ip[10] | ^n.Mask[10],
		ip[11] | ^n.Mask[11],
		ip[12] | ^n.Mask[12],
		ip[13] | ^n.Mask[13],
		ip[14] | ^n.Mask[14],
		ip[15] | ^n.Mask[15],
	}
}

func FirstAddr(n *net.IPNet) net.IP {
	ip := n.IP.To4()
	if ip != nil {
		return net.IP{
			ip[0] & n.Mask[0],
			ip[1] & n.Mask[1],
			ip[2] & n.Mask[2],
			ip[3] & n.Mask[3],
		}
	}

	return net.IP{
		ip[0] & n.Mask[0],
		ip[1] & n.Mask[1],
		ip[2] & n.Mask[2],
		ip[3] & n.Mask[3],
		ip[4] & n.Mask[4],
		ip[5] & n.Mask[5],
		ip[6] & n.Mask[6],
		ip[7] & n.Mask[7],
		ip[8] & n.Mask[8],
		ip[9] & n.Mask[9],
		ip[10] & n.Mask[10],
		ip[11] & n.Mask[11],
		ip[12] & n.Mask[12],
		ip[13] & n.Mask[13],
		ip[14] & n.Mask[14],
		ip[15] & n.Mask[15],
	}
}

func NetworkID(n *net.IPNet) *net.IPNet {
	return &net.IPNet{IP: FirstAddr(n), Mask: n.Mask}
}

func RandAddr(n *net.IPNet) (net.IP, error) {
	// ip & (mask | random) should generate a random ip
	ip := n.IP.To4()
	if ip != nil {
		rand_bytes := make([]byte, 4)
		_, err := rand.Read(rand_bytes)
		if err != nil {
			return nil, err
		}
		return net.IP{
			ip[0] | (^n.Mask[0] & rand_bytes[0]),
			ip[1] | (^n.Mask[1] & rand_bytes[1]),
			ip[2] | (^n.Mask[2] & rand_bytes[2]),
			ip[3] | (^n.Mask[3] & rand_bytes[3]),
		}, nil
	}

	rand_bytes := make([]byte, 16)
	_, err := rand.Read(rand_bytes)
	if err != nil {
		return nil, err
	}
	return net.IP{
		ip[0] | (^n.Mask[0] & rand_bytes[0]),
		ip[1] | (^n.Mask[1] & rand_bytes[1]),
		ip[2] | (^n.Mask[2] & rand_bytes[2]),
		ip[3] | (^n.Mask[3] & rand_bytes[3]),
		ip[4] | (^n.Mask[4] & rand_bytes[4]),
		ip[5] | (^n.Mask[5] & rand_bytes[5]),
		ip[6] | (^n.Mask[6] & rand_bytes[6]),
		ip[7] | (^n.Mask[7] & rand_bytes[7]),
		ip[8] | (^n.Mask[8] & rand_bytes[8]),
		ip[9] | (^n.Mask[9] & rand_bytes[9]),
		ip[10] | (^n.Mask[10] & rand_bytes[10]),
		ip[11] | (^n.Mask[11] & rand_bytes[11]),
		ip[12] | (^n.Mask[12] & rand_bytes[12]),
		ip[13] | (^n.Mask[13] & rand_bytes[13]),
		ip[14] | (^n.Mask[14] & rand_bytes[14]),
		ip[15] | (^n.Mask[15] & rand_bytes[15]),
	}, nil
}
