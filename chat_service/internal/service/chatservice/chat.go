package chatservice

import (
	"context"

	"github.com/vakhrushevk/chat-server-service/internal/converter"
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/chat-server-service/internal/service"
	servicelevelmodel "github.com/vakhrushevk/chat-server-service/internal/service/serviceLevelModel"
	"github.com/vakhrushevk/local-platform/db"
)

type chatService struct {
	repository repository.ChatRepository
	txManager  db.TxManager
}

// AddChatMember - Добавить пользователя в чат
func (s *chatService) AddChatMember(ctx context.Context, chat *servicelevelmodel.ChatMemberInfo) error {
	return s.repository.AddChatMember(ctx, converter.ServiceToRepositoryChatMemberInfo(chat))
}

// RemoveChatMember - Удалить пользователя из чата
func (s *chatService) RemoveChatMember(ctx context.Context, chat *servicelevelmodel.ChatMemberInfo) error {
	return s.repository.RemoveChatMember(ctx, converter.ServiceToRepositoryChatMemberInfo(chat))
}

// ListChatsByIdUser - Получить список чатов пользователя по userID
// TODO: Переписать, убрать геморой с указателями
func (s *chatService) ListChatsByIdUser(ctx context.Context, UserID int64) ([]*servicelevelmodel.Chat, error) {

	chats, err := s.repository.ListChatsByIdUser(ctx, UserID)
	if err != nil {
		return nil, err
	}
	var result []*servicelevelmodel.Chat

	for _, chat := range chats {
		IDsUser, err := s.repository.ListMemberChat(ctx, chat.ID)
		if err != nil {
			return nil, err
		}

		chServ := servicelevelmodel.Chat{
			ID:       chat.ID,
			ChatInfo: converter.RepositoryToServiceChatInfo(&chat.ChatInfo, IDsUser),
		}

		result = append(result, &chServ)
	}
	return result, nil
}

// DeleteChat - Удалить чат
func (s *chatService) DeleteChat(ctx context.Context, idChat int64) error {
	return s.repository.DeleteChat(ctx, idChat)
}

// New - creates a new chat level service
func New(chatRepository repository.ChatRepository, txManager db.TxManager) service.ChatService {
	return &chatService{repository: chatRepository, txManager: txManager}
}
