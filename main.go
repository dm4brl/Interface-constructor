package main

import (
    "context"
    "github.com/jackc/pgx/v4"
    "github.com/spf13/viper"
)

// DatabaseService interface: Defines the methods that should be implemented by all database services.
type DatabaseService interface {
    Connect() (*pgx.Conn, error)
    Close(conn *pgx.Conn)
    InsertData(conn *pgx.Conn, data string) error
    GetData(conn *pgx.Conn) (string, error)
}

// PostgreSQLDatabaseService struct: Implements the DatabaseService interface for PostgreSQL.
type PostgreSQLDatabaseService struct {
    config DatabaseConfig
}

// DatabaseConfig struct: Holds the PostgreSQL database connection parameters.
type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    Database string
}

// NewPostgreSQLDatabaseService constructor: Creates a new instance of PostgreSQLDatabaseService.
func NewPostgreSQLDatabaseService(config DatabaseConfig) *PostgreSQLDatabaseService {
    return &PostgreSQLDatabaseService{config: config}
}

// Connect establishes a connection to the database.
func (s *PostgreSQLDatabaseService) Connect() (*pgx.Conn, error) {
    connStr := "postgres://" + s.config.User + ":" + s.config.Password + "@" + s.config.Host + ":" + s.config.Port + "/" + s.config.Database
    conn, err := pgx.Connect(context.Background(), connStr)
    if err != nil {
        return nil, err
    }
    return conn, nil
}

// Close closes the database connection.
func (s *PostgreSQLDatabaseService) Close(conn *pgx.Conn) {
    conn.Close(context.Background())
}

// InsertData inserts data into the database.
func (s *PostgreSQLDatabaseService) InsertData(conn *pgx.Conn, data string) error {
    _, err := conn.Exec(context.Background(), "INSERT INTO your_table_name (data_column) VALUES ($1)", data)
    return err
}

// GetData retrieves data from the database.
func (s *PostgreSQLDatabaseService) GetData(conn *pgx.Conn) (string, error) {
    var data string
    err := conn.QueryRow(context.Background(), "SELECT data_column FROM your_table_name").Scan(&data)
    if err != nil {
        return "", err
    }
    return data, nil
}

func main() {
    // Use Viper to read configuration data.
    viper.SetConfigFile("config.yaml")
    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }

    var dbConfig DatabaseConfig
    err := viper.UnmarshalKey("database", &dbConfig)
    if err != nil {
        panic(err)
    }

    // Create an instance of the database service.
    dbService := NewPostgreSQLDatabaseService(dbConfig)

    // Establish a connection to the database.
    conn, err := dbService.Connect()
    if err != nil {
        panic(err)
    }
    defer dbService.Close(conn)

    // Insert data into the database.
    err = dbService.InsertData(conn, "Hello, PostgreSQL!")
    if err != nil {
        panic(err)
    }

    // Retrieve data from the database.
    data, err := dbService.GetData(conn)
    if err != nil {
        panic(err)
    }

    // Print the data.
    println("Data from the database:", data)
}
