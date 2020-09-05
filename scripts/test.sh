#!/bin/bash

DATA_FOO="{\"description\": \"Mech of Brixton\",\"name\": \"foo\",\"release\": \"2019121\",\"version\": \"1.0.1\"}"
DATA_BAR="{\"description\": \"Mech of London\",\"name\": \"bar\",\"release\": \"20200802\",\"version\": \"2.0.1\"}"

curl -XPOST  -H "Content-Type: application/json"  -d "$DATA_FOO" http://localhost:9999/api/v1/mechs
curl -XPOST  -H "Content-Type: application/json"  -d "$DATA_BAR" http://localhost:9999/api/v1/mechs

curl -XGET -H "Content-Type: application/json" http://localhost:9999/api/v1/mechs
