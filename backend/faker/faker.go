package faker

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Faker interface {
	Fake(locale string, rules map[string]interface{}) string
}

var localizedFirstNames = map[string][]string{
	"en": {"Alice", "Bob", "Charlie", "Diane", "Edward", "Fiona", "George", "Hannah", "Ian", "Julia", "Kevin", "Laura", "Michael", "Nina", "Oscar", "Paula", "Quentin", "Rachel", "Steve", "Tina", "Uma", "Victor", "Wendy", "Xavier", "Yvonne"},
	"fr": {"Alice", "Bob", "Charlie", "Diane", "Émile", "François", "Giselle", "Hugo", "Inès", "Julien", "Karim", "Léa", "Mathieu", "Nathalie", "Olivier", "Pauline", "Quentin", "Romain", "Sophie", "Théo", "Ulysse", "Valérie", "William", "Xavier", "Yasmine"},
}
var localizedLastNames = map[string][]string{
	"en": {"Smith", "Johnson", "Williams", "Brown", "Jones", "Miller", "Davis", "Garcia", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin", "Lee", "Perez", "Thompson", "White", "Harris"},
	"fr": {"Martin", "Dupont", "Lemoine", "Bernard", "Dubois", "Thomas", "Robert", "Richard", "Petit", "Durand", "Leroy", "Moreau", "Simon", "Laurent", "Lefebvre", "Michel", "Garcia", "David", "Bertrand", "Roux", "Vincent", "Fournier", "Morel", "Girard", "Andre"},
}
var localizedDomains = map[string][]string{
	"en": {"example.com", "mail.com", "test.org", "demo.net", "sample.io", "alpha.com", "beta.org", "gamma.net", "delta.com", "epsilon.org", "zeta.net", "theta.com", "iota.org", "kappa.net", "lambda.com", "mu.org", "nu.net", "omicron.com", "pi.org", "rho.net", "sigma.com", "tau.org", "upsilon.net", "phi.com", "chi.org"},
	"fr": {"exemple.fr", "courriel.fr", "test.fr", "domaine.fr", "mail.fr", "alpha.fr", "beta.fr", "gamma.fr", "delta.fr", "epsilon.fr", "zeta.fr", "theta.fr", "iota.fr", "kappa.fr", "lambda.fr", "mu.fr", "nu.fr", "omicron.fr", "pi.fr", "rho.fr", "sigma.fr", "tau.fr", "upsilon.fr", "phi.fr", "chi.fr"},
}
var localizedStreets = map[string][]string{
	"en": {"Main St", "High St", "Broadway", "Elm St", "Maple Ave", "Oak St", "Pine St", "Cedar Ave", "Walnut St", "Chestnut St", "Spruce St", "Willow Ave", "Birch St", "Ash St", "Cherry Ln", "Poplar St", "Sycamore Ave", "Beech St", "Magnolia Blvd", "Hickory St", "Dogwood Dr", "Cottonwood St", "Palm Ave", "Redwood St", "Sequoia Ave"},
	"fr": {"rue de Paris", "avenue de Lyon", "boulevard Haussmann", "rue Victor Hugo", "rue de la Paix", "avenue des Champs-Élysées", "rue de Rivoli", "rue du Faubourg Saint-Honoré", "rue de la République", "rue Nationale", "rue de la Gare", "rue du Général de Gaulle", "rue de l'Église", "rue des Écoles", "rue du Moulin", "rue du Château", "rue des Lilas", "rue des Acacias", "rue des Fleurs", "rue des Jardins", "rue des Peupliers", "rue des Platanes", "rue des Tilleuls", "rue des Cerisiers", "rue des Rosiers"},
}
var localizedCities = map[string][]string{
	"en": {"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose", "Austin", "Jacksonville", "Fort Worth", "Columbus", "Charlotte", "San Francisco", "Indianapolis", "Seattle", "Denver", "Washington", "Boston", "El Paso", "Nashville", "Detroit", "Oklahoma City"},
	"fr": {"Paris", "Marseille", "Lyon", "Toulouse", "Nice", "Nantes", "Strasbourg", "Montpellier", "Bordeaux", "Lille", "Rennes", "Reims", "Le Havre", "Saint-Étienne", "Toulon", "Grenoble", "Dijon", "Angers", "Nîmes", "Villeurbanne", "Saint-Denis", "Aix-en-Provence", "Clermont-Ferrand", "Brest", "Limoges", "Tours"},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type NameFaker struct{}
func (n NameFaker) Fake(locale string, rules map[string]interface{}) string {
	if rules != nil {
		if custom, ok := rules["name"].(string); ok && custom != "" {
			return custom
		}
	}
	firsts := localizedFirstNames[locale]
	lasts := localizedLastNames[locale]
	if firsts == nil { firsts = localizedFirstNames["en"] }
	if lasts == nil { lasts = localizedLastNames["en"] }
	return firsts[rand.Intn(len(firsts))] + " " + lasts[rand.Intn(len(lasts))]
}

type EmailFaker struct{}
func (e EmailFaker) Fake(locale string, rules map[string]interface{}) string {
	var name string
	if rules != nil {
		if n, ok := rules["name"].(string); ok && n != "" {
			name = n
		}
	}
	if name == "" {
		name = NameFaker{}.Fake(locale, nil)
	}
	domain := ""
	if rules != nil {
		if d, ok := rules["domain"].(string); ok && d != "" {
			domain = d
		}
	}
	domains := localizedDomains[locale]
	if domains == nil { domains = localizedDomains["en"] }
	if domain == "" {
		domain = domains[rand.Intn(len(domains))]
	}
	namePart := name
	namePart = strings.ToLower(namePart)
	namePart = strings.ReplaceAll(namePart, " ", ".")
	return namePart + "@" + domain
}

type PhoneFaker struct{}

func (p PhoneFaker) Fake(locale string, rules map[string]interface{}) string {
	switch locale {
	case "fr":
		// France : 06 ou 07
		prefix := "06"
		if rules != nil {
			if v, ok := rules["prefix"].(string); ok && v != "" {
				prefix = v
			}
		}
		num := prefix
		for i := 0; i < 8; i++ {
			num += strconv.Itoa(rand.Intn(10))
		}
		return num[:2] + " " + num[2:4] + " " + num[4:6] + " " + num[6:8] + " " + num[8:10]

	case "en":
		// US : +1 (XXX) XXX-XXXX
		area := rand.Intn(800) + 200
		if rules != nil {
			if a, ok := rules["area"].(string); ok && a != "" {
				if areaInt, err := strconv.Atoi(a); err == nil {
					area = areaInt
				}
			}
		}
		central := rand.Intn(900) + 100
		station := rand.Intn(10000)
		return fmt.Sprintf("+1 (%03d) %03d-%04d", area, central, station)

	default:
		num := "+"
		for i := 0; i < 10; i++ {
			num += strconv.Itoa(rand.Intn(10))
		}
		return num
	}
}



type AddressFaker struct{}
func (a AddressFaker) Fake(locale string, rules map[string]interface{}) string {
	streets := localizedStreets[locale]
	if streets == nil { streets = localizedStreets["en"] }
	return fmt.Sprintf("%d %s", rand.Intn(200)+1, streets[rand.Intn(len(streets))])
}

type CityFaker struct{}
func (c CityFaker) Fake(locale string, rules map[string]interface{}) string {
	if rules != nil {
		if city, ok := rules["city"].(string); ok && city != "" {
			return city
		}
	}
	cities := localizedCities[locale]
	if cities == nil { cities = localizedCities["en"] }
	return cities[rand.Intn(len(cities))]
}

type AgeFaker struct{}
func (a AgeFaker) Fake(locale string, rules map[string]interface{}) string {
	min := 18
	max := 99
	if v, ok := rules["min"].(int); ok {
		min = v
	}
	if v, ok := rules["max"].(int); ok {
		max = v
	}
	if min > max {
		min, max = max, min
	}
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

var Fakers = map[string]Faker{
	"name":    NameFaker{},
	"email":   EmailFaker{},
	"phone":   PhoneFaker{},
	"address": AddressFaker{},
	"city":    CityFaker{},
	"age":     AgeFaker{},
	"postalcode": PostalCodeFaker{},
}

type PostalCodeFaker struct{}
var cityToPostalPrefix = map[string]string{
	// France
	"Paris":       "75",
	"Marseille":   "13",
	"Lyon":        "69",
	"Toulouse":    "31",
	"Nice":        "06",
	"Nantes":      "44",
	// US
	"New York":    "100",
	"Los Angeles": "900",
	"Chicago":     "606",
	"Houston":     "770",
	"Phoenix":     "850",
}

func (p PostalCodeFaker) Fake(locale string, rules map[string]interface{}) string {
	city := ""
	if rules != nil {
		if c, ok := rules["city"].(string); ok {
			city = c
		}
	}

	prefix := ""
	if city != "" {
		if pre, ok := cityToPostalPrefix[city]; ok {
			prefix = pre
		}
	}

	if prefix == "" {
		if locale == "fr" {
			prefix = strconv.Itoa(rand.Intn(95) + 1)
			if len(prefix) == 1 {
				prefix = "0" + prefix
			}
		} else if locale == "en" {
			prefix = strconv.Itoa(rand.Intn(899) + 100)
		} else {
			prefix = "00"
		}
	}

	suffix := ""
	if locale == "fr" {
		suffix = fmt.Sprintf("%03d", rand.Intn(1000))
	} else if locale == "en" {
		suffix = fmt.Sprintf("%02d", rand.Intn(100))
	}

	return prefix + suffix
}
