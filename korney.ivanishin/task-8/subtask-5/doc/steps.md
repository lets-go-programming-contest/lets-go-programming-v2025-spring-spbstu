- To check that embedded files' contents can be directly used by the program:

```bash
echo "hello go" > data/single_file.txt
echo "123" > data/file1.hash
echo "456" > data/file2.hash
go run embeddemo.go
```
