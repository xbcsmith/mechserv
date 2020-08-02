package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
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

func TestApi(t *testing.T) {
	cfg := NewConfig("localhost", "9999")
	router := NewAPI(cfg)
	fmt.Printf("%v\n", router)
}
