package main

import (
    "fmt"
    "github.com/tonnom/faker/faker"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go [name|email|phone|address]")
        return
    }

    switch os.Args[1] {
    case "name":
        fmt.Println(faker.FakeName())
    case "email":
        fmt.Println(faker.FakeEmail())
    case "phone":
        fmt.Println(faker.FakePhone())
    case "address":
        fmt.Println(faker.FakeAddress())
    default:
        fmt.Println("Commande inconnue.")
    }
}
