package commands

import "github.com/bwmarrin/discordgo"

func Get() []discordgo.ApplicationCommand {
	return []discordgo.ApplicationCommand{
		{
			Name:        ADD_REPO,
			Description: "Adds a new repo to monitor PRs from",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "url",
					Description: "GitHub repo webhook",
					Required:    true,
				},
			},
		},
	}
}
