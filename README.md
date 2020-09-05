# mechserv

## build

`make`

## test

```bash
DATA="{\"description\": \"Mech of Brixton\",\"name\": \"foo\",\"release\": \"2019121\",\"version\": \"1.0.1\"}"

curl -XPOST  -H "Content-Type: application/json"  -d "$DATA" http://localhost:9999/api/v1/mechs

curl -XGET -H "Content-Type: application/json" http://localhost:9999/api/v1/mechs
```

```bash
export MECHSERV_HOST="http://localhost"
export MECHSERV_PORT="9999"

# Create initial mechs
./mechit.py init

# Create more mechs
./mechit.py create --name "brak" --version "11.1.2" --description "Mech of the Barbarian" --debug
```
