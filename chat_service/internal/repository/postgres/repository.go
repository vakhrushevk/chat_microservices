package postgres

import (
	"context"
	"github.com/vakhrushevk/chat-server-service/internal/logger"
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/local-platform/db"

	"github.com/Masterminds/squirrel"
	"github.com/vakhrushevk/chat-server-service/internal/repository/repositoryLevelModel"
)

const (
	table_chat_members   = "chat_service.chat_members"
	chatColumnName       = "name"
	table_chat           = "chat_service.chats"
	chatColumnCreated_by = "created_by"
	//columnCreatedAt      = "created_at"

)

type repo struct {
	db db.Client
}

// NewChatRepository - Создаем новый экземлпяр репозитория
func NewChatRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

// CreateChat - Создать чат
func (r *repo) CreateChat(ctx context.Context, chat *repositoryLevelModel.ChatInfo) (int64, error) {
	query, args, err := squirrel.
		Insert(table_chat).
		Columns(chatColumnName, chatColumnCreated_by).
		Values(chat.Name, chat.CreatedBy).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("returning id").
		ToSql()
	if err != nil {
		logger.Error("Failed to create chat", logger.ErrAttr(err))
		return 0, err
	}

	var id int64
	q := db.Query{Name: "CreateChat ", QueryRaw: query}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// AddChatMember - Добавить пользователя в чат
func (r *repo) AddChatMember(ctx context.Context, chat *repositoryLevelModel.ChatMemberInfo) error {
	query, args, err := squirrel.
		Insert(table_chat_members).
		Columns("chat_id",
			"user_id",
			"joined_at").
		PlaceholderFormat(squirrel.Dollar).
		Values(chat.ChatID, chat.UserID, "now()").
		ToSql()

	if err != nil {
		return err
	}
	q := db.Query{Name: "AddChatMember", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// RemoveChatMember - Удалить пользователя из чата
func (r *repo) RemoveChatMember(ctx context.Context, chat *repositoryLevelModel.ChatMemberInfo) error {
	query, args, err := squirrel.Delete(table_chat_members).
		Where(squirrel.Eq{"chat_id": chat.ChatID, "user_id": chat.UserID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{Name: "RemoveChatMember", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// ListChatsByIdUser - Получить список чатов пользователя по userID
func (r *repo) ListChatsByIdUser(ctx context.Context, userID int64) ([]*repositoryLevelModel.Chat, error) {
	query, args, err := squirrel.
		Select("c.id", "c.name", "c.created_by", "c.created_at", "c.updated_at").
		From("chat_service.chats c").
		Join("chat_service.chat_members cm ON c.id = cm.chat_id").
		Where(squirrel.Eq{"cm.user_id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{Name: "ListChatsByIdUser", QueryRaw: query}
	var chats []*repositoryLevelModel.Chat
	err = r.db.DB().ScanAllContext(ctx, &chats, q, args...)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

// ListMemberChat - Получить список участников чата
func (r *repo) ListMemberChat(ctx context.Context, chatID int64) ([]int64, error) {
	query, args, err := squirrel.
		Select("user_id").
		From(table_chat_members).
		Where(squirrel.Eq{"chat_id": chatID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{Name: "ListMemberChat", QueryRaw: query}
	var IDs []int64
	err = r.db.DB().ScanAllContext(ctx, &IDs, q, args...)
	if err != nil {
		return nil, err
	}

	return IDs, nil
}

// DeleteChat - Delete chat
func (r *repo) DeleteChat(ctx context.Context, idChat int64) error {
	query, args, err := squirrel.Delete(table_chat).
		Where(squirrel.Eq{"id": idChat}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}
	q := db.Query{Name: "DeleteChat", QueryRaw: query}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}
