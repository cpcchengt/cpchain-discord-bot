package airdrop

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func AirdropApply(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	/*
		如果用户私聊你'!airdrop'
		如果是在公告频道则返回提示信息
	*/
	if m.Content == "!airdrop" {
		fmt.Println(m.Content, m.ChannelID, m.Author, m.Type, m.Components, m.Application, m.Embeds, m.GuildID, m.Mentions)

		// private chat
		if m.GuildID == "" {
			applyReply(s, m)
		} else {
			s.ChannelMessageSend(m.ChannelID, "请私信机器人领取空投")
		}
	}
}

func applyReply(s *discordgo.Session, m *discordgo.MessageCreate) {
	token := uuid.New()
	s.ChannelMessageSend(m.ChannelID, token.String())
}
