package faker

import (
    "math/rand"
    "strconv"
    "time"
)

var firstNames = []string{"Alice", "Bob", "Charlie", "Diane"}
var lastNames = []string{"Martin", "Dupont", "Lemoine", "Bernard"}
var domains = []string{"example.com", "mail.com", "test.org"}
var streets = []string{"rue de Paris", "avenue de Lyon", "boulevard Haussmann"}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func FakeName() string {
    return firstNames[rand.Intn(len(firstNames))] + " " + lastNames[rand.Intn(len(lastNames))]
}

func FakeEmail() string {
    return FakeName() + "@" + domains[rand.Intn(len(domains))]
}

func FakePhone() string {
    return "06" + strconv.Itoa(rand.Intn(90000000)+10000000)
}

func FakeAddress() string {
    return strconv.Itoa(rand.Intn(200)+1) + " " + streets[rand.Intn(len(streets))]
}
