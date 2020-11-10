package application

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jmwri/luckydice/domain"
	"log"
	"strings"
)

const messagePrefix = "!roll"

func NewHandler(inputParser domain.InputParser, roller domain.Roller, outputBuilder domain.OutputBuilder) *Handler {
	return &Handler{
		inputParser:   inputParser,
		roller:        roller,
		outputBuilder: outputBuilder,
	}
}

type Handler struct {
	inputParser   domain.InputParser
	roller        domain.Roller
	outputBuilder domain.OutputBuilder
}

func (h *Handler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.ToLower(m.Content)

	if !strings.HasPrefix(content, messagePrefix) {
		return
	}
	content = strings.TrimPrefix(content, messagePrefix)
	content = strings.TrimSpace(content)

	input, err := h.inputParser.Parse(content)
	if err != nil {
		return
	}
	output := h.roller.Roll(input)
	response := h.outputBuilder.Build(m.Author.Mention(), output)
	log.Println(content, response)

	_, err = s.ChannelMessageSend(m.ChannelID, response)
	if err != nil {
		log.Println("failed to send msg")
	}
}
