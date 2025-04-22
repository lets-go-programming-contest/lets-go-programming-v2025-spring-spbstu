package data

import _ "embed"

//go:embed hello.txt
var Hello []byte

//go:embed world.txt
var World []byte
