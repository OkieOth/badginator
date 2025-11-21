package dummy

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTables(db *sql.DB) {
	stmt := `
  CREATE TABLE IF NOT EXISTS Name (
    _id_ INTEGER PRIMARY KEY
    ,first TEXT
    ,middle TEXT
    ,last TEXT
  );`
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatalf("error while creating table: %v", err)
	}
	fmt.Println("Created table 'Name' :)")
}

func getTablesImpl(db *sql.DB) int {
	sqlStmt := "SELECT name FROM sqlite_master WHERE type='table' ORDER BY name;"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
	defer rows.Close()

	rowCount := 0
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			log.Fatal("error while scanning for colum 'name'")
		}
		fmt.Println("Found Table:", name)
		rowCount++
	}
	return rowCount
}

func GetTables(dbFile string) error {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	tableCount := getTablesImpl(db)
	if tableCount == 0 {
		createTables(db)
	} else {
		return nil
	}
	tableCount = getTablesImpl(db)
	if tableCount == 0 {
		log.Fatal("wasn't able to create and read tables :-/")
	}
	return nil
}
