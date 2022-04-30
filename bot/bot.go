package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/discord-github-bot/bot/commands"
)

type Bot struct {
	l    *log.Logger
	dg   *discordgo.Session
	chId string
}

func NewBot(l *log.Logger, token string, chId string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return &Bot{}, err
	}

	return &Bot{l, dg, chId}, nil
}

func (b *Bot) Start() {
	b.l.Println("Starting bot...")
	err := b.dg.Open()
	if err != nil {
		b.l.Fatal("Cannot open the session: ", err)
	}

	b.l.Println("Adding commands...")
	b.registerCommands()
}

func (b *Bot) SendMsg(m string) error {
	b.l.Println("Sending message: ", m)
	_, err := b.dg.ChannelMessageSend(b.chId, m)

	return err
}

func (b *Bot) EditMsg(i string, m string) error {
	b.l.Println("Editing message to: ", m)
	_, err := b.dg.ChannelMessageEdit(b.chId, i, m)

	return err
}

func (b *Bot) Close() {
	b.dg.Close()
}

func (b *Bot) registerCommands() error {
	c := commands.Get()
	ch := commands.GetHandlers()

	b.dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := ch[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	cmdIds := make([]*discordgo.ApplicationCommand, len(c))
	for i, v := range c {
		cmd, err := b.dg.ApplicationCommandCreate(b.dg.State.User.ID, "", &v)
		if err != nil {
			b.l.Panicf("Cannot create '%v' command: %v", v.Name, err)
			return err
		}
		cmdIds[i] = cmd
	}

	return nil
}
