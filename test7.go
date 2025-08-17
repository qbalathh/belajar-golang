package main

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"time"
)

type Sample struct {
	Name string `required: "true" max: "10`
}

type User struct {
	Name string
	Age  int
}

type Userslice []User

func (userSlice Userslice) Len() int {
	return len(userSlice)
}

func (userSlice Userslice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice Userslice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"Muhammad Iqbal Athhar", 20},
		{"Joko Widodo", 62},
		{"Prabowo Subianto", 72},
	}

	sort.Sort(Userslice(users))

	fmt.Println(users)

	now := time.Now()
	fmt.Println("Current time:", now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Location:", now.Location())
	fmt.Println("Unix timestamp:", now.Unix())
	fmt.Println("Formatted time:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("Is zero time:", now.IsZero())
	fmt.Println("Time after 1 hour:", now.Add(time.Hour))
	fmt.Println("Time before 1 hour:", now.Add(-time.Hour))
	fmt.Println("Time difference:", now.Sub(now.Add(time.Hour)))
	fmt.Println("Time after 1 day:", now.AddDate(0, 0, 1))
	fmt.Println("Time before 1 day:", now.AddDate(0, 0, -1))
	fmt.Println("Time after 1 month:", now.AddDate(0, 1, 0))

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2023-10-01 12:00:00")
	fmt.Println("Parsed time:", parse)

	sample := Sample{"Katino"}
	sampleType := reflect.TypeOf(sample)
	fmt.Println("Sample type:", sampleType)
	fmt.Println("Sample name:", sampleType.Name())
	fmt.Println("Sample kind:", sampleType.Kind())
	fmt.Println("Sample field count:", sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println("Sample field tag:", sampleType.Field(0).Tag)

	fmt.Println("Is sample valid:", IsValid(sample))
	fmt.Println("Is empty sample valid:", IsValid(Sample{}))

	var nameRegex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println("Regex match:", nameRegex.MatchString("eao"))
	fmt.Println("Regex find all:", nameRegex.FindAllString("eao ebo eco edo efo", -1))
	fmt.Println("Regex find all submatches:", nameRegex.FindAllStringSubmatch("emo eko eio eyo ezo", -1))
	fmt.Println("Regex find all submatch indices:", nameRegex.FindAllStringSubmatchIndex("exo epo evo euo", -1))
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}
