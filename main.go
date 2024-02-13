package main

import (
	"dimploma/getflags"
	"fmt"
)

func main() {
	// todos.SendRequest()
	//check if q enabled (send request)
	sendRequest, err := getflags.CheckIfWebRequest()
	if sendRequest {
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}

	//send req with id
	//create dict
}
