package iputil

import (
	"fmt"
	"net"
	"testing"
)

// nolint dupl
func TestSubnetEqualSubnetTrue6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80::/64")
	if !SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to equal %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetEqualSubnetFalseDifferentMask6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80::/48")
	if SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetEqualSubnetFalseDifferentNet6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80:1::/64")
	if SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetEqualSubnetTrueIpInNet6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	ip2, net2, _ := net.ParseCIDR("fe80::1/64")
	net2.IP = ip2
	if !SubnetEqualSubnet(net1, net2) {
		t.Errorf("Expected %v to not equal %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetEqualSubnetFalseN1Nil6(t *testing.T) {
	_, net2, _ := net.ParseCIDR("fe80::/64")
	if SubnetEqualSubnet(nil, net2) {
		t.Errorf("Expected nil not equal %v", net2)
	}
}

// nolint dupl
func TestSubnetEqualSubnetFalseN2Nil6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	if SubnetEqualSubnet(net1, nil) {
		t.Errorf("Expected %v not equal nil", net1)
	}
}

// nolint dupl
func TestSubnetEqualSubnetTrueN1NilN2Global6(t *testing.T) {
	_, net2, _ := net.ParseCIDR("::/0")
	if !SubnetEqualSubnet(nil, net2) {
		t.Errorf("Expected nil equal %v", net2)
	}
}

// nolint dupl
func TestSubnetEqualSubnetTrueN1GlobalN2Nil6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("::/0")
	if !SubnetEqualSubnet(net1, nil) {
		t.Errorf("Expected %v equal nil", net1)
	}
}

// nolint dupl
func TestSubnetContainSubnetFalseN1Nil6(t *testing.T) {
	_, net2, _ := net.ParseCIDR("fe80::/64")
	if !SubnetContainsSubnet(nil, net2) {
		t.Errorf("Expected nil contains %v", net2)
	}
}

// nolint dupl
func TestSubnetContainSubnetFalseN2Nil6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	if SubnetContainsSubnet(net1, nil) {
		t.Errorf("Expected %v not contains nil", net1)
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueN1NilN2Global6(t *testing.T) {
	_, net2, _ := net.ParseCIDR("::/0")
	if !SubnetContainsSubnet(nil, net2) {
		t.Errorf("Expected nil contains %v", net2)
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueN1GlobalN2Nil6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("::/0")
	if !SubnetContainsSubnet(net1, nil) {
		t.Errorf("Expected %v contains nil", net1)
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueEqual6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80::/64")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueSmaller6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/48")
	_, net2, _ := net.ParseCIDR("fe80::/64")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetContainSubnetTrueSmallerDiff6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/48")
	_, net2, _ := net.ParseCIDR("fe80:0:0:1::/64")
	if !SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetContainSubnetFalseLarger6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80::/48")
	if SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestSubnetContainSubnetFalseDifferent6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	_, net2, _ := net.ParseCIDR("fe80:1::/64")
	if SubnetContainsSubnet(net1, net2) {
		t.Errorf("Expected %v to contain %v", net1.String(), net2.String())
	}
}

// nolint dupl
func TestLastAddr6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	lr := net.ParseIP("fe80::ffff:ffff:ffff:ffff")
	if !LastAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", LastAddr(net1).String(), lr.String())
	}
}

// nolint dupl
func TestFirstAddr6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::55/64")
	lr := net.ParseIP("fe80::")
	if !FirstAddr(net1).Equal(lr) {
		t.Errorf("Expected %v to equal %v", FirstAddr(net1).String(), lr.String())
	}
}

// nolint dupl
func TestNetworkID6(t *testing.T) {
	ip1, net1, _ := net.ParseCIDR("fe80::55/64")
	net1.IP = ip1
	_, net2, _ := net.ParseCIDR("fe80::/64")
	nid := NetworkID(net1)
	if !SubnetEqualSubnet(nid, net2) {
		t.Errorf("Expected %v to equal %v", nid.String(), net2.String())
	}
}

// nolint dupl
func TestRandomAddr6(t *testing.T) {
	_, net1, _ := net.ParseCIDR("fe80::/64")
	for i := 1; i <= 10; i++ {
		rip := RandAddr(net1)
		if !net1.Contains(rip) {
			t.Errorf("IP %v outside subnet %v", rip.String(), net1.String())
		}
	}
}

