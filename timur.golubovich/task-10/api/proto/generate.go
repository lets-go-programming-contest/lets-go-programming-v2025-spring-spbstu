package proto

//go:generate buf format -w
//go:generate buf generate
//go:generate buf lint
//go:generate buf generate --template buf.gen.doc.manager.yaml --path phonebook/v1/manager.proto
