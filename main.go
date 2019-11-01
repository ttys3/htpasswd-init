package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/foomo/htpasswd"
	"os"
	"strings"
)

var Version = "dev"

var outfile string
var showVersion bool

func init() {
	flag.StringVar(&outfile, "out", "/etc/nginx/fancyindex.htpasswd", "htpasswd output path")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Parse()
}

func main() {
	if showVersion {
		fmt.Println(Version)
		return
	}
	file := outfile
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