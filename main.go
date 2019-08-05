package main

import "log"

var bump = &Bump{}

func main() {
	accessToken, err := ensureAccessTokenAvailableAndValid()
	for err != nil {
		err = promptAccessToken()
		if err != nil {
			log.Println(err.Error())
			return
		}
		accessToken, err = ensureAccessTokenAvailableAndValid()
	}
	err = promptAction(bump, accessToken)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Yayyy!")
}
