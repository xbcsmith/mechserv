// SPDX-FileCopyrightText: © 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("FOO", "1")
	foo := GetEnv("FOO", "2")
	assert.Assert(t, foo == "1")
	bar := GetEnv("BAR", "42")
	assert.Assert(t, bar == "42")
}

func TestNewULID(t *testing.T) {
	u := NewULID()
	assert.Assert(t, len(u) == 26)
}

func TestDecodeMechFromJson(t *testing.T) {
	content := `{
  "description": "Mech from Brixton",
  "id": "01DM3PADN0FQFGD2SDJR04DGF8",
  "name": "foo",
  "release": "20190906.1567787180",
  "version": "0.0.1"
} `

	mech, err := DecodeMechFromJSON(strings.NewReader(content))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, mech.Name, "foo")
}

func TestNewMech(t *testing.T) {
	mech := NewMech()
	assert.Assert(t, len(mech.ID) == 26)
}

func TestNewConfig(t *testing.T) {
	cfg := NewConfig("localhost", "9999")
	assert.Assert(t, cfg.Host == "localhost")
}
