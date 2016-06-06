// -*- Mode: Go; indent-tabs-mode: t -*-
// +build !excludeintegration

/*
 * Copyright (C) 2016 Canonical Ltd
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

package tests

import (
	"github.com/snapcore/snapd/integration-tests/testutils/cli"
	"github.com/snapcore/snapd/integration-tests/testutils/common"

	"gopkg.in/check.v1"
)

var _ = check.Suite(&installedUnitsSuite{})

type installedUnitsSuite struct {
	common.SnappySuite
}

func (s *installedUnitsSuite) TestInstalledServicesAreUp(c *check.C) {
	for _, unit := range []string{"snapd.socket", "snapd.refresh.timer"} {
		output := cli.ExecCommand(c, "systemctl", "status", unit)

		c.Check(output, check.Matches, "(?ms).*\n +Active: active.*")
	}
}
