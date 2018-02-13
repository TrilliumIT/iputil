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

func TestLastAddrLongMask(t *testing.T) {
	net1 := &net.IPNet{
		IP:   net.IP{10, 1, 6, 0},
		Mask: net.IPMask{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 0},
	}
	lr := net.ParseIP("10.1.6.255")
	if !LastAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", FirstAddr(net1).String(), lr.String())
	}
}

func TestFirstAddr(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.6.0/24")
	lr := net.ParseIP("10.1.6.0")
	if !FirstAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", FirstAddr(net1).String(), lr.String())
	}
}

func TestFirstAddrLongMask(t *testing.T) {
	net1 := &net.IPNet{
		IP:   net.IP{10, 1, 6, 0},
		Mask: net.IPMask{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 0},
	}
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
		rip := RandAddr(net1)
		if !net1.Contains(rip) {
			t.Errorf("IP %v outside subnet %v", rip.String(), net1.String())
		}
	}
}

func TestRandomAddrLongMask(t *testing.T) {
	net1 := &net.IPNet{
		IP:   net.IP{10, 1, 6, 0},
		Mask: net.IPMask{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 0},
	}
	for i := 1; i <= 10; i++ {
		rip := RandAddr(net1)
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

func TestIPBefore(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ip2 := net.ParseIP("10.1.0.2")
	ret := IPBefore(ip, ip2)
	if !ret {
		t.Errorf("%v should be before %v", ip.String(), ip2.String())
	}
}

func TestIPNotBefore(t *testing.T) {
	ip := net.ParseIP("10.1.0.2")
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPBefore(ip, ip2)
	if ret {
		t.Errorf("%v should not be before %v", ip.String(), ip2.String())
	}
}

func TestIPBeforeEqual(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPBefore(ip, ip2)
	if ret {
		t.Errorf("%v should not be before %v", ip.String(), ip2.String())
	}
}

func TestIPBeforeNilFirst(t *testing.T) {
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPBefore(nil, ip2)
	if !ret {
		t.Errorf("nil should be before %v", ip2.String())
	}
}

func TestIPBeforeNilSecond(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ret := IPBefore(ip, nil)
	if ret {
		t.Errorf("%v should not be before nil", ip.String())
	}
}

func TestIPBeforeNilBoth(t *testing.T) {
	ret := IPBefore(nil, nil)
	if ret {
		t.Errorf("nil should not be before nil")
	}
}

func TestIPDiffNeg(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ip2 := net.ParseIP("10.1.0.2")
	ret := IPDiff(ip, ip2)
	if ret != -1 {
		t.Errorf("%v minus %v should be -1", ip, ip2)
	}
}

func TestIPDiffPos(t *testing.T) {
	ip := net.ParseIP("10.1.0.2")
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPDiff(ip, ip2)
	if ret != 1 {
		t.Errorf("%v minus %v should be 1", ip, ip2)
	}
}

func TestIPDiffEq(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPDiff(ip, ip2)
	if ret != 0 {
		t.Errorf("%v minus %v should be 0", ip, ip2)
	}
}

func TestRandomAddrWithExclude(t *testing.T) {
	_, sn, _ := net.ParseCIDR("10.1.0.0/24")
	ip := RandAddrWithExclude(sn, 0, 0)
	if !sn.Contains(ip) {
		t.Errorf("%v should be an IP in %v", ip.String(), sn.String())
	}
}

func TestRandomAddrWithBadExcludeFirst(t *testing.T) {
	_, sn, _ := net.ParseCIDR("10.1.0.0/24")
	ip := RandAddrWithExclude(sn, 300, 0)
	if ip != nil {
		t.Errorf("exclusions that land outside the subnet's range should return a nil IP")
	}
}

func TestRandomAddrWithBadExcludeSecond(t *testing.T) {
	_, sn, _ := net.ParseCIDR("10.1.0.0/24")
	ip := RandAddrWithExclude(sn, 0, 300)
	if ip != nil {
		t.Errorf("exclusions that land outside the subnet's range should return a nil IP")
	}
}

func TestRandomAddrWithBadExcludeBoth(t *testing.T) {
	_, sn, _ := net.ParseCIDR("10.1.0.0/24")
	ip := RandAddrWithExclude(sn, 150, 150)
	if ip != nil {
		t.Errorf("exclusions that land outside the subnet's range should return a nil IP")
	}
}
