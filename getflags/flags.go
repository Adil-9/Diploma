package getflags

import (
	"dimploma/todos"
	"dimploma/variables"
	"errors"
	"flag"
	"strconv"
	"strings"
)

func CheckIfWebRequest(request bool, name, surname, country string) (bool, error) {

	if !request {
		return request, nil
	}

	countryId, err := strconv.Atoi(country)
	if err != nil || countryId < 0 || countryId > 224 {
		return request, errors.New("incorrect country id")
	}

	if name == "" || surname == "" {
		return request, errors.New("name or surname is not provided")
	}
	todos.SendRequest(name, surname, country)
	return true, nil
}

func ParseAllFlags() variables.AllVariables {
	var AllFlags variables.AllVariables

	flag.BoolVar(&AllFlags.Request, "q", false, "provide q flag if you want to send requset")
	flag.StringVar(&AllFlags.Name, "name", "", "provide name of person you searching info about")
	flag.StringVar(&AllFlags.Surname, "surname", "", "provide surname of person you searching info about")
	flag.StringVar(&AllFlags.Country, "country", "0", "provide country of person you searching info about")

	var strs string
	flag.StringVar(&strs, "strings", "", "provide string to add strings")
	AllFlags.StringVar = splitStrings(strs)

	var ints string
	flag.StringVar(&ints, "integers", "", "provide integers to add integers")
	AllFlags.StringVar = splitStrings(ints)

	flag.BoolVar(&AllFlags.SpecialCharAll, "country", false, "provide -c for all special charachters")
	flag.Parse()

	return AllFlags
}

func splitStrings(str string) []string {
	splitted := strings.Split(str, ",")
	return splitted
}
