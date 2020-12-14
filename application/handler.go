package application

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jmwri/luckydice/domain"
	"go.uber.org/zap"
	"strings"
)

const messagePrefix = "!roll"

func NewHandler(logger *zap.Logger, inputParser domain.InputParser, roller domain.Roller, outputBuilder domain.OutputBuilder, recorder domain.InputRecorder) *Handler {
	return &Handler{
		logger:        logger,
		inputParser:   inputParser,
		roller:        roller,
		outputBuilder: outputBuilder,
		recorder:      recorder,
	}
}

type Handler struct {
	logger        *zap.Logger
	inputParser   domain.InputParser
	roller        domain.Roller
	outputBuilder domain.OutputBuilder
	recorder      domain.InputRecorder
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
		h.recorder.RecordHelp()
		response = h.buildHelp(m)
	} else {
		response = h.parseRoll(content, m)
	}

	h.logger.Info(
		"responded to message",
		zap.String("request", content),
		zap.String("response", response),
	)

	_, err := s.ChannelMessageSend(m.ChannelID, response)
	if err != nil {
		h.logger.Error("failed to send message", zap.Error(err))
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
		h.recorder.RecordMisunderstanding()
		return fmt.Sprintf("Sorry %s, I don't understand. You can ask me for help with `%s help`.", m.Author.Mention(), messagePrefix)
	} else {
		h.recorder.RecordRoll(input)
		output := h.roller.Roll(input)
		return h.outputBuilder.Build(m.Author.Mention(), output)
	}
}
