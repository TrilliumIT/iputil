package iputil

import (
	"net"
	"testing"
)

func TestSubnetEqualSubnetTrue(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if !SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to equal %v", net1.String(), net2.String())
	}
}

func TestSubnetEqualSubnetFalseDifferentMask(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.0.0/16")
	_, net2, _ := net.ParseCIDR("10.1.0.0/24")
	if SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

func TestSubnetEqualSubnetFalseDifferentNet(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.0.0/16")
	_, net2, _ := net.ParseCIDR("10.2.0.0/16")
	if SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

func TestSubnetEqualSubnetTrueIpInNet(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.6/16")
	ip2, net2, _ := net.ParseCIDR("10.1.2.1/16")
	net2.IP = ip2
	if !SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

func TestSubnetEqualSubnetTrueBothNil(t *testing.T) {
	if !SubnetEqualSubnet(nil, nil) {
		t.Error("Expected nil subnets to be equal.")
	}
}

func TestSubnetEqualSubnetFalseN1Nil(t *testing.T) {
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if SubnetEqualSubnet(nil, net2) {
		t.Errorf("Expected nil not equal %v", net2)
	}
}

func TestSubnetEqualSubnetFalseN2Nil(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	if SubnetEqualSubnet(net1, nil) {
		t.Errorf("Expected %v not equal nil", net1)
	}
}

func TestSubnetEqualSubnetTrueN1NilN2Global(t *testing.T) {
	_, net2, _ := net.ParseCIDR("0.0.0.0/0")
	if !SubnetEqualSubnet(nil, net2) {
		t.Errorf("Expected nil equal %v", net2)
	}
}

func TestSubnetEqualSubnetTrueN1GlobalN2Nil(t *testing.T) {
	_, net1, _ := net.ParseCIDR("0.0.0.0/0")
	if !SubnetEqualSubnet(net1, nil) {
		t.Errorf("Expected %v equal nil", net1)
	}
}

func TestSubnetContainSubnetTrueBothNil(t *testing.T) {
	if !SubnetContainsSubnet(nil, nil) {
		t.Error("Expected nil contains nil.")
	}
}

func TestSubnetContainSubnetFalseN1Nil(t *testing.T) {
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if !SubnetContainsSubnet(nil, net2) {
		t.Errorf("Expected nil contains %v", net2)
	}
}

func TestSubnetContainSubnetFalseN2Nil(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	if SubnetContainsSubnet(net1, nil) {
		t.Errorf("Expected %v not contains nil", net1)
	}
}

func TestSubnetContainSubnetTrueN1NilN2Global(t *testing.T) {
	_, net2, _ := net.ParseCIDR("0.0.0.0/0")
	if !SubnetContainsSubnet(nil, net2) {
		t.Errorf("Expected nil contains %v", net2)
	}
}

func TestSubnetContainSubnetTrueN1GlobalN2Nil(t *testing.T) {
	_, net1, _ := net.ParseCIDR("0.0.0.0/0")
	if !SubnetContainsSubnet(net1, nil) {
		t.Errorf("Expected %v contains nil", net1)
	}
}

func TestSubnetContainSubnetTrueEqual(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

func TestSubnetContainSubnetTrueSmaller(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	_, net2, _ := net.ParseCIDR("10.1.5.128/25")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

func TestSubnetContainSubnetFalseLarger(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/25")
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

func TestSubnetContainSubnetFalseDifferent(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	_, net2, _ := net.ParseCIDR("10.1.6.0/24")
	if SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

func TestLastAddr(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.6.0/24")
	lr := net.ParseIP("10.1.6.255")
	if !LastAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", LastAddr(net1).String(), lr.String())
	}
}

func TestFirstAddr(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.6.0/24")
	lr := net.ParseIP("10.1.6.0")
	if !FirstAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", FirstAddr(net1).String(), lr.String())
	}
}

func TestNetworkID(t *testing.T) {
	ip1, net1, _ := net.ParseCIDR("10.1.6.88/24")
	net1.IP = ip1
	_, net2, _ := net.ParseCIDR("10.1.6.0/24")
	nid := NetworkID(net1)
	if !SubnetEqualSubnet(nid, net2) {
		t.Errorf("Expected %v to equal %v", nid.String(), net2.String())
	}
}

func TestRandomAddr(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.6.0/24")
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

func TestIPAdd(t *testing.T) {
	ip := net.ParseIP("10.1.5.0")
	r := net.ParseIP("10.1.5.1")
	if !IPAdd(ip, 1).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

func TestIPAddCarryover(t *testing.T) {
	ip := net.ParseIP("10.1.5.255")
	r := net.ParseIP("10.1.6.4")
	if !IPAdd(ip, 5).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

func TestIPAddCarryover2(t *testing.T) {
	ip := net.ParseIP("10.1.255.255")
	r := net.ParseIP("10.2.0.4")
	if !IPAdd(ip, 5).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

func TestIPSub(t *testing.T) {
	ip := net.ParseIP("10.1.5.5")
	r := net.ParseIP("10.1.5.4")
	if !IPAdd(ip, -1).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

func TestIPSubCarryover(t *testing.T) {
	ip := net.ParseIP("10.1.5.0")
	r := net.ParseIP("10.1.4.251")
	rip := IPAdd(ip, -5)
	if !rip.Equal(r) {
		t.Errorf("%v should equal %v", rip.String(), r.String())
	}
}

func TestIPSubCarryover2(t *testing.T) {
	ip := net.ParseIP("10.1.0.0")
	r := net.ParseIP("10.0.255.251")
	rip := IPAdd(ip, -5)
	if !rip.Equal(r) {
		t.Errorf("%v should equal %v", rip.String(), r.String())
	}
}
