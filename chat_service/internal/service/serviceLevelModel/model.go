package servicelevelmodel

import (
	"time"
)

// Chat - информация о чате
type Chat struct {
	ID        int64
	ChatInfo  ChatInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ChatInfo - информация о структуре чате
type ChatInfo struct {
	Name           string
	CreatedBy      int64
	ChatMembersIds []int64
}

// ChatMember - информация о пользователе в чате
type ChatMember struct {
	ID             int64
	ChatMemberInfo ChatMemberInfo
	JoinedAt       time.Time
}

// ChatMemberInfo - информация о пользователе в чате
type ChatMemberInfo struct {
	ChatID int64 `db:"chat_id"`
	UserID int64 `db:"user_id"`
}
