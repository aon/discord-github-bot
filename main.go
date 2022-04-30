package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/discord-github-bot/bot"
)

var Token string
var ChannelId string

func init() {
	flag.StringVar(&Token, "t", "", "Discord bot token")
	flag.StringVar(&ChannelId, "c", "", "Discord channel ID")
	flag.Parse()

	if Token == "" || ChannelId == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags)

	bot, _ := bot.NewBot(l, Token, ChannelId)
	bot.Start()

	// Wait here until CTRL-C or other term signal is received.
	l.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	// Block until a signal is received
	<-sc
	l.Println("Received terminate, gracefully shutting down.")

	// Cleanly close down the Discord bot.
	bot.Close()
}
