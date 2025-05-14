package proto

//go:generate buf format -w
//go:generate buf generate
//go:generate buf lint
//go:generate buf generate --template buf.gen.doc.contact_manager.yaml --path contact_manager/v1/contact_manager.proto
