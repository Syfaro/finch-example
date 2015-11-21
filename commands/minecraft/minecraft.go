package finchcommandminecraft

import (
	"fmt"
	"github.com/syfaro/finch"
	"github.com/syfaro/mcapi/client"
	"github.com/syfaro/telegram-bot-api"
	"strings"
)

func init() {
	finch.RegisterCommand(&minecraftCommand{})
}

type minecraftCommand struct {
	finch.CommandBase
}

func (cmd *minecraftCommand) Help() finch.Help {
	return finch.Help{
		Name:        "Minecraft Server Status",
		Description: "Gets information about a Minecraft server",
		Example:     "/mc@@",
	}
}

func (cmd *minecraftCommand) ShouldExecute(update tgbotapi.Update) bool {
	return finch.SimpleCommand("mc", update.Message.Text)
}

func (cmd *minecraftCommand) Execute(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter your server IP")
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = tgbotapi.ForceReply{
		ForceReply: true,
		Selective:  true,
	}

	cmd.MyState.WaitingForReply = true

	return cmd.Finch.SendMessage(msg)
}

func (cmd *minecraftCommand) ExecuteKeyboard(update tgbotapi.Update) error {
	cmd.MyState.WaitingForReply = false

	ip := strings.Trim(update.Message.Text, " ")

	status, _ := mcapi.GetServerStatus(ip, 25565)

	var text string
	if status.Online {
		text = fmt.Sprintf("Server is online! There are %d players online", status.Players.Now)
	} else {
		text = "Server is offline"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	return cmd.Finch.SendMessage(msg)
}
