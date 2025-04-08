echo "AST:"
go tool compile -W main.go

echo "SSA:"
GOSSAFUNC=main go tool compile main.go

echo "asm code in genssa section in ssa.html"

echo "building binary"
go build main.go