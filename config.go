//Copyright 2010 Cory Kolbeck <ckolbeck@gmail.com>.
//So long as this notice remains in place, you are welcome 
//to do whatever you like to or with this code.  This code is 
//provided 'As-Is' with no warrenty expressed or implied. 
//If you like it, and we happen to meet, buy me a beer sometime

package main

import (
	"cbeck/ircbot"
	"json"
	"io/ioutil"
	"os"
)

type botConfig struct {
	BotName string
	AttnChar byte
	IdentPW string

	Server string
	Port int

	Channels []string

	Version string
	SourceLoc string

	Help map[string]string
	Trusted map[string]bool
	TitleWhitelist map[string]bool
	Ignores map[string]bool
}

func parseConfig(filename string) (*botConfig, os.Error){
	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	c := &botConfig{}

	err = json.Unmarshal(raw, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func initParseConfig() {
	var err os.Error

	config, err = parseConfig(configPath)

	if err != nil {
		errors.Printf("Failed to parse config file: %s\n --exiting--", err.String())
		os.Exit(1)
	}

	if config.BotName == "" || config.Server == "" {
		errors.Printf("Malformed config file or missing item,\n%#v\n exiting", config)
		os.Exit(1)
	}

	helpList = "I understand the following commands: "

	for k := range config.Help {
		helpList += k + ", "
	}

	helpList += "source"	
}

func saveConfig(bc *botConfig) os.Error {
	_, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		return err
	}

	//mv old to path.bak
	//attempt to write
	//Fail: attempt to undo mv
	//success:

	return nil
}

func reparseConfig(bot *ircbot.Bot) bool {
	c, err := parseConfig(configPath)

	if err != nil {
		return false
	}

	err = bot.Send(c.Server, &ircbot.Message{
	Command : "NICK",
	Args : []string{c.BotName},
	})

	if err != nil {
		return false
	}

	if c.IdentPW != "" {
		bot.Send(c.Server, &ircbot.Message{
		Command : "PRIVMSG",
		Args : []string{"NickServ"},
		Trailing : "identify " + c.IdentPW,
		})
	}

	for _, channel := range c.Channels {
		bot.Send(c.Server, &ircbot.Message{
		Command : "JOIN",
		Args : []string{channel},
		})
	}

	helpList = "I understand the following commands: "

	for k := range c.Help {
		helpList += k + ", "
	}
	helpList += "source"	

	config = c
	return true
}
