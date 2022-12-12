package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"reflect"
	"time"

	gocron "github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var MEMORY map[string][]string = make(map[string][]string)

func GetIfeces() (map[string][]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return map[string][]string{}, err
	}

	ifall := make(map[string][]string)
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return map[string][]string{}, err
		}
		var tmp []string
		for _, a := range addrs {
			tmp = append(tmp, a.String())
		}
		ifall[i.Name] = tmp
	}
	return ifall, nil
}

func formatAddr(addrs map[string][]string) string {
	buff := ""
	for ifcace, addr := range addrs {
		buff += fmt.Sprintf("%s:\n\t", ifcace)
		for _, a := range addr {
			buff += a + "\n\t"
		}
		buff += "\n"
	}
	return buff
}

func main() {
	token := flag.String("token", "", "TOKEN")
	myID := flag.Int64("id", 0, "telegram chat id")
	flag.Parse()

	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(1).Hour().Do(func() {
		addrs, err := GetIfeces()
		if err != nil {
			fmt.Println(err)
		}
		if reflect.DeepEqual(MEMORY, addrs) {
			msg := tgbotapi.NewMessage(*myID, formatAddr(addrs))
			bot.Send(msg)
			MEMORY = addrs
		}
	})
	scheduler.StartBlocking()
}
