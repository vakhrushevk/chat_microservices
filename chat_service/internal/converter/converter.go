package converter

import (
	modelRepo "github.com/vakhrushevk/chat-server-service/internal/repository/repositoryLevelModel"
	modelService "github.com/vakhrushevk/chat-server-service/internal/service/serviceLevelModel"
	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
)

// ServiceToRepositoryChatInfo - Конвертируем модель ChatInfo сервиса в модель ListChatsResponse desc
func ServiceToDescListChatsResponse(chats []*modelService.Chat) *chat_v1.ListChatsResponse {
	chatDesc := chat_v1.ListChatsResponse{}
	for _, chat := range chats {
		c := &chat_v1.ChatInfo{
			ChatId:    chat.ID,
			Name:      chat.ChatInfo.Name,
			MemberIds: chat.ChatInfo.ChatMembersIds,
		}
		chatDesc.Chats = append(chatDesc.Chats, c)
	}

	return &chatDesc
}

// ServiceToRepositoryChatInfo - Конвертируем модель ChatInfo cервиса в модель ChatInfo репозитория
func ServiceToRepositoryChatInfo(chat *modelService.ChatInfo) *modelRepo.ChatInfo {
	return &modelRepo.ChatInfo{
		Name:      chat.Name,
		CreatedBy: chat.CreatedBy,
	}
}

// ServiceToRepositoryChatMemberInfo - Конвертируем модель ChatMemberInfo cервиса в модель ChatMemberInfo репозитория
func ServiceToRepositoryChatMemberInfo(chat *modelService.ChatMemberInfo) *modelRepo.ChatMemberInfo {
	return &modelRepo.ChatMemberInfo{
		ChatID: chat.ChatID,
		UserID: chat.UserID,
	}
}

/*
func ChatToServiceLevel(chat *modelRepo.Chat) *modelService.Chat {
	var createdAt time.Time
	if chat.CreatedAt.Valid {
		createdAt = chat.CreatedAt.Time
	}

	return &modelService.Chat{
		ID:        chat.ID,
		ChatInfo:  RepositoryToServiceChatInfo(&chat.ChatInfo),
		CreatedAt: createdAt,
	}
}
*/

// RepositoryToServiceChatInfo - Конвертируем модель ChatInfo репозитория в модель ChatInfo сервиса
func RepositoryToServiceChatInfo(chatInfo *modelRepo.ChatInfo, ids []int64) modelService.ChatInfo {
	return modelService.ChatInfo{
		Name:           chatInfo.Name,
		CreatedBy:      chatInfo.CreatedBy,
		ChatMembersIds: ids,
	}
}

/*
func RepositoryToServiceChat(chat *modelRepo.Chat) *modelService.Chat {
	return &modelService.Chat{
		ID:        chat.ID,
		ChatInfo:  RepositoryToServiceChatInfo(&chat.ChatInfo),
		CreatedAt: chat.CreatedAt.Time,
		UpdatedAt: chat.UpdatedAt.Time,
	}
}
*/

// DescToServiceChatInfo - Конвертируем модель ChatInfo desc в модель ChatInfo сервиса
func DescToServiceChatInfo(chatInfo *chat_v1.CreateChatRequest) *modelService.ChatInfo {
	return &modelService.ChatInfo{
		Name:      chatInfo.Name,
		CreatedBy: chatInfo.CreatedBy,
	}
}

// DescToServiceChatMemberInfo - Конвертируем модель ChatMemberInfo desc в модель ChatMemberInfo сервиса
func DescToServiceChatMemberInfo(chatInfo *chat_v1.ChatMemberInfo) *modelService.ChatMemberInfo {
	return &modelService.ChatMemberInfo{
		ChatID: chatInfo.ChatId,
		UserID: chatInfo.UserId,
	}
}
