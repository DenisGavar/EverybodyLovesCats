package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/DenisGavar/EverybodyLovesCats/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) initAutharizationProcess(message *tgbotapi.Message) error {
	authLink, err := b.gererateAuthorizationLink(message.Chat.ID)
	if err != nil {
		log.Fatal(err)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(b.messages.Start, authLink))

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.tokenRepository.Get(chatID, repository.AccessTokens)
}

func (b *Bot) gererateAuthorizationLink(chatID int64) (string, error) {
	redirectURL := b.generateRedirectURL(chatID)
	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", err
	}

	err = b.tokenRepository.Save(chatID, requestToken, repository.RequestTokens)
	if err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken, redirectURL)
}

func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatID)
}
