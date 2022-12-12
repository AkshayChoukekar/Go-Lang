package main

import (
	
"testing"
"database/sql")


func TestGetDataFromTable(t *testing.T){
	getQuery := "select * from Persons"
    want:='[{"ID":1,"LastName":"Thakur","FirstName":"Mahi","Age":27}]'
	got := GetDataFromTable(db, getQuery)
	if got!=want{
		println("Error while getting data from")
	}
}

func TestCreateTable(t *testing.T) {
	createQuery := "CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text,product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)"
	want := error(nil)
	got := CreateTable(db, createQuery)
	if want != got {
		println("Error while creating table")
	}
}

func TestDeleteDataFromTable(t *testing.T){
	deleteQuery := "DELETE FROM Persons WHERE FirstName = 'Akshay'"
	want:= error("")
	got:=DeleteDataFromTable(db,deleteQuery)
	if got!=want{
		println("Error while deleting data from table")
	}
}

func TestUpdateDataInTable(t *testing.T){
	updateQuery := "UPDATE Persons SET Age = '27' WHERE FirstName = 'Akshay'"
	want:= error(nil)
	got:=DeleteDataFromTable(db,deleteQuery)
	if got!=want{
		println("Error while deleting data from table")
	}
}
