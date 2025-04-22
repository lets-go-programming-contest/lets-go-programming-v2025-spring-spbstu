1. Построение AST (Abstract Syntax Tree):
```
go build -gcflags="-S" main.go
```
2. Генерация SSA (Static Single Assignment):
```
GOSSAFUNC=main go build main.go
```
3. Ассемблерный код:
```
go tool compile -S main.go
```
4. Создание объектного файла:
```
go tool compile -o main.o main.go 
```
5. Генерация исполняемого файла:
```
go build -o bubble_sort main.go
```
