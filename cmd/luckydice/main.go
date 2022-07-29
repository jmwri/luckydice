package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jmwri/luckydice/internal"
	"github.com/jmwri/luckydice/internal/core"
	"github.com/jmwri/luckydice/internal/domain"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Panicf("unable to create logger: %s", err)
	}
}

var dg *discordgo.Session

func init() {
	var err error
	// Get the bot token from environment
	token := os.Getenv("DISCORD_TOKEN")
	// Create a new Discord session using the provided bot token
	dg, err = discordgo.New("Bot " + token)
	if err != nil {
		logger.Fatal("error creating discord session", zap.Error(err))
		return
	}
}

var opts = domain.ServiceOpts{
	RollCmdName:          "roll",
	RollCmdInputName:     "input",
	RollUtilCmdName:      "roll-util",
	RollUtilHelpCmdName:  "help",
	RollUtilStatsCmdName: "stats",
	OldPrefix:            "!roll",
}
var svc internal.Service = core.NewService(opts)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        opts.RollCmdName,
		Description: "Roll dice",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        opts.RollCmdInputName,
				Description: "Dice roll input",
				Required:    true,
			},
		},
	},
	{
		Name:        opts.RollUtilCmdName,
		Description: "Give other info on the bot",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        opts.RollUtilHelpCmdName,
				Description: "Display help on how to use the bot",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        opts.RollUtilStatsCmdName,
				Description: "Display some statistics on the bot",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	},
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	opts.RollCmdName: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		content := ""
		var err error

		for _, opt := range options {
			switch opt.Name {
			case opts.RollCmdInputName:
				content, err = svc.HandleRoll(i.Member.Mention(), opt.StringValue())
				break
			default:
				content = fmt.Sprintf("Unknown command: %s", opt.Name)
				break
			}
		}

		if err != nil {
			logger.Error("failed to handle cmd", zap.String("command", opts.RollCmdName), zap.Error(err))
			content = "Oops, something went wrong!"
		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
		if err != nil {
			logger.Error("failed to respond to interaction", zap.Error(err))
		}
	},
	opts.RollUtilCmdName: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		content := ""
		var err error

		for _, opt := range options {
			switch opt.Name {
			case opts.RollUtilHelpCmdName:
				content, err = svc.HandleHelp(i.Member.Mention())
				break
			case opts.RollUtilStatsCmdName:
				content, err = svc.HandleStats(i.Member.Mention())
				break
			default:
				content = fmt.Sprintf("Unknown command: %s", opt.Name)
				break
			}
		}

		if err != nil {
			logger.Error("failed to handle cmd", zap.String("command", opts.RollUtilCmdName), zap.Error(err))
			content = "Oops, something went wrong!"
		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
		if err != nil {
			logger.Error("failed to respond to interaction", zap.Error(err))
		}
	},
}

func init() {
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handler, ok := commandHandlers[i.ApplicationCommandData().Name]
		if !ok {
			// No handler found
			return
		}
		// Run the handler
		handler(s, i)
	})
}

func main() {
	// Always try to sync to logger on exit
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	// Log when logged in
	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	// Connect to discord
	err := dg.Open()
	if err != nil {
		logger.Error("error connecting to discord", zap.Error(err))
		return
	}

	logger.Info("registering commands")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, cmd := range commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", cmd)
		if err != nil {
			logger.Error("failed to register command", zap.Any("command", cmd), zap.Error(err))
			continue
		}
		logger.Info("registered command", zap.String("command", cmd.Name))
		registeredCommands[i] = cmd
	}

	// Close discord session on exit
	defer func(dg *discordgo.Session) {
		_ = dg.Close()
	}(dg)

	// Wait for stop signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	log.Println("Press Ctrl+C to exit")
	<-stop

	// Unregister commands on shutdown
	log.Println("Removing commands...")
	for _, cmd := range registeredCommands {
		err := dg.ApplicationCommandDelete(dg.State.User.ID, "", cmd.ID)
		if err != nil {
			logger.Error("failed to unregister command", zap.String("command", cmd.Name), zap.Error(err))
			continue
		}
		logger.Info("unregistered command", zap.String("command", cmd.Name))
	}
}
