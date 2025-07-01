package main

import (
	"encoding/json"
	"go-faker/backend/faker"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

type FakeUser struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	locale := r.URL.Query().Get("locale")
	if locale == "" {
		locale = "fr"
	}

	phoneRules := make(map[string]interface{})
	if prefix := r.URL.Query().Get("phone_prefix"); prefix != "" {
		phoneRules["prefix"] = prefix
	}
	if lengthStr := r.URL.Query().Get("phone_length"); lengthStr != "" {
		if length, err := strconv.Atoi(lengthStr); err == nil {
			phoneRules["length"] = length
		}
	}

	name := faker.Fakers["name"].Fake(locale, nil)
	user := FakeUser{
		Name:    name,
		Email:   faker.Fakers["email"].Fake(locale, map[string]interface{}{"name": name}),
		Phone:   faker.Fakers["phone"].Fake(locale, phoneRules),
		Address: faker.Fakers["address"].Fake(locale, nil),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func generateFromSchema(schema map[string]interface{}, locale string, context map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, val := range schema {
		if key == "name" {
			if v, ok := val.(string); ok && v == "name" {
				name := faker.Fakers["name"].Fake(locale, nil)
				result[key] = name
				context["name"] = name
			} else if v, ok := val.(map[string]interface{}); ok && v["faker"] == "name" {
				name := faker.Fakers["name"].Fake(locale, deepCopyMap(v))
				result[key] = name
				context["name"] = name
			}
		}
		if key == "city" {
			if v, ok := val.(string); ok && v == "city" {
				city := faker.Fakers["city"].Fake(locale, nil)
				result[key] = city
				context["city"] = city
			} else if v, ok := val.(map[string]interface{}); ok && v["faker"] == "city" {
				city := faker.Fakers["city"].Fake(locale, deepCopyMap(v))
				result[key] = city
				context["city"] = city
			}
		}
	}

	for key, val := range schema {
		if key == "name" || key == "city" {
			continue
		}
		switch v := val.(type) {
		case string:
			if v == "email" && context["name"] != nil {
				result[key] = faker.Fakers["email"].Fake(locale, map[string]interface{}{"name": context["name"].(string)})
			} else if v == "age" || regexp.MustCompile(`^age(:|$)`).MatchString(v) {
				rules := make(map[string]interface{})
				if v != "age" {
					re := regexp.MustCompile(`min=(\d+)`)
					if m := re.FindStringSubmatch(v); len(m) == 2 {
						minVal, _ := strconv.Atoi(m[1])
						rules["min"] = minVal
					}
					re = regexp.MustCompile(`max=(\d+)`)
					if m := re.FindStringSubmatch(v); len(m) == 2 {
						maxVal, _ := strconv.Atoi(m[1])
						rules["max"] = maxVal
					}
				}
				result[key] = faker.Fakers["age"].Fake(locale, rules)
			} else if fakerFn, ok := faker.Fakers[v]; ok {
				result[key] = fakerFn.Fake(locale, nil)
			} else {
				result[key] = v
			}
		case map[string]interface{}:
			if f, ok := v["faker"]; ok {
				fakerName, _ := f.(string)
				opts := deepCopyMap(v)

				if fakerName == "email" && context["name"] != nil {
					if _, hasName := opts["name"]; !hasName {
						opts["name"] = context["name"].(string)
					}
				}
				if fakerName == "postalcode" && context["city"] != nil {
					if _, hasCity := opts["city"]; !hasCity {
						opts["city"] = context["city"].(string)
					}
				}

				if fakerFn, ok := faker.Fakers[fakerName]; ok {
					result[key] = fakerFn.Fake(locale, opts)
				} else {
					result[key] = opts
				}
			} else {
				result[key] = generateFromSchema(v, locale, context)
			}
		default:
			result[key] = v
		}
	}
	return result
}

func deepCopyMap(m map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{}, len(m))
	for k, v := range m {
		if sub, ok := v.(map[string]interface{}); ok {
			newMap[k] = deepCopyMap(sub)
		} else {
			newMap[k] = v
		}
	}
	return newMap
}

func schemaHandler(w http.ResponseWriter, r *http.Request) {
	type SchemaRequest struct {
		Schema map[string]interface{} `json:"schema"`
		Count  int                    `json:"count"`
		Locale string                 `json:"locale"`
	}
	var req SchemaRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if req.Count <= 0 {
		req.Count = 1
	}
	if req.Locale == "" {
		req.Locale = "fr"
	}
	results := make([]map[string]interface{}, req.Count)
	for i := 0; i < req.Count; i++ {
		results[i] = generateFromSchema(req.Schema, req.Locale, make(map[string]interface{}))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func fakersListHandler(w http.ResponseWriter, r *http.Request) {
	fakersList := make([]string, 0, len(faker.Fakers))
	for k := range faker.Fakers {
		fakersList = append(fakersList, k)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fakersList)
}

func main() {
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)
	http.HandleFunc("/api/fake", handler)
	http.HandleFunc("/api/fake/schema", schemaHandler)
	http.HandleFunc("/api/fakers", fakersListHandler)
	println("Serveur Go Faker sur http://localhost:9090")
	http.ListenAndServe(":8080", nil)
}
