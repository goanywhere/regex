/* ----------------------------------------------------------------------
 *       ______      ___                         __
 *      / ____/___  /   |  ____  __  ___      __/ /_  ___  ________
 *     / / __/ __ \/ /| | / __ \/ / / / | /| / / __ \/ _ \/ ___/ _ \
 *    / /_/ / /_/ / ___ |/ / / / /_/ /| |/ |/ / / / /  __/ /  /  __/
 *    \____/\____/_/  |_/_/ /_/\__. / |__/|__/_/ /_/\___/_/   \___/
 *                            /____/
 *
 * (C) Copyright 2014 GoAnywhere (http://goanywhere.io).
 * ----------------------------------------------------------------------
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 * ----------------------------------------------------------------------*/

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
	Convey("Regex. IPv6", t, func() {
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
