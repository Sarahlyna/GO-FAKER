package main

import (
	"encoding/json"
	"net/http"
	"go-faker/backend/faker"
)

type FakeData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*") 
    w.Header().Set("Content-Type", "application/json")

    data := FakeData{
        Name:    faker.FakeName(),
        Email:   faker.FakeEmail(),
        Phone:   faker.FakePhone(),
        Address: faker.FakeAddress(),
    }

    json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/api/fake", apiHandler)

	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}
