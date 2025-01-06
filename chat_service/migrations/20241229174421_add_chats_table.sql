-- +goose Up
-- Создание схемы chat_service
CREATE SCHEMA IF NOT EXISTS chat_service;

-- Таблица для хранения чатов
CREATE TABLE chat_service.chats (
                                    id SERIAL PRIMARY KEY, -- Уникальный идентификатор чата
                                    name VARCHAR(255) NOT NULL, -- Название чата
                                    created_by INT NOT NULL, -- Идентификатор пользователя, создавшего чат
                                    created_at TIMESTAMP DEFAULT NOW(),-- Время создания чата
                                    updated_at TIMESTAMP
);

-- Таблица для хранения участников чатов
CREATE TABLE chat_service.chat_members (
                                           id SERIAL PRIMARY KEY, -- Уникальный идентификатор записи
                                           chat_id INT NOT NULL REFERENCES chat_service.chats(id) ON DELETE CASCADE, -- Ссылка на чат
                                           user_id INT NOT NULL, -- Идентификатор участника чата
                                           joined_at TIMESTAMP DEFAULT NOW(), -- Время добавления в чат
                                           UNIQUE (chat_id, user_id) -- Ограничение уникальности для пары chat_id и user_id
);

-- +goose Down
drop table chat_service.chat_members;
drop table chat_service.chats;
