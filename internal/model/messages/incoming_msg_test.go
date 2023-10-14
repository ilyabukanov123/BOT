package messages

import (
	"github.com/golang/mock/gomock"
	mock "github.com/ilyabukanov123/BOT/internal/mocks/messages"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mock.NewMockMessageSender(ctrl)

	model := New(sender)

	sender.EXPECT().SendMessage("hello", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 123,
	})
	assert.NoError(t, err)
}

func Test_OnUnknownCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mock.NewMockMessageSender(ctrl)
	sender.EXPECT().SendMessage("не знаю эту команду", int64(123))
	model := New(sender)

	err := model.IncomingMessage(Message{
		Text:   "some text",
		UserID: 123,
	})
	assert.NoError(t, err)
}
