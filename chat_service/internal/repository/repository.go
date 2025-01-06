package repository

import (
	"context"

	"github.com/vakhrushevk/chat-server-service/internal/repository/repositoryLevelModel"
)

// ChatRepository -
type ChatRepository interface {
	// CreateChat - Создает чат
	CreateChat(ctx context.Context, chat *repositoryLevelModel.ChatInfo) (int64, error)
	// DeleteChat - Удалить чат
	DeleteChat(ctx context.Context, idChat int64) error
	// AddChatMember - Добавить пользователя в чат
	AddChatMember(ctx context.Context, chat *repositoryLevelModel.ChatMemberInfo) error
	// RemoveChatMember - Удалить пользователя из чата
	RemoveChatMember(ctx context.Context, chat *repositoryLevelModel.ChatMemberInfo) error
	// ListChatsByIdUser - Получить список чатов пользователя по userID
	ListChatsByIdUser(ctx context.Context, userID int64) ([]*repositoryLevelModel.Chat, error)
	// ListMemberChat - Получить список участников чата
	ListMemberChat(ctx context.Context, chatID int64) ([]int64, error)
}
