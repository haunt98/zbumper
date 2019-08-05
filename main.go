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
	//err := promptAction(bump)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//log.Print(bump)
}
