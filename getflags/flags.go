package getflags

import (
	"dimploma/todos"
	"errors"
	"flag"
	"strconv"
)

func CheckIfWebRequest() (bool, error) {
	var (
		request bool
		name    string
		surname string
		country string
	)
	flag.BoolVar(&request, "q", false, "provide q flag if you want to send requset")
	flag.StringVar(&name, "name", "", "provide name of person you searching info about")
	flag.StringVar(&surname, "surname", "", "provide surname of person you searching info about")
	flag.StringVar(&country, "country", "0", "provide country of person you searching info about")
	flag.Parse()

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
