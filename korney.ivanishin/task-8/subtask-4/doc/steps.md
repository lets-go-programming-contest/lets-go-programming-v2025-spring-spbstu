- To generate mock objects for the project directly:

```bash
mockgen -destination=mock.go -package=gendemo . Doer
```

- To do the same thing with a single `go generate` command:

```bash
go generate ./...
```
