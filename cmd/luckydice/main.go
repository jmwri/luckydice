package main

import (
	"context"
	"fmt"
	"github.com/jmwri/luckydice/application"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Get the bot token from environment
	token := os.Getenv("DISCORD_TOKEN")
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating discord session: %s", err)
		return
	}

	inputParser := application.NewInputParser()
	roller := application.NewRoller()
	outputBuilder := application.NewOutputBuilder()
	handler := application.NewHandler(inputParser, roller, outputBuilder)

	// Register the handler func as a callback for MessageCreate events.
	dg.AddHandler(handler.Handle)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	ctx := context.Background()
	guildReporter := application.NewGuildReporter(dg, time.Hour)
	guildReporter.Start(ctx)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	ctx.Done()

	// Cleanly close down the Discord session.
	dg.Close()
}
