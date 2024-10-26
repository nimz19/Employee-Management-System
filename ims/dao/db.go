package dao

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // Use the MySQL driver
)

func Connect() (*sql.DB, error) {
    // Replace the connection string with your MySQL details
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ims_db")
    if err != nil {
        return nil, err
    }

    // Test the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
