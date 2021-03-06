// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016-2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package builtin

const timeControlSummary = `allows setting system date and time`

const timeControlBaseDeclarationSlots = `
  time-control:
    allow-installation:
      slot-snap-type:
        - core
    deny-auto-connection: true
`

const timeControlConnectedPlugAppArmor = `
# Description: Can set time and date via systemd' timedated D-Bus interface.
# Can read all properties of /org/freedesktop/timedate1 D-Bus object; see
# https://www.freedesktop.org/wiki/Software/systemd/timedated/; This also
# gives full access to the RTC device nodes and relevant parts of sysfs.

#include <abstractions/dbus-strict>

# Introspection of org.freedesktop.timedate1
dbus (send)
    bus=system
    path=/org/freedesktop/timedate1
    interface=org.freedesktop.DBus.Introspectable
    member=Introspect
    peer=(label=unconfined),

dbus (send)
    bus=system
    path=/org/freedesktop/timedate1
    interface=org.freedesktop.timedate1
    member="Set{Time,LocalRTC}"
    peer=(label=unconfined),

# Read all properties from timedate1
dbus (send)
    bus=system
    path=/org/freedesktop/timedate1
    interface=org.freedesktop.DBus.Properties
    member=Get{,All}
    peer=(label=unconfined),

# Receive timedate1 property changed events
dbus (receive)
    bus=system
    path=/org/freedesktop/timedate1
    interface=org.freedesktop.DBus.Properties
    member=PropertiesChanged
    peer=(label=unconfined),

# As the core snap ships the timedatectl utility we can also allow
# clients to use it now that they have access to the relevant
# D-Bus methods for setting the time via timedatectl's set-time and
# set-local-rtc commands.
/usr/bin/timedatectl{,.real} ixr,

# Allow write access to system real-time clock
# See 'man 4 rtc' for details.

capability sys_time,

/dev/rtc[0-9]* rw,

# Access to the sysfs nodes are needed by rtcwake for example
# to program scheduled wakeups in the future.
/sys/class/rtc/*/ rw,
/sys/class/rtc/*/** rw,

# As the core snap ships the hwclock utility we can also allow
# clients to use it now that they have access to the relevant
# device nodes.
/sbin/hwclock ixr,
`

const timeControlConnectedPlugUDev = `SUBSYSTEM=="rtc", TAG+="###CONNECTED_SECURITY_TAGS###"`

func init() {
	registerIface(&commonInterface{
		name:                  "time-control",
		summary:               timeControlSummary,
		implicitOnCore:        true,
		implicitOnClassic:     true,
		baseDeclarationSlots:  timeControlBaseDeclarationSlots,
		connectedPlugAppArmor: timeControlConnectedPlugAppArmor,
		connectedPlugUDev:     timeControlConnectedPlugUDev,
		reservedForOS:         true,
	})
}
