package iputil

import (
	"net"
	"testing"
)

// nolint dupl
func TestSubnetEqualSubnetTrue(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if !SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to equal %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetEqualSubnetFalseDifferentMask(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.0.0/16")
	_, net2, _ := net.ParseCIDR("10.1.0.0/24")
	if SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetEqualSubnetFalseDifferentNet(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.0.0/16")
	_, net2, _ := net.ParseCIDR("10.2.0.0/16")
	if SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetEqualSubnetTrueIpInNet(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.6/16")
	ip2, net2, _ := net.ParseCIDR("10.1.2.1/16")
	net2.IP = ip2
	if !SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetEqualSubnetTrueBothNil(t *testing.T) {
	if !SubnetEqualSubnet(nil, nil) {
		t.Error("Expected nil subnets to be equal.")
	}
}

// nolint dupl
func TestSubnetEqualSubnetFalseN1Nil(t *testing.T) {
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if SubnetEqualSubnet(nil, net2) {
		t.Errorf("Expected nil not equal %v", net2)
	}
}

// nolint dupl
func TestSubnetEqualSubnetFalseN2Nil(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	if SubnetEqualSubnet(net1, nil) {
		t.Errorf("Expected %v not equal nil", net1)
	}
}

// nolint dupl
func TestSubnetEqualSubnetTrueN1NilN2Global(t *testing.T) {
	_, net2, _ := net.ParseCIDR("0.0.0.0/0")
	if !SubnetEqualSubnet(nil, net2) {
		t.Errorf("Expected nil equal %v", net2)
	}
}

// nolint dupl
func TestSubnetEqualSubnetTrueN1GlobalN2Nil(t *testing.T) {
	_, net1, _ := net.ParseCIDR("0.0.0.0/0")
	if !SubnetEqualSubnet(net1, nil) {
		t.Errorf("Expected %v equal nil", net1)
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueBothNil(t *testing.T) {
	if !SubnetContainsSubnet(nil, nil) {
		t.Error("Expected nil contains nil.")
	}
}

// nolint dupl
func TestSubnetContainSubnetFalseN1Nil(t *testing.T) {
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if !SubnetContainsSubnet(nil, net2) {
		t.Errorf("Expected nil contains %v", net2)
	}
}

// nolint dupl
func TestSubnetContainSubnetFalseN2Nil(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	if SubnetContainsSubnet(net1, nil) {
		t.Errorf("Expected %v not contains nil", net1)
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueN1NilN2Global(t *testing.T) {
	_, net2, _ := net.ParseCIDR("0.0.0.0/0")
	if !SubnetContainsSubnet(nil, net2) {
		t.Errorf("Expected nil contains %v", net2)
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueN1GlobalN2Nil(t *testing.T) {
	_, net1, _ := net.ParseCIDR("0.0.0.0/0")
	if !SubnetContainsSubnet(net1, nil) {
		t.Errorf("Expected %v contains nil", net1)
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueEqual(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueSmaller(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	_, net2, _ := net.ParseCIDR("10.1.5.128/25")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetContainSubnetFalseLarger(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/25")
	_, net2, _ := net.ParseCIDR("10.1.5.0/24")
	if SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetContainSubnetFalseDifferent(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.5.0/24")
	_, net2, _ := net.ParseCIDR("10.1.6.0/24")
	if SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

// nolint dupl
// nolint dupl
func TestLastAddr(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.6.0/24")
	lr := net.ParseIP("10.1.6.255")
	if !LastAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", LastAddr(net1).String(), lr.String())
	}
}

// nolint dupl
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

// nolint dupl
// nolint dupl
func TestFirstAddr(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.6.0/24")
	lr := net.ParseIP("10.1.6.0")
	if !FirstAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", FirstAddr(net1).String(), lr.String())
	}
}

// nolint dupl
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

// nolint dupl
func TestNetworkID(t *testing.T) {
	ip1, net1, _ := net.ParseCIDR("10.1.6.88/24")
	net1.IP = ip1
	_, net2, _ := net.ParseCIDR("10.1.6.0/24")
	nid := NetworkID(net1)
	if !SubnetEqualSubnet(nid, net2) {
		t.Errorf("Expected %v to equal %v", nid.String(), net2.String())
	}
}

// nolint dupl
func TestRandomAddr(t *testing.T) {
	_, net1, _ := net.ParseCIDR("10.1.6.0/24")
	for i := 1; i <= 10; i++ {
		rip := RandAddr(net1)
		if !net1.Contains(rip) {
			t.Errorf("IP %v outside subnet %v", rip.String(), net1.String())
		}
	}
}

// nolint dupl
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

// nolint dupl
// nolint dupl
func TestIPAdd(t *testing.T) {
	ip := net.ParseIP("10.1.5.0")
	r := net.ParseIP("10.1.5.1")
	if !IPAdd(ip, 1).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

// nolint dupl
func TestIPAddCarryover(t *testing.T) {
	ip := net.ParseIP("10.1.5.255")
	r := net.ParseIP("10.1.6.4")
	if !IPAdd(ip, 5).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

// nolint dupl
func TestIPAddCarryover2(t *testing.T) {
	ip := net.ParseIP("10.1.255.255")
	r := net.ParseIP("10.2.0.4")
	if !IPAdd(ip, 5).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

// nolint dupl
func TestIPSub(t *testing.T) {
	ip := net.ParseIP("10.1.5.5")
	r := net.ParseIP("10.1.5.4")
	if !IPAdd(ip, -1).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

// nolint dupl
// nolint dupl
func TestIPSubCarryover(t *testing.T) {
	ip := net.ParseIP("10.1.5.0")
	r := net.ParseIP("10.1.4.251")
	rip := IPAdd(ip, -5)
	if !rip.Equal(r) {
		t.Errorf("%v should equal %v", rip.String(), r.String())
	}
}

// nolint dupl
// nolint dupl
func TestIPSubCarryover2(t *testing.T) {
	ip := net.ParseIP("10.1.0.0")
	r := net.ParseIP("10.0.255.251")
	rip := IPAdd(ip, -5)
	if !rip.Equal(r) {
		t.Errorf("%v should equal %v", rip.String(), r.String())
	}
}

// nolint dupl
func TestIPBefore(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ip2 := net.ParseIP("10.1.0.2")
	ret := IPBefore(ip, ip2)
	if !ret {
		t.Errorf("%v should be before %v", ip.String(), ip2.String())
	}
}

// nolint dupl
func TestIPNotBefore(t *testing.T) {
	ip := net.ParseIP("10.1.0.2")
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPBefore(ip, ip2)
	if ret {
		t.Errorf("%v should not be before %v", ip.String(), ip2.String())
	}
}

// nolint dupl
func TestIPBeforeEqual(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPBefore(ip, ip2)
	if ret {
		t.Errorf("%v should not be before %v", ip.String(), ip2.String())
	}
}

// nolint dupl
func TestIPBeforeNilFirst(t *testing.T) {
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPBefore(nil, ip2)
	if !ret {
		t.Errorf("nil should be before %v", ip2.String())
	}
}

// nolint dupl
func TestIPBeforeNilSecond(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ret := IPBefore(ip, nil)
	if ret {
		t.Errorf("%v should not be before nil", ip.String())
	}
}

// nolint dupl
func TestIPBeforeNilBoth(t *testing.T) {
	ret := IPBefore(nil, nil)
	if ret {
		t.Errorf("nil should not be before nil")
	}
}

// nolint dupl
func TestIPDiffNeg(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ip2 := net.ParseIP("10.1.0.2")
	ret := IPDiff(ip, ip2)
	if ret != -1 {
		t.Errorf("%v minus %v should be -1", ip, ip2)
	}
}

// nolint dupl
func TestIPDiffPos(t *testing.T) {
	ip := net.ParseIP("10.1.0.2")
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPDiff(ip, ip2)
	if ret != 1 {
		t.Errorf("%v minus %v should be 1", ip, ip2)
	}
}

// nolint dupl
func TestIPDiffEq(t *testing.T) {
	ip := net.ParseIP("10.1.0.1")
	ip2 := net.ParseIP("10.1.0.1")
	ret := IPDiff(ip, ip2)
	if ret != 0 {
		t.Errorf("%v minus %v should be 0", ip, ip2)
	}
}

// nolint dupl
func TestRandomAddrWithExclude(t *testing.T) {
	_, sn, _ := net.ParseCIDR("10.1.0.0/24")
	ip := RandAddrWithExclude(sn, 0, 0)
	if !sn.Contains(ip) {
		t.Errorf("%v should be an IP in %v", ip.String(), sn.String())
	}
}

// nolint dupl
func TestRandomAddrWithBadExcludeFirst(t *testing.T) {
	_, sn, _ := net.ParseCIDR("10.1.0.0/24")
	ip := RandAddrWithExclude(sn, 300, 0)
	if ip != nil {
		t.Errorf("exclusions that land outside the subnet's range should return a nil IP")
	}
}

// nolint dupl
func TestRandomAddrWithBadExcludeSecond(t *testing.T) {
	_, sn, _ := net.ParseCIDR("10.1.0.0/24")
	ip := RandAddrWithExclude(sn, 0, 300)
	if ip != nil {
		t.Errorf("exclusions that land outside the subnet's range should return a nil IP")
	}
}

// nolint dupl
func TestRandomAddrWithBadExcludeBoth(t *testing.T) {
	_, sn, _ := net.ParseCIDR("10.1.0.0/24")
	ip := RandAddrWithExclude(sn, 150, 150)
	if ip != nil {
		t.Errorf("exclusions that land outside the subnet's range should return a nil IP")
	}
}

// nolint dupl
func TestMakeSameLenghtNoChange(t *testing.T) {
	ip := net.IP{10, 10, 10, 10}
	oip := net.IP{10, 10, 10, 10}
	ip2 := net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 10, 10, 10, 11}
	oip2 := net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 10, 10, 10, 11}

	nip, nip2 := makeSameLength(ip, ip2)

	if !strictEqual(oip, ip) {
		t.Errorf("value of ip should not have changed")
	}
	if !strictEqual(oip2, ip2) {
		t.Errorf("value of ip2 should not have changed")
	}

	if !ip.Equal(nip) {
		t.Errorf("ip should equal value nip")
	}

	if !ip2.Equal(nip2) {
		t.Errorf("ip should equal value of nip")
	}

	if strictEqual(ip, nip) {
		t.Errorf("ip should not strictly equal nip")
	}
}

// nolint dupl
func TestMakeNilZeroNoChange(t *testing.T) {
	var ip, oip net.IP
	ip = nil
	oip = nil
	ip2 := net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 10, 10, 10, 11}
	oip2 := net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 10, 10, 10, 11}

	nip, nip2 := makeNilZero(ip, ip2)

	if !strictEqual(oip, ip) {
		t.Errorf("value of ip should not have changed")
	}
	if !strictEqual(oip2, ip2) {
		t.Errorf("value of ip2 should not have changed")
	}

	if ip != nil {
		t.Errorf("ip should equal value nil")
	}

	if !nip.Equal(net.IPv4zero) {
		t.Errorf("nip should equal zero")
	}

	if !ip2.Equal(nip2) {
		t.Errorf("ip should equal value of nip")
	}

	if strictEqual(ip, nip) {
		t.Errorf("ip should not strictly equal nip")
	}
}

// nolint dupl
func strictEqual(ip, ip2 net.IP) bool {
	if len(ip) != len(ip2) {
		return false
	}
	for i := range ip {
		if ip[i] != ip2[i] {
			return false
		}
	}
	return true
}

// nolint dupl
func TestCIDRToIPNet(t *testing.T) {
	ip := "10.1.0.1/24"
	ipnet, _ := CIDRToIPNet(ip)
	if !ipnet.IP.Equal(net.ParseIP("10.1.0.1")) {
		t.Errorf("ip of %v should equal 10.1.0.1", ipnet.String())
	}
}

// nolint dupl
func TestBadCIDRToIPNet(t *testing.T) {
	ip := "not an ip address"
	_, err := CIDRToIPNet(ip)
	if err == nil {
		t.Errorf("parsing something not a CIDR should return an error")
	}
}
