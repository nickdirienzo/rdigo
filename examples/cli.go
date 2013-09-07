package main

import (
	"fmt"
	"github.com/nickdirienzo/rdigo"
	"log"
)

func main() {
	rdio := rdigo.NewClient("CONSUMER_TOKEN", "CONSUMER_SECRET")
	requestToken, url, err := rdio.BeginAuthentication("oob")
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(requestToken, url)
	verificationCode := ""
	fmt.Scanln(&verificationCode)

	err = rdio.CompleteAuthentication(requestToken, verificationCode)
	if err != nil {
		log.Println(err.Error())
		return
	}
	values := make(map[string]string)
	values["keys"] = "a184236,a254895,a242205"
	resp, err := rdio.Call("get", values)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(resp)
}
