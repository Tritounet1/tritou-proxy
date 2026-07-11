# TritouProxy

## Run project in dev

### Install air

With homebrew :

```sh
brew install air
```

### Run project

```sh
air
```

### Test proxy

```sh
curl -v -H "Host: tritounet.fr" http://localhost
```

### Run tests

```bash
go test ./...              # start all the tests
go test -v ./...           # verbose (name of each test)
go test -run TestHealth    # start one test with his name
go test -cover ./...       # coverage rate
go test -race ./...        # data race detector
```

### Structure of tests

- Un fichier de test se nomme `xxx_test.go` et se place **à côté** du fichier testé.
- Une fonction de test : `func TestXxx(t *testing.T) { ... }`.
- Pour signaler un échec : `t.Errorf(...)` (continue) ou `t.Fatalf(...)` (arrête ce test).
- Les sous-tests via `t.Run("nom", func(t *testing.T){ ... })` aident à organiser les cas.
