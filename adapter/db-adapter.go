package adapter

import (
	"coba-db-mssql/config"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func ConnectDB() *sql.DB {
	dbConfig := config.GetDBConfig()
	connString := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%s;database=%s;",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Port,
		dbConfig.Dbname,
	)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")

	return conn
}
