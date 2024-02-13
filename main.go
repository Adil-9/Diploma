package main

import (
	"dimploma/getflags"
	"fmt"
)

func main() {
	AllFlags := getflags.ParseAllFlags()
	sendRequest, err := getflags.CheckIfWebRequest(AllFlags.Request, AllFlags.Name, AllFlags.Surname, AllFlags.Country)
	if sendRequest {
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
	//create a dictionary
}
