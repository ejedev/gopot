package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Port string `yaml:"listening_port"`
	Logs string `yaml:"log_folder"`
	Save bool   `yaml:"save_logs"`
}

func (auth *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, auth)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return auth
}

func main() {
	var config conf
	config.getConf()

	ssh.Handle(func(s ssh.Session) {
		log.Println(fmt.Sprintf("%s logged in as user %s", s.RemoteAddr().String(), s.User()))
		// Simulating an ubuntu server prompt with the chosen name.
		term := terminal.NewTerminal(s, fmt.Sprintf("[%s@ubuntu ~]$ ", s.User()))
		line := ""
		for {
			line, _ = term.ReadLine()
			// Probably should avoid reusing code for this and login.
			if config.Save {
				file, err := os.OpenFile(fmt.Sprintf("%sactions/%s.log", config.Logs, s.RemoteAddr().String()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
				if err != nil {
					log.Fatal(err)
				}
				log.SetOutput(file)
			}
			log.Println(fmt.Sprintf("[Action - %s] %s", s.RemoteAddr().String(), line))
			io.WriteString(s, fmt.Sprintf("Input: %s\n", line))
		}
	})

	log.Println(fmt.Sprintf("GoPot v0.1 starting on port %s.", config.Port))
	log.Fatal(ssh.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil, ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {

		if config.Save {
			file, err := os.OpenFile(fmt.Sprintf("%saccess.log", config.Logs), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}
			log.SetOutput(file)
		}

		// Set to return true. This can be modified to allow a config of different passwords or usernames.
		log.Println(fmt.Sprintf("%s authenticating with username %s and password %s.", ctx.RemoteAddr().String(), ctx.User(), pass))
		return true
	})))
}
