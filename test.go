package main

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	var pswd = "123456"

	hash, err := bcrypt.GenerateFromPassword([]byte(pswd), 4); if err != nil {
		log.Println("Failed to Hash Password!")
	}
	var pswdHashed = string(hash)


	if err := bcrypt.CompareHashAndPassword([]byte(pswdHashed), []byte(pswd)); err != nil {
		log.Println("Failed to Valid Password!")
	}
	log.Println("", err == nil)

	
	log.Println("", string(pswdHashed) == pswdHashed)
}