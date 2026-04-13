package foundation

import "database/sql"

type PostgresDB struct {
	conn *sql.Conn
}

func NewPostgresDB(conn *sql.Conn) *PostgresDB {
	return &PostgresDB{conn: conn}
}
