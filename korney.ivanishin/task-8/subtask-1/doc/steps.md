`go tool compile` *for some reason has trouble linking modules in version 1.23*

- To generate the AST:

```bash
go build -gcflags=-W
```

- To generate SSA IR and assembly code for a main.go file; generate `main` object file:

```bash
GOSSAFUNC=main go build cmd/service/main.go > ssa.html
```

- To build the final executable:

```bash
go build -o service
```
