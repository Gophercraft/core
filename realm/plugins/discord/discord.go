package discord

import (
	"fmt"

	"github.com/Gophercraft/core/realm"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/log"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/websocket"
)

const blurple = "7289DAFF"

var AnnounceEnabled models.PropID = "gophercraft.realm.chat.AnnounceEnabled"

var (
	ErrBotTokenNotSet = fmt.Errorf("discord: Discord.BotToken not set in world config")
	ErrChannelNotSet  = fmt.Errorf("discord: Discord.ChannelID not set in world config")
)

type DiscordPlugin struct {
	Server        *realm.Server
	Session       *discordgo.Session
	BridgeHook    *discordgo.Webhook
	ChannelID     string
	NonFatalError error
}

func (p *DiscordPlugin) Activated() (bool, error) {
	return p.Session != nil, p.NonFatalError
}

func (p *DiscordPlugin) Terminate() error {
	return p.Session.Close()
}

func (p *DiscordPlugin) SendMessageFrom(s *realm.Session, body string) error {
	whook, err := p.Session.WebhookCreate(p.ChannelID, s.PlayerName(), "")
	if err != nil {
		return err
	}

	p.Session.WebhookExecute(whook.ID, whook.Token, false, &discordgo.WebhookParams{
		Username: s.PlayerName(),
		Content:  body,
	})

	return nil
}

func cmdToggleAnn(s *realm.Session) {
	if s.BoolProp(AnnounceEnabled) {
		s.Warnf("Announce disabled.")
		s.RemoveProp(AnnounceEnabled)
	} else {
		s.Warnf("Announce enabled.")
		s.SetProp(AnnounceEnabled, true)
	}
}

func (p *DiscordPlugin) GetChannelName() string {
	for _, g := range p.Session.State.Guilds {
		for _, c := range g.Channels {
			if c.ID == p.ChannelID {
				return c.Name
			}
		}
	}

	panic("unknown name")
}

func (p *DiscordPlugin) Init(server *realm.Server, info *realm.PluginInfo) error {
	p.Server = server

	p.Server.Cmd(realm.Player, "discord toggle", "Enable or disable Discord on your account", cmdToggleAnn)

	info.ID = "gophercraft.core.discord"
	info.Name = "Gophercraft Official Discord Plugin"
	info.Version = "0.1"
	info.Authors = []string{"The Gophercraft Developers"}

	token := server.StringVar("Discord.BotToken")

	// User may not have Discord configured.
	if token == "" {
		p.NonFatalError = ErrBotTokenNotSet
		log.Warn(p.NonFatalError)
		return nil
	}

	p.ChannelID = server.StringVar("Discord.ChannelID")

	if p.ChannelID == "" {
		return ErrChannelNotSet
	}

	// p.Server.On(realm.ChatEvent, nil, func(session *realm.Session, msg *chat.Message) {

	// })

	// User network may be temporarily unreachable.
	var err error
	p.Session, err = discordgo.New("Bot " + token)
	if err != nil {
		p.NonFatalError = err
		return nil
	}

	p.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == p.Session.State.User.ID {
			return
		}

		if m.ChannelID != p.ChannelID {
			return
		}

		msgData := fmt.Sprintf("[|c%s%s|r] %s: %s", blurple, p.GetChannelName(), m.Message.Author.Username, m.Message.Content)

		p.Server.AllSessions().Iter(func(s *realm.Session) {
			if s.BoolProp(AnnounceEnabled) {
				s.SystemChat(msgData)
			}
		})
	})

	p.Server.Cmd(realm.Player, "announce", "Send a message to the server's discord channel", func(s *realm.Session, msg string) {
		p.Session.WebhookExecute(p.BridgeHook.ID, p.BridgeHook.Token, false, &discordgo.WebhookParams{
			Content:  msg,
			Username: s.PlayerName(),
		})
	})

	if err = p.Session.Open(); err != nil {
		switch err.(type) {
		case *websocket.CloseError:
			return fmt.Errorf("discord: %s", err)
		}
		p.NonFatalError = err
		return nil
	}

	hooks, err := p.Session.ChannelWebhooks(p.ChannelID)
	if err != nil {
		return err
	}

	msghookExists := false
	for _, hk := range hooks {
		if hk.Name == "BridgeHook" {
			p.BridgeHook = hk
			msghookExists = true
			break
		}
	}

	if !msghookExists {
		p.BridgeHook, err = p.Session.WebhookCreate(p.ChannelID, "BridgeHook", "")
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	realm.RegisterPlugin("discord", &DiscordPlugin{})
}
