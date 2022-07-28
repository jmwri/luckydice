package main

import (
	"context"
	"fmt"
	"github.com/jmwri/luckydice/internal"
	"github.com/jmwri/luckydice/internal/adapter"
	"github.com/jmwri/luckydice/internal/core"
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

	var svc internal.Service = core.NewService("!roll")
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		err := svc.Handle(m.Author.Mention(), m.Content, func(response string) error {
			_, err := s.ChannelMessageSend(m.ChannelID, response)
			if err != nil {
				return fmt.Errorf("failed to send message: %w", err)
			}
			logger.Info(
				"handled message",
				zap.String("request", m.Content),
				zap.String("response", response),
			)
			return nil
		})

		if err != nil {
			logger.Error(
				"failed to handle message",
				zap.Error(err),
			)
		}
	})

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		logger.Error("error connecting to discord", zap.Error(err))
		return
	}

	ctx := context.Background()
	guildCountProvider := adapter.NewGuildCountProvider(dg)
	periodicReporter := core.NewPeriodicReporter(logger, time.Minute*30, guildCountProvider, svc)
	go periodicReporter.Start(ctx)

	// Wait here until CTRL-C or other term signal is received.
	logger.Info("bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	ctx.Done()

	// Cleanly close down the Discord session.
	_ = dg.Close()
}
