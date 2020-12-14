package main

import (
	"context"
	"github.com/jmwri/luckydice/application"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Panicf("unable to create logger: %s", err)
	}
	defer logger.Sync()
	// Get the bot token from environment
	token := os.Getenv("DISCORD_TOKEN")
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Fatal("error creating discord session", zap.Error(err))
		return
	}

	inputParser := application.NewInputParser()
	roller := application.NewRoller()
	outputBuilder := application.NewOutputBuilder()
	handler := application.NewHandler(logger, inputParser, roller, outputBuilder)

	// Register the handler func as a callback for MessageCreate events.
	dg.AddHandler(handler.Handle)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		logger.Error("error connecting to discord", zap.Error(err))
		return
	}

	ctx := context.Background()
	guildReporter := application.NewGuildReporter(logger, dg, time.Minute*30)
	go guildReporter.Start(ctx)

	// Wait here until CTRL-C or other term signal is received.
	logger.Info("bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	ctx.Done()

	// Cleanly close down the Discord session.
	dg.Close()
}
