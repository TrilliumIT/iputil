package iputil

import (
	"net"
	"testing"
)

func TestSubnetEqualSubnetTrue6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80::/64")
	if !SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to equal %v", net1.String(), net2.String())
	}
}

func TestSubnetEqualSubnetFalseDifferentMask6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80::/48")
	if SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

func TestSubnetEqualSubnetFalseDifferentNet6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80:1::/64")
	if SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

func TestSubnetEqualSubnetTrueIpInNet6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	ip2, net2, _ := net.ParseCIDR("fe80::1/64")
	net2.IP = ip2
	if !SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

func TestSubnetContainSubnetTrueEqual6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80::/64")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

func TestSubnetContainSubnetTrueSmaller6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/48")
	_, net2, _ := net.ParseCIDR("fe80::/64")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

func TestSubnetContainSubnetTrueSmallerDiff6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/48")
	_, net2, _ := net.ParseCIDR("fe80:0:0:1::/64")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

func TestSubnetContainSubnetFalseLarger6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80::/48")
	if SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

func TestSubnetContainSubnetFalseDifferent6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80:1::/64")
	if SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

func TestLastAddr6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	lr := net.ParseIP("fe80::ffff:ffff:ffff:ffff")
	if !LastAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", LastAddr(net1).String(), lr.String())
	}
}

func TestFirstAddr6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::55/64")
	lr := net.ParseIP("fe80::")
	if !FirstAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", FirstAddr(net1).String(), lr.String())
	}
}

func TestNetworkID6(t *testing.T) {
	ip1, net1, _ := net.ParseCIDR("fe80::55/64")
	net1.IP = ip1
	_, net2, _ := net.ParseCIDR("fe80::/64")
	nid := NetworkID(net1)
	if !SubnetEqualSubnet(nid, net2) {
		t.Errorf("Expected %v to equal %v", nid.String(), net2.String())
	}
}

func TestRandomAddr6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	for i := 1; i <= 10; i++ {
		rip, err := RandAddr(net1)
		if err != nil {
			t.Errorf("Error generating random bytes")
		}
		if !net1.Contains(rip) {
			t.Errorf("IP %v outside subnet %v", rip.String(), net1.String())
		}
	}
}

func TestIPAdd6(t *testing.T) {
	ip := net.ParseIP("fe80::")
	r := net.ParseIP("fe80::1")
	if !IPAdd(ip, 1).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

func TestIPAddCarryover6(t *testing.T) {
	ip := net.ParseIP("fe80::ffff")
	r := net.ParseIP("fe80::1:4")
	if !IPAdd(ip, 5).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

func TestIPAddCarryover26(t *testing.T) {
	ip := net.ParseIP("fe80::ffff:ffff")
	r := net.ParseIP("fe80::1:0:4")
	if !IPAdd(ip, 5).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

func TestIPSub6(t *testing.T) {
	ip := net.ParseIP("fe80::5")
	r := net.ParseIP("fe80::4")
	if !IPAdd(ip, -1).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

func TestIPSubCarryover6(t *testing.T) {
	ip := net.ParseIP("fe80::100")
	r := net.ParseIP("fe80::fb")
	rip := IPAdd(ip, -5)
	if !rip.Equal(r) {
		t.Errorf("%v should equal %v", rip.String(), r.String())
	}
}

func TestIPSubCarryover26(t *testing.T) {
	ip := net.ParseIP("fe80::1:0")
	r := net.ParseIP("fe80::fffb")
	rip := IPAdd(ip, -5)
	if !rip.Equal(r) {
		t.Errorf("%v should equal %v", rip.String(), r.String())
	}
}
