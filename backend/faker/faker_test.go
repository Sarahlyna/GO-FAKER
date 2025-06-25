package faker

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestNameFaker(t *testing.T) {
	name := NameFaker{}.Fake("en", nil)
	if len(strings.Split(name, " ")) != 2 {
		t.Errorf("Expected name to have two parts, got: %s", name)
	}
}

func TestEmailFaker(t *testing.T) {
	name := "Alice Smith"
	email := EmailFaker{}.Fake("en", map[string]interface{}{ "name": name })
	if !strings.HasPrefix(email, "alice.smith@") {
		t.Errorf("Expected email to start with alice.smith@, got: %s", email)
	}
	if !strings.Contains(email, "@") {
		t.Errorf("Expected email to contain @, got: %s", email)
	}
}

func TestPhoneFaker(t *testing.T) {
	phone := PhoneFaker{}.Fake("fr", map[string]interface{}{ "prefix": "07", "length": 9 })
	if !strings.HasPrefix(phone, "07") {
		t.Errorf("Expected phone to start with 07, got: %s", phone)
	}
	if len(phone) != 11 {
		t.Errorf("Expected phone length 11, got: %d (%s)", len(phone), phone)
	}
}

func TestAddressFaker(t *testing.T) {
	address := AddressFaker{}.Fake("fr", nil)
	if !regexp.MustCompile(`^\d+ `).MatchString(address) {
		t.Errorf("Expected address to start with a number, got: %s", address)
	}
}

func TestCityFaker(t *testing.T) {
	city := CityFaker{}.Fake("fr", nil)
	if city == "" {
		t.Error("Expected non-empty city")
	}
}

func TestAgeFakerDefault(t *testing.T) {
	ageStr := AgeFaker{}.Fake("en", nil)
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		t.Fatalf("Age is not a number: %s", ageStr)
	}
	if age < 18 || age > 99 {
		t.Errorf("Default age out of range: %d", age)
	}
}

func TestAgeFakerMinMax(t *testing.T) {
	min, max := 25, 30
	for i := 0; i < 20; i++ {
		ageStr := AgeFaker{}.Fake("en", map[string]interface{}{ "min": min, "max": max })
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			t.Fatalf("Age is not a number: %s", ageStr)
		}
		if age < min || age > max {
			t.Errorf("Age out of range: %d (expected %d-%d)", age, min, max)
		}
	}
} 