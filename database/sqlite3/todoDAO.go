package sqlite3

import (
	"database/sql"
)

//CreateDatabase : create Sqlite3 Db file about filePath
func CreateDatabase(filePath string) *sql.DB {
	sqliteDB, err := sql.Open("sqlite3", filePath)
	procError(err)
	return sqliteDB
}

//CreateTable : create table about tb_todo
func CreateTable(db *sql.DB) {
	tableSQL := `
	CREATE TABLE tb_todo (
        id integer primary key autoincrement, 
        title text not null,
        doneYN boolean not null,
        createDt varchar(14) not null,
        updateDt varchar(14) not null
    );
	`

	_, err := db.Exec(tableSQL)
	procError(err)
}

func procError(err error) {
	if err != nil {
		panic(err)
	}
}
