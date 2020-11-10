package application

import (
	"fmt"
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

	var response string
	if content == "help" {
		response = h.buildHelp(m)
	} else {
		response = h.parseRoll(content, m)
	}

	log.Println(content, response)

	_, err := s.ChannelMessageSend(m.ChannelID, response)
	if err != nil {
		log.Println("failed to send msg")
	}
}

func (h *Handler) buildHelp(m *discordgo.MessageCreate) string {
	lines := []string{
		fmt.Sprintf("Hi %s! You can use me by typing the following: `%s {number of rolls} d{sides on die} {modifier}`. For example: `%s 2 d20 +3`.", m.Author.Mention(), messagePrefix, messagePrefix),
		fmt.Sprintf("All whitespace is optional, and you can exclude {number of rolls} and {modifier}."),
		fmt.Sprintf("`%s 1d20+0` is the same as `%s d20`.", messagePrefix, messagePrefix),
	}
	return strings.Join(lines, "\n")
}

func (h *Handler) parseRoll(content string, m *discordgo.MessageCreate) string {
	input, err := h.inputParser.Parse(content)
	if err != nil {
		return fmt.Sprintf("Sorry %s, I don't understand. You can ask me for help with `%s help`.", m.Author.Mention(), messagePrefix)
	} else {
		output := h.roller.Roll(input)
		return h.outputBuilder.Build(m.Author.Mention(), output)
	}
}
