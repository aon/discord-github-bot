package commands

import "github.com/bwmarrin/discordgo"

type commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)

func GetHandlers() commandHandlers {
	return commandHandlers{
		ADD_REPO: addRepo,
	}
}

func addRepo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	url := i.ApplicationCommandData().Options[0].StringValue()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: url,
		},
	})
}
