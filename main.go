package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/foomo/htpasswd"
)

var Version = "dev"

var outfile string
var showVersion bool

func init() {
	flag.StringVar(&outfile, "out", "/etc/nginx/.htpasswd", "htpasswd output path")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Parse()
}

func main() {
	if showVersion {
		fmt.Println("Name:    htpasswd-init")
		fmt.Println("Version: " + Version)
		fmt.Println("Author:  荒野無燈")
		return
	}
	file := outfile
	name := "admin"
	password := "admin"
	envName := os.Getenv("HTTP_AUTH_USER")
	envPasswd := os.Getenv("HTTP_AUTH_PASSWD")
	if envName != "" {
		name = envName
	}
	if envPasswd != "" {
		password = envPasswd
	}
	if !strings.HasSuffix(file, "htpasswd") {
		file += ".htpasswd"
	}
	os.Remove(file)
	if err := htpasswd.SetPassword(file, name, password, htpasswd.HashBCrypt); err == nil {
		log.Printf("[htpasswd] %s generated successfully", file)
	} else {
		log.Printf("[htpasswd] %s generation failed: %s", file, err)
	}
}
