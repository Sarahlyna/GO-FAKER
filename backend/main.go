package main

var firstNames = [4]string{"Alice", "Bob", "Charlie", "Diane"}
var lastNames = [4]string{"Martin", "Dupont", "Lemoine", "Bernard"}
var domains = [3]string{"example.com", "mail.com", "test.org"}
var streets = [3]string{"rue de Paris", "avenue de Lyon", "boulevard Haussmann"}

var currentIndex = 0

func fakeName() string {
	first := firstNames[currentIndex%len(firstNames)]
	last := lastNames[currentIndex%len(lastNames)]
	return first + " " + last
}

func fakeEmail(name string) string {
	domain := domains[currentIndex%len(domains)]
	email := name + "@" + domain
	return email
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
	return itoa(numero) + " " + street
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var digits = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var buf [10]byte
	i := 10
	for n > 0 {
		i--
		buf[i] = digits[n%10]
		n /= 10
	}
	return string(buf[i:])
}

func main() {
	println("=== Faux utilisateur ===")
	name := fakeName()
	println("Nom: " + name)
	println("Email: " + fakeEmail(name))
	println("Téléphone: " + fakePhone())
	println("Adresse: " + fakeAddress())

	currentIndex++
}
