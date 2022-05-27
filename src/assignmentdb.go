package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db2, _ = sql.Open("sqlite3", "assignmentdb.db")

type Assignment_Form struct {
	IssueNumber int
	From	string
	To 	string
	DueDate	string
	Problem	string
	Comment	string
	Status string
}

var assignList []Assignment_Form

func getAssigment (db2 *sql.DB) []Assignment_Form{

	assignList = nil

	rows, err := db2.Query("Select * from tblAssignment")

	if err == nil{
		for rows.Next(){
			var number int
			var from string
			var to string
			var due string
			var problem string
			var comment string
			var status string

			err = rows.Scan(&number,&from,&to,&due,&problem,&comment,&status)

			if err == nil{
				assignList = append(assignList, Assignment_Form{IssueNumber: number, From: from, To: to, DueDate: due, Problem: problem, Comment: comment,Status: status})
			}else {
				fmt.Println(err)
			}
		}
	}else {
		fmt.Println(err)
	}

	rows.Close()
	return assignList
}

func addNewDb(db2 *sql.DB, from, to,dueDate, problem ,comment,status string) {


	stmt, err := db2.Prepare("INSERT INTO tblAssignment(Name, Toname, DueDate, Problem ,Comment, Status) values (?,?,?,?,?,?)") 
	
	checkError(err)

	res,err := stmt.Exec(from, to, dueDate, problem ,comment,status)
	checkError(err)

	id,err := res.LastInsertId()

	checkError(err)

	fmt.Println(" Assignment successfully added : ", id)

	fmt.Println(getAssigment(db2))
}

func updateAssignment(db2 *sql.DB, Status string, issueNumber int){

	stmt, err := db2.Prepare("update tblAssignment set Status=? where IssueNumber=?") 
	
	checkError(err)

	res,err := stmt.Exec(Status, issueNumber)
	checkError(err)

	_, err = res.RowsAffected()

	checkError(err)

	fmt.Println("Status Updated!")

	getAssigment(db2)
}
