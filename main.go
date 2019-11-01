package main

import (
	"log"
	"github.com/foomo/htpasswd"
	"os"
)

func main() {
	file := "/tmp/fancyindex.htpasswd"
	name := "admin"
	password := "admin"
	envName := os.Getenv("HTTP_USERNAME")
	envPasswd := os.Getenv("HTTP_PASSWD")
	if envName != "" {
		name = envName
	}
	if envPasswd != "" {
		password = envPasswd
	}
	os.Remove(file)
	if err := htpasswd.SetPassword(file, name, password, htpasswd.HashBCrypt); err == nil {
		log.Printf("[htpasswd] %s generated successfully", file)
	} else {
		log.Printf("[htpasswd] %s generation failed: %s", file, err)
	}
}