1. Default build:
```
go build .
 ```
2. Build with flags:
```
go build -ldflags "-X main.version=1000 -X main.date=$(date +'%Y-%m-%d')" -o flags .
```
