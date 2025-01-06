package chat

import (
	"context"

	"github.com/vakhrushevk/chat-server-service/internal/converter"
	"github.com/vakhrushevk/chat-server-service/internal/logger"
	"github.com/vakhrushevk/chat-server-service/internal/service"
	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
)

var _ chat_v1.ChatV1Server = (*Implementation)(nil)

type Implementation struct {
	chatService service.ChatService
	chat_v1.UnimplementedChatV1Server
}

func (c *Implementation) CreateChat(ctx context.Context, request *chat_v1.CreateChatRequest) (*chat_v1.CreateChatResponse, error) {
	id, err := c.chatService.CreateChat(ctx, converter.DescToServiceChatInfo(request))
	if err != nil {

		logger.Error("Error creating chat", logger.ErrAttr(err))
		return nil, err
	}

	return &chat_v1.CreateChatResponse{
		ChatId: id,
	}, nil
}

// DeleteChat - Удалить чат
func (c *Implementation) DeleteChat(ctx context.Context, request *chat_v1.DeleteChatRequest) (*chat_v1.DeleteChatResponse, error) {
	err := c.chatService.DeleteChat(ctx, request.ChatId)
	if err != nil {
		return &chat_v1.DeleteChatResponse{Success: false}, err
	}
	return &chat_v1.DeleteChatResponse{
		Success: true,
	}, nil
}

// AddChatMember - Добавить пользователя в чат
func (c *Implementation) AddChatMember(ctx context.Context, request *chat_v1.ChatMemberInfo) (*chat_v1.AddChatMemberResponse, error) {
	err := c.chatService.AddChatMember(ctx, converter.DescToServiceChatMemberInfo(request))
	if err != nil {
		return &chat_v1.AddChatMemberResponse{Success: false}, err
	}

	return &chat_v1.AddChatMemberResponse{Success: true}, nil
}

// RemoveChatMember - Удалить пользователя из чата
func (c *Implementation) RemoveChatMember(ctx context.Context, request *chat_v1.RemoveChatMemberRequest) (*chat_v1.RemoveChatMemberResponse, error) {
	err := c.chatService.RemoveChatMember(ctx, converter.DescToServiceChatMemberInfo(request.ChatMemberInfo))
	if err != nil {
		return &chat_v1.RemoveChatMemberResponse{Success: false}, err
	}

	return &chat_v1.RemoveChatMemberResponse{Success: true}, nil
}

// ListChats - Получить список чатов пользователя по userID
func (c *Implementation) ListChats(ctx context.Context, request *chat_v1.ListChatsRequest) (*chat_v1.ListChatsResponse, error) {
	chatInfo, err := c.chatService.ListChatsByIdUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}
	return converter.ServiceToDescListChatsResponse(chatInfo), nil
}

// NewChatImplementation - creates a new chat implementation
func NewChatImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{chatService: chatService}
}
