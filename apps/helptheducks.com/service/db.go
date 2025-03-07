package main

import (
	"log"

	"github.com/dxe/helptheducks.com/service/config"
	"github.com/jmoiron/sqlx"
)

// TODO: make a separate script to init db b/c this is kinda sketch for prod.
func getDb() *sqlx.DB {
	//dsn, dbName := splitDbNameFromDsn(config.Dsn)
	//
	//db, err := sqlx.Open("mysql", dsn)
	//if err != nil {
	//	log.Fatalf("Error connecting to database: %v", err)
	//}
	//fmt.Println("Connected to database")
	//
	//// Create the database if it doesn't exist.
	//query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	//_, err = db.Exec(query)
	//if err != nil {
	//	log.Fatalf("Error creating database: %v", err)
	//}

	// Connect to the specific database.
	db, err := sqlx.Open("mysql", config.Dsn)
	if err != nil {
		log.Fatalf("Error connecting to specified database: %v", err)
	}

	//createTableQuery := `
	//	CREATE TABLE IF NOT EXISTS messages (
	//		id INT AUTO_INCREMENT PRIMARY KEY,
	//		submitted_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	//		ip_address VARCHAR(255),
	//		name VARCHAR(255) NOT NULL,
	//		email VARCHAR(255) NOT NULL,
	//		phone VARCHAR(255),
	//		outside_us BOOLEAN NOT NULL DEFAULT FALSE,
	//		zip VARCHAR(5),
	//		city VARCHAR(255),
	//		message TEXT,
	//		status VARCHAR(20) NOT NULL DEFAULT 'PENDING'
	//	)`
	//_, err = db.Exec(createTableQuery)
	//if err != nil {
	//	log.Fatalf("Error creating messages table: %v", err)
	//}

	return db
}

//func splitDbNameFromDsn(dsn string) (string, string) {
//	parts := strings.Split(dsn, "/")
//	return strings.Join(parts[:len(parts)-1], "/") + "/", parts[len(parts)-1]
//}