// nolint dupl
func TestIPAdd6(t *testing.T) {
	ip := net.ParseIP("fe80::")
	r := net.ParseIP("fe80::1")
	if !IPAdd(ip, 1).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

// nolint dupl
func TestIPAddCarryover6(t *testing.T) {
	ip := net.ParseIP("fe80::ffff")
	r := net.ParseIP("fe80::1:4")
	if !IPAdd(ip, 5).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

// nolint dupl
func TestIPAddCarryover26(t *testing.T) {
	ip := net.ParseIP("fe80::ffff:ffff")
	r := net.ParseIP("fe80::1:0:4")
	if !IPAdd(ip, 5).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

// nolint dupl
func TestIPSub6(t *testing.T) {
	ip := net.ParseIP("fe80::5")
	r := net.ParseIP("fe80::4")
	if !IPAdd(ip, -1).Equal(r) {
		t.Errorf("%v should equal %v", ip.String(), r.String())
	}
}

// nolint dupl
func TestIPSubCarryover6(t *testing.T) {
	ip := net.ParseIP("fe80::100")
	r := net.ParseIP("fe80::fb")
	rip := IPAdd(ip, -5)
	if !rip.Equal(r) {
		t.Errorf("%v should equal %v", rip.String(), r.String())
	}
}

// nolint dupl
func TestIPSubCarryover26(t *testing.T) {
	ip := net.ParseIP("fe80::1:0")
	r := net.ParseIP("fe80::fffb")
	rip := IPAdd(ip, -5)
	if !rip.Equal(r) {
		t.Errorf("%v should equal %v", rip.String(), r.String())
	}
}

// nolint dupl
func TestIPBefore6(t *testing.T) {
	ip := net.ParseIP("fe80::1")
	ip2 := net.ParseIP("fe80::2")
	ret := IPBefore(ip, ip2)
	if !ret {
		t.Errorf("%v should be before %v", ip.String(), ip2.String())
	}
}

// nolint dupl
func TestIPNotBefore6(t *testing.T) {
	ip := net.ParseIP("fe80::2")
	ip2 := net.ParseIP("fe80::1")
	ret := IPBefore(ip, ip2)
	if ret {
		t.Errorf("%v should not be before %v", ip.String(), ip2.String())
	}
}

// nolint dupl
func TestIPBeforeEqual6(t *testing.T) {
	ip := net.ParseIP("fe80::1")
	ip2 := net.ParseIP("fe80::1")
	ret := IPBefore(ip, ip2)
	if ret {
		t.Errorf("%v should not be before %v", ip.String(), ip2.String())
	}
}

// nolint dupl
func TestIPBeforeNilFirst6(t *testing.T) {
	ip2 := net.ParseIP("fe80::1")
	ret := IPBefore(nil, ip2)
	if !ret {
		t.Errorf("nil should be before %v", ip2.String())
	}
}

// nolint dupl
func TestIPBeforeNilSecond6(t *testing.T) {
	ip := net.ParseIP("fe80::1")
	ret := IPBefore(ip, nil)
	if ret {
		t.Errorf("%v should not be before nil", ip.String())
	}
}

// nolint dupl
func TestMakeSameLengthWithNil6(t *testing.T) {
	ip := net.ParseIP("fe80::1")
	nip, nip2 := makeNilZero(ip, nil)
	_, nip2 = makeSameLength(nip, nip2)
	if !nip2.Equal(net.IPv6zero) {
		fmt.Printf("nip2: %v", nip2.String())
		t.Errorf("nip2 should be zero")
	}
}

// nolint dupl
func TestIPDiffNeg6(t *testing.T) {
	ip := net.ParseIP("fe80::1")
	ip2 := net.ParseIP("fe80::2")
	ret := IPDiff(ip, ip2)
	if ret != -1 {
		t.Errorf("%v minus %v should be -1", ip, ip2)
	}
}

// nolint dupl
func TestIPDiffPos6(t *testing.T) {
	ip := net.ParseIP("fe80::2")
	ip2 := net.ParseIP("fe80::1")
	ret := IPDiff(ip, ip2)
	if ret != 1 {
		t.Errorf("%v minus %v should be 1", ip, ip2)
	}
}

// nolint dupl
func TestIPDiffEq6(t *testing.T) {
	ip := net.ParseIP("fe80::1")
	ip2 := net.ParseIP("fe80::1")
	ret := IPDiff(ip, ip2)
	if ret != 0 {
		t.Errorf("%v minus %v should be 0", ip, ip2)
	}
}

// nolint dupl
func TestRandomAddrWithExclude6(t *testing.T) {
	_, sn, _ := net.ParseCIDR("fe80::/120")
	ip := RandAddrWithExclude(sn, 0, 0)
	if !sn.Contains(ip) {
		t.Errorf("%v should be an IP in %v", ip.String(), sn.String())
	}
}

// nolint dupl
func TestRandomAddrWithBadExcludeFirst6(t *testing.T) {
	_, sn, _ := net.ParseCIDR("fe80::/120")
	ip := RandAddrWithExclude(sn, 300, 0)
	if ip != nil {
		t.Errorf("exclusions that land outside the subnet's range should return a nil IP")
	}
}

// nolint dupl
func TestRandomAddrWithBadExcludeSecond6(t *testing.T) {
	_, sn, _ := net.ParseCIDR("fe80::/120")
	ip := RandAddrWithExclude(sn, 0, 300)
	if ip != nil {
		t.Errorf("exclusions that land outside the subnet's range should return a nil IP")
	}
}

// nolint dupl
func TestRandomAddrWithBadExcludeBoth6(t *testing.T) {
	_, sn, _ := net.ParseCIDR("fe80::/120")
	ip := RandAddrWithExclude(sn, 150, 150)
	if ip != nil {
		t.Errorf("exclusions that land outside the subnet's range should return a nil IP")
	}
}

// nolint dupl
func TestCIDRToIPNet6(t *testing.T) {
	ip := "fe80::1/64"
	ipnet, _ := CIDRToIPNet(ip)
	if !ipnet.IP.Equal(net.ParseIP("fe80::1")) {
		t.Errorf("ip of %v should equal 10.1.0.1", ipnet.String())
	}
}
