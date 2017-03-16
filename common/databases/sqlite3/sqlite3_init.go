package sqlite3

import (
	"database/sql"
	"todoGolang/common/exceptions"
)

//CreateDatabase : create Sqlite3 Database file about filePath
func CreateDatabase(filePath string) *sql.DB {
	sqliteDB, err := sql.Open("sqlite3", filePath)

	exceptions.ProcError(err)
	return sqliteDB
}

//CreateTable : create table about tb_todo table
func CreateTable(db *sql.DB) {
	tableSQL := `
	CREATE TABLE IF NOT EXISTS tb_todo (
        id integer primary key autoincrement, 
        title text not null,
        doneYN boolean not null,
        createDt varchar(14) not null,
        updateDt varchar(14) not null
    );
	`
	_, err := db.Exec(tableSQL)
	exceptions.ProcError(err)
}
