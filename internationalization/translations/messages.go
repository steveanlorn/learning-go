package translations

import "github.com/nicksnyder/go-i18n/v2/i18n"

type MessageID string

const (
	HelloWorldMessageID MessageID = "HelloWorld"
	GoodByeMessageID    MessageID = "GoodBye"
)

var _staticMessages = map[MessageID]*i18n.LocalizeConfig{
	HelloWorldMessageID: &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          string(HelloWorldMessageID),
			Other:       "はじめまして",
			Description: "context:message to greet new people",
		},
	},
	GoodByeMessageID: &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          string(GoodByeMessageID),
			Other:       "さようなら",
			Description: "context:message to say good bye",
		},
	},
}
