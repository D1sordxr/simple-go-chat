package migrations

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func Migrate(conn *pgx.Conn) {
	migrateUsersTable(conn)
	migrateMessagesTable(conn)
}

func migrateUsersTable(conn *pgx.Conn) {
	ctx := context.Background()
	query := `CREATE TABLE IF NOT EXISTS users (
		    id SERIAL PRIMARY KEY,
		    created_at TIMESTAMPTZ DEFAULT NOW(),
		    updated_at TIMESTAMPTZ DEFAULT NOW(),
		    user_id UUID NOT NULL UNIQUE,
		    username VARCHAR NOT NULL
		)`

	_, err := conn.Exec(ctx, query)
	if err != nil {
		log.Printf("failed migrate users table: %v", err)
	}
}

func migrateMessagesTable(conn *pgx.Conn) {
	ctx := context.Background()
	query := `CREATE TABLE IF NOT EXISTS messages (
		    id SERIAL PRIMARY KEY,
			user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
			content TEXT NOT NULL,
			is_edited BOOLEAN DEFAULT FALSE,
    		created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		)`

	_, err := conn.Exec(ctx, query)
	if err != nil {
		log.Printf("failed migrate messages table: %v", err)
	}
}
