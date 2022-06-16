package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db, _ = sql.Open("sqlite3", "./dataBase/accounts.db")

type UserAccounts struct {
	Number int
	Username string
	Password string
}

var accountList []UserAccounts

func getAccount (db *sql.DB) []UserAccounts{

	accountList = nil

	rows, err := db.Query("Select * from tblAccounts")

	if err == nil{
		for rows.Next(){
			var number int
			var usersName string
			var passwords string

			err = rows.Scan(&number,&usersName,&passwords)

			if err == nil{
				accountList = append(accountList, UserAccounts{Number: number, Username: usersName,Password: passwords})
			}else {
				fmt.Println(err)
			}
		}
	}else {
		fmt.Println(err)
	}

	rows.Close()
	return accountList
}

func addNew(db *sql.DB, userName, password string){

	stmt, err := db.Prepare("INSERT INTO tblAccounts(Username, Password) values (?,?)") 
	
	checkError(err)

	res,err := stmt.Exec(userName, password)
	checkError(err)

	id,err := res.LastInsertId()

	checkError(err)

	fmt.Println("son eklenen id : ", id)

	fmt.Println(getAccount (db))
}

func deleteAccount(db *sql.DB, Number int){

	stmt, err := db.Prepare("DELETE FROM tblAccounts where Number=?") 
	
	checkError(err)
	fmt.Println(Number)
	res,err := stmt.Exec(Number)
	fmt.Println(res)
	checkError(err)

	_, err = res.RowsAffected()

	checkError(err)

	fmt.Println("Account deleted")
	getAccount(db)
	//seq:= string(Number-1)
}

func updateAccount(db *sql.DB, userName, Password string, Number int){

	stmt, err := db.Prepare("update tblAccounts set userName=?, Password=? where Number=?") 
	
	checkError(err)

	res,err := stmt.Exec(userName,Password, Number)
	checkError(err)

	_, err = res.RowsAffected()

	checkError(err)

	fmt.Println("Password Updated!")

	getAccount(db)

}

func checkError (err error){
	if err != nil {
		panic("NOLUYİ")
	}
}
