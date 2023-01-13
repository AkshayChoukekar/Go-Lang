package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Persons struct {
	ID        int
	LastName  string
	FirstName string
	Age       int
}

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "Asdf@12345",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "test_project_go_lang",
	}
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		println(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		println(pingErr)
	}
	println("Connected!")
	//getQuery := "select * from Persons"
	//insertQuery := "INSERT INTO Persons (LastName, FirstName, Age) VALUES ('Thakur', 'Anup', 25);"
	//createQuery := `CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text,
	// product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	//updateQuery := "UPDATE Persons SET Age = '27' WHERE FirstName = 'Akshay'"
	//deleteQuery := "DELETE FROM Persons WHERE FirstName = 'Akshay'"
	//println(CreateTable(db, insertQuery))
	//println(GetDataFromTable(db, getQuery))
	//println(UpdateDataInTable(db, updateQuery))
	//println(DeleteDataFromTable(db, deleteQuery))
}

func CreateTable(db *sql.DB, query string) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		println("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		println("Error %s when getting rows affected", err)
		return err
	}
	println("Rows affected when creating table:", rows)
	return nil
}

func GetDataFromTable(db *sql.DB, query string) string {
	var persons []Persons
	result, err := db.Query(query)
	if err != nil {
		println(err)
	}
	defer result.Close()
	for result.Next() {
		var per Persons
		if err := result.Scan(&per.ID, &per.LastName, &per.FirstName, &per.Age); err != nil {
			println(err)
		}
		persons = append(persons, per)
	}
	data, err := json.Marshal(persons)
	if err != nil {
		println(err)
	}
	return string(data)
}

func UpdateDataInTable(db *sql.DB, query string) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	println(num)
	return nil
}

func DeleteDataFromTable(db *sql.DB, query string) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		println("Error while deleting table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		println("Error while getting rows affected", err)
		return err
	}
	println("Rows affected when deleting table:", rows)
	return nil
}
