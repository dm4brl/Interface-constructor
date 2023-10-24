The "PostgreSQLDatabaseService" structure and constructor, the use of Viper for configuration data, and the purpose of the main function. 


# Interface-constructor
Interface method and constructor architecture with interface configuration option to work with PostgreSQL. This approach allows you to create configurable database service instances


The "PostgreSQLDatabaseService" structure and constructor implementation:

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

PostgreSQLDatabaseService is an implementation of the DatabaseService interface for working with PostgreSQL. It contains a config field of type DatabaseConfig to store the connection configuration.

NewPostgreSQLDatabaseService(config DatabaseConfig) *PostgreSQLDatabaseService is a constructor that creates a new instance of PostgreSQLDatabaseService. The constructor takes the configuration as a parameter and returns a pointer to the new instance.

Using Viper to read configuration data:

The example uses the Viper library to read connection parameters from the configuration file (config.yaml). This provides secure storage of credentials and other configuration parameters.

The main function of the code creates an instance of the database service, establishes a connection to the database, performs data insertion and retrieval operations, and displays the results.



