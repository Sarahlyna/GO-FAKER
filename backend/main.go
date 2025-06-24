package main

import (
	"encoding/json"
	"go-faker/backend/faker"
	"net/http"
	"strconv"
)

type FakeUser struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

var currentIndex = 0

var firstNames = []string{"Alice", "Bob", "Charlie", "Diane"}
var lastNames = []string{"Martin", "Dupont", "Lemoine", "Bernard"}
var domains = []string{"example.com", "mail.com", "test.org"}
var streets = []string{"rue de Paris", "avenue de Lyon", "boulevard Haussmann"}

func fakeName() string {
	first := firstNames[currentIndex%len(firstNames)]
	last := lastNames[currentIndex%len(lastNames)]
	return first + " " + last
}

func fakeEmail(name string) string {
	domain := domains[currentIndex%len(domains)]
	return name + "@" + domain
}

func fakePhone() string {
	num := "06"
	for i := 0; i < 8; i++ {
		chiffre := ((currentIndex + i) % 9) + 1
		num += string(rune('0' + chiffre))
	}
	return num
}

func fakeAddress() string {
	numero := 1 + (currentIndex % 100)
	street := streets[currentIndex%len(streets)]
	return strconv.Itoa(numero) + " " + street
}

func handler(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("mode")

	var user FakeUser

	if mode == "random" {
		user = FakeUser{
			Name:    faker.FakeName(),
			Email:   faker.FakeEmail(),
			Phone:   faker.FakePhone(),
			Address: faker.FakeAddress(),
		}
	} else {
		name := fakeName()
		user = FakeUser{
			Name:    name,
			Email:   fakeEmail(name),
			Phone:   fakePhone(),
			Address: fakeAddress(),
		}
		currentIndex++
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)
	http.HandleFunc("/api/fake", handler)
	println("Serveur Go Faker sur http://localhost:9090")
	http.ListenAndServe(":8080", nil)
}
