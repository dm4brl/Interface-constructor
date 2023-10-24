# Interface-method-and-constructor
Interface method and constructor architecture with interface configuration option to work with PostgreSQL. This approach allows you to create configurable database service instances


go mod init Interface-method-and-constructor

DatabaseService interface: Defines methods to be implemented by all database services. This interface provides an abstraction for database operations and includes the following methods:

Connect() (*pgx.Conn, error): Establishes a connection to the database and returns a connection object (pgx.Conn) or an error if no connection is established.

Close(conn *pgx.Conn): Closes the database connection, taking the connection object as a parameter.

InsertData(conn *pgx.Conn, data string) error: Inserts data into the database using the connection object and returns an error if something went wrong.

GetData(conn *pgx.Conn) (string, error): Retrieves data from the database using a connection object and returns the retrieved data and an error if there was a problem.

DatabaseConfig structure: Stores the parameters for connecting to a PostgreSQL database. This structure contains the following fields:

Host (database host).

Port (database port).

User (database user name).

Password (database user password).

Database (database name).

PostgreSQLDatabaseService structure and constructor:



