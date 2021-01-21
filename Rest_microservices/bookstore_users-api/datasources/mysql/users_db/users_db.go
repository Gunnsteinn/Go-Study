package users_db

import (
	"database/sql"
	"fmt"
	// _ "github.com/go-sql-driver/mysql" is necessary to mitigate the follow error:
	// >>> panic: sql: unknown driver "mysql" (forgotten import?)
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

// Client is a variable. DB is a database handle representing a pool of zero or more underlying connections.
// username, password, host, schema are variables that are settings in Environment at the begins like this:
//  >> mysql_users_username=database configuration;
//  >> mysql_users_password=database configuration;
//  >> mysql_users_host=database configuration;
//  >> mysql_users_schema=database configuration;
var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured.")
}
