package main

func main() {
    a := 1
    b := 2
    if true {
        add(a, b)
    }
}

func add(a, b int) {
    println(a + b)
}
