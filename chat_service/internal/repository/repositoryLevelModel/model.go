package repositoryLevelModel

import "database/sql"

type Chat struct {
	ID        int64        `db:"id"`
	ChatInfo  ChatInfo     `db:""`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type ChatInfo struct {
	Name      string `db:"name"`
	CreatedBy int64  `db:"created_by"`
}

type ChatMembers struct {
	ID             int64          `db:"id"`
	ChatMemberInfo ChatMemberInfo `db:""`
	JoinedAt       sql.NullTime   `db:"joined_at"`
}

type ChatMemberInfo struct {
	ChatID int64 `db:"chat_id"`
	UserID int64 `db:"user_id"`
}
