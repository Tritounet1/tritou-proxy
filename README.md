# TritouProxy

## Run project

### Install air

With homebrew :

```sh
brew install air
```

### Run project in development

```sh
air
```

### Test proxy with curl

```sh
curl -v -H "Host: tritounet.fr" http://localhost
```

### Run project in production

```sh
sudo chmod +x install.sh
sh install.sh
```

## Tests

### Run tests

```bash
go test ./...              # start all the tests
go test -v ./...           # verbose (name of each test)
go test -run TestHealth    # start one test with his name
go test -cover ./...       # coverage rate
go test -race ./...        # data race detector
```

### Structure of tests

- A test file is named `xxx_test.go` and is placed **next to** the file being tested.
- A test function : `func TestXxx(t *testing.T) { ... }`.
- To report a failure: `t.Errorf(...)` (continues) or `t.Fatalf(...)` (stops this test).
- Subtests using `t.Run("name", func(t *testing.T){ ... })` help organize test cases.

## Libraries

| Package             | Role                                                                                             |
| ------------------- | ------------------------------------------------------------------------------------------------ |
| `testing`           | The foundation. `TestXxx(t *testing.T)` functions in `_test.go` files. Run with `go test ./...`. |
| `net/http/httptest` | HTTP testing. Two key tools (see below).                                                         |

### The Two Tools in `httptest`

- **`httptest.NewRecorder()`** — a mock `ResponseWriter` that captures what your handler writes (status code, headers, body) without starting a real server. For testing a handler in isolation.
- **`httptest.NewServer(handler)`** — a real, ephemeral HTTP server on a random port. Perfect for simulating a target backend and testing end-to-end proxying. Remember to use `defer server.Close()`.
