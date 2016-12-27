package regex

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEmail(t *testing.T) {
	Convey("Regex. Email", t, func() {
		So(Email.MatchString("golang@goanywhere.io"), ShouldBeTrue)
		So(Email.MatchString("golang@goanywhere"), ShouldBeFalse)
	})
}

func TestIPv4(t *testing.T) {
	Convey("Regex. IPv4", t, func() {
		So(IPv4.MatchString("8.8.8.8"), ShouldBeTrue)
		So(IPv4.MatchString("8.8.8."), ShouldBeFalse)
		So(IPv4.MatchString("192.168.1.1"), ShouldBeTrue)
	})
}

func TestIPv6(t *testing.T) {
	Convey("Regex. IPv6 (Simple Match)", t, func() {
		So(IPv6.MatchString("1:2:3:4:5:6:7:8"), ShouldBeTrue)
		So(IPv6.MatchString("1::"), ShouldBeTrue)
		So(IPv6.MatchString("1:2:3:4:5:6:7::"), ShouldBeTrue)
		So(IPv6.MatchString("1::8"), ShouldBeTrue)
		So(IPv6.MatchString("1:2:3:4:5:6::8"), ShouldBeTrue)
		So(IPv6.MatchString("1::7:8"), ShouldBeTrue)
		So(IPv6.MatchString("1:2:3:4:5::7:8"), ShouldBeTrue)
		So(IPv6.MatchString("1:2:3:4:5::8"), ShouldBeTrue)
		So(IPv6.MatchString("1::6:7:8"), ShouldBeTrue)
		So(IPv6.MatchString("1:2:3:4::6:7:8"), ShouldBeTrue)
		So(IPv6.MatchString("1:2:3:4::8"), ShouldBeTrue)
		So(IPv6.MatchString("1::5:6:7:8"), ShouldBeTrue)
		So(IPv6.MatchString("1:2:3::5:6:7:8"), ShouldBeTrue)
		So(IPv6.MatchString("1:2:3::8"), ShouldBeTrue)
		So(IPv6.MatchString("1::4:5:6:7:8"), ShouldBeTrue)
		So(IPv6.MatchString("1:2::4:5:6:7:8"), ShouldBeTrue)
		So(IPv6.MatchString("1:2::8"), ShouldBeTrue)
		So(IPv6.MatchString("1::3:4:5:6:7:8"), ShouldBeTrue)
		So(IPv6.MatchString("::2:3:4:5:6:7:8"), ShouldBeTrue)
		// link-local IPv6 addresses with zone index
		So(IPv6.MatchString("fe80::7:8%eth0"), ShouldBeTrue)
		So(IPv6.MatchString("fe80::7:8%eth1"), ShouldBeTrue)
		So(IPv6.MatchString("fe80::7:8%1"), ShouldBeTrue)
		// IPv4-mapped IPv6 addresses and IPv4-translated addresses
		So(IPv6.MatchString("::255.255.255.255"), ShouldBeTrue)
		So(IPv6.MatchString("::ffff:255.255.255.255"), ShouldBeTrue)
		So(IPv6.MatchString("::ffff:0:255.255.255.255"), ShouldBeTrue)
		// IPv4-Embedded IPv6 Address
		So(IPv6.MatchString("2001:db8:3:4::192.0.2.33"), ShouldBeTrue)
		So(IPv6.MatchString("64:ff9b::192.0.2.33"), ShouldBeTrue)
	})
	Convey("Regex. IPv6 (Find String)", t, func() {
		So(IPv6.FindString(`1:2:3:4:5:6:7:8`), ShouldEqual, `1:2:3:4:5:6:7:8`)
		So(IPv6.FindString(`1::`), ShouldEqual, `1::`)
		So(IPv6.FindString(`1:2:3:4:5:6:7::`), ShouldEqual, `1:2:3:4:5:6:7::`)
		So(IPv6.FindString(`1::8`), ShouldEqual, `1::8`)
		So(IPv6.FindString(`1:2:3:4:5:6::8`), ShouldEqual, `1:2:3:4:5:6::8`)
		So(IPv6.FindString(`1::7:8`), ShouldEqual, `1::7:8`)
		So(IPv6.FindString(`1:2:3:4:5::7:8`), ShouldEqual, `1:2:3:4:5::7:8`)
		So(IPv6.FindString(`1:2:3:4:5::8`), ShouldEqual, `1:2:3:4:5::8`)
		So(IPv6.FindString(`1::6:7:8`), ShouldEqual, `1::6:7:8`)
		So(IPv6.FindString(`1:2:3:4::6:7:8`), ShouldEqual, `1:2:3:4::6:7:8`)
		So(IPv6.FindString(`1:2:3:4::8`), ShouldEqual, `1:2:3:4::8`)
		So(IPv6.FindString(`1::5:6:7:8`), ShouldEqual, `1::5:6:7:8`)
		So(IPv6.FindString(`1:2:3::5:6:7:8`), ShouldEqual, `1:2:3::5:6:7:8`)
		So(IPv6.FindString(`1:2:3::8`), ShouldEqual, `1:2:3::8`)
		So(IPv6.FindString(`1::4:5:6:7:8`), ShouldEqual, `1::4:5:6:7:8`)
		So(IPv6.FindString(`1:2::4:5:6:7:8`), ShouldEqual, `1:2::4:5:6:7:8`)
		So(IPv6.FindString(`1:2::8`), ShouldEqual, `1:2::8`)
		So(IPv6.FindString(`1::3:4:5:6:7:8`), ShouldEqual, `1::3:4:5:6:7:8`)
		So(IPv6.FindString(`::2:3:4:5:6:7:8`), ShouldEqual, `::2:3:4:5:6:7:8`)
		// link-local IPv6 addresses with zone index
		So(IPv6.FindString(`fe80::7:8%eth0`), ShouldEqual, `fe80::7:8%eth0`)
		So(IPv6.FindString(`fe80::7:8%eth1`), ShouldEqual, `fe80::7:8%eth1`)
		So(IPv6.FindString(`fe80::7:8%1`), ShouldEqual, `fe80::7:8%1`)
		// IPv4-mapped IPv6 addresses and IPv4-translated addresses
		So(IPv6.FindString(`::255.255.255.255`), ShouldEqual, `::255.255.255.255`)
		So(IPv6.FindString(`::ffff:255.255.255.255`), ShouldEqual, `::ffff:255.255.255.255`)
		So(IPv6.FindString(`::ffff:0:255.255.255.255`), ShouldEqual, `::ffff:0:255.255.255.255`)
		// IPv4-Embedded IPv6 Address
		So(IPv6.FindString(`2001:db8:3:4::192.0.2.33`), ShouldEqual, `2001:db8:3:4::192.0.2.33`)
		So(IPv6.FindString(`64:ff9b::192.0.2.33`), ShouldEqual, `64:ff9b::192.0.2.33`)
	})
}

func TestURL(t *testing.T) {
	Convey("regex.URL", t, func() {
		So(URL.MatchString("https://github.com"), ShouldBeTrue)
		So(URL.MatchString("https://api.example.com/v1/75KImuE2obLJrYpWOAH4OE"), ShouldBeTrue)
		So(URL.MatchString("api.example.com/v1/75KImuE2obLJrYpWOAH4OE"), ShouldBeTrue)
		So(URL.MatchString("api.io"), ShouldBeTrue)
		So(URL.MatchString("api"), ShouldBeFalse)
	})
}
