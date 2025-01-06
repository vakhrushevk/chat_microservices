package chatservice

import (
	"context"

	"github.com/vakhrushevk/chat-server-service/internal/converter"
	servicelevelmodel "github.com/vakhrushevk/chat-server-service/internal/service/serviceLevelModel"
)

// CreateChat - Создает чат и добавляет создателя в него
func (s *chatService) CreateChat(ctx context.Context, chat *servicelevelmodel.ChatInfo) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.repository.CreateChat(ctx, converter.ServiceToRepositoryChatInfo(chat))
		if errTx != nil {
			return errTx
		}

		smi := servicelevelmodel.ChatMemberInfo{ChatID: id, UserID: chat.CreatedBy}
		errTx = s.repository.AddChatMember(ctx, converter.ServiceToRepositoryChatMemberInfo(&smi))

		if errTx != nil {
			return errTx
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
