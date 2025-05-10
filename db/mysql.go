package db

import (
	"database/sql"
	"fmt"
	"termsql/types"

	_ "github.com/go-sql-driver/mysql"
)

type SQL interface {
	ListDatabases(keyword string) ([]string, error)
	Close()
	SwitchDatabase(database string) error
	GetTables() ([]string, error)
}

type Mysql struct {
	db *sql.DB
}

// Initialize connection
func Init(conn types.Connection) (SQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conn.Username, conn.Password, conn.Host, conn.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Ping to check if the connection is alive
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Mysql{db: db}, nil
}

func (m *Mysql) Close() {
	m.db.Close()
}

func (m *Mysql) SwitchDatabase(database string) error {
	_, err := m.db.Exec("USE " + database)
	return err
}

// Fetch list of databases with optional search keyword
func (m *Mysql) ListDatabases(keyword string) ([]string, error) {
	var query string
	if keyword != "" {
		query = fmt.Sprintf("SHOW DATABASES LIKE '%%%s%%'", keyword)
	} else {
		query = "SHOW DATABASES"
	}

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching databases: %w", err)
	}
	defer rows.Close()

	var databases []string
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			return nil, fmt.Errorf("error scanning database name: %w", err)
		}
		databases = append(databases, dbName)
	}

	return databases, nil
}

func (m *Mysql) GetTables() ([]string, error) {
	rows, err := m.db.Query("SHOW TABLES")
	if err != nil {
		return nil, fmt.Errorf("error retrieving tables: %w", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("error scanning table name: %w", err)
		}
		tables = append(tables, tableName)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row iteration: %w", err)
	}

	return tables, nil
}
