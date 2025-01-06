package service

import (
	"context"

	servicelevelmodel "github.com/vakhrushevk/chat-server-service/internal/service/serviceLevelModel"
)

// ChatService - интерфейс для работы с чатами
type ChatService interface {
	CreateChat(ctx context.Context, chat *servicelevelmodel.ChatInfo) (int64, error)
	DeleteChat(ctx context.Context, idChat int64) error
	AddChatMember(ctx context.Context, chat *servicelevelmodel.ChatMemberInfo) error
	RemoveChatMember(ctx context.Context, chat *servicelevelmodel.ChatMemberInfo) error
	ListChatsByIdUser(ctx context.Context, UserID int64) ([]*servicelevelmodel.Chat, error)
}
