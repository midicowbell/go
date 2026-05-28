package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	url := "postgres://postgres:12345@localhost:5432/warehouse_iot"
	return pgx.Connect(ctx, url)
}
func CreateTables(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS products (
		product_id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		weight_grams INTEGER NOT NULL CHECK (weight_grams > 0),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS shelves (
		shelf_id SERIAL PRIMARY KEY,
		product_id INT REFERENCES products(product_id) ON DELETE SET NULL,
		current_weight INTEGER DEFAULT 0 CHECK (current_weight >= 0),
		min_count_limit INTEGER DEFAULT 0 CHECK (min_count_limit >= 0),
		status VARCHAR(50) DEFAULT 'OK',
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS weight_logs (
		id SERIAL PRIMARY KEY,
		shelf_id INTEGER REFERENCES shelves(shelf_id) ON DELETE CASCADE, -- Исправили имя таблицы
		raw_weight INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := conn.Exec(ctx, sqlQuery); err != nil {
		return err
	}
	return nil
}

func InsertRow(conn pgx.Conn, ctx context.Context) error {
	sqlQuery := `
		INSERT INTO products(name, weight_grams)
		VALUES('Зерновой кофе', 3500);
	`
	if _, err := conn.Exec(ctx, sqlQuery); err != nil {
		return err
	}
	return nil
}
