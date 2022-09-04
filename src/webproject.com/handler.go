package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	_ "github.com/mattn/go-sqlite3"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  fmt.Sprintf("/%s",userName),
			MaxAge: 300,
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func index(response http.ResponseWriter, request *http.Request) {
	
	http.ServeFile(response, request, "../../docs/index.html")
	//userPage("deniz", response)
}

// login handler
func loginPage(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "../../docs/login.html")
	fmt.Println(getUserName(request),6)
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	vars := mux.Vars(request)

	arrayData := getAccount(db)
	
	for i:=0 ; i<len(arrayData) ; i++ {
		oneData := arrayData[i]
		fmt.Println(oneData.Username, name)
		fmt.Println(oneData.Password, pass)
		if oneData.Username == name && oneData.Password == pass && name != "" && pass != ""{
				// checking ..
				vars["Uname"] = name
				setSession(name, response)
				redirectTarget = fmt.Sprintf("/denizcamalan.github.io/%s", name)
				break
		}else if i==len(arrayData) && oneData.Username != name && oneData.Password != pass || (name == "" && pass == "") {
			redirectTarget = "/denizcamalan.github.io/loginpage"
			log.Println("Wrong Password")
		}else{
			redirectTarget = "/denizcamalan.github.io/loginpage"
			log.Println("Wrong Password")
		}
	}
	http.Redirect(response, request, redirectTarget, http.StatusFound)
}

func addNewAssign(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("from")
	toname := request.FormValue("to")
	due := request.FormValue("date")
	issue := request.FormValue("issue")
	comment := request.FormValue("comment")
	status := request.FormValue("status")
	
	addNewDb(db2, name,toname,due,issue,comment,status)
	http.Redirect(response, request, "/"+name, http.StatusFound)
	//fmt.Println(getUserName(request),3453)
}

// add assignment and send other user

func dbpage(response http.ResponseWriter, request *http.Request){
	http.ServeFile(response, request, "../../docs/madeform.html")

}

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		tmpl := template.Must(template.ParseFiles("../../docs/innerPage.html"))
        	variable := AccountName{
				PageName: userName,
				WserName: userName,
				ActionLog: userName,
				ActionCreate: userName,
				ActionGetAss: userName,
				ActionUpdate: userName,
				ActionDelete: userName,
			}
        	tmpl.Execute(response, variable)
			
		http.ServeFile(response, request, "../../docs/innerPage.html")
		fmt.Println(getUserName(request),7)
	}
}
// obseved assignments

func assignmentPage(response http.ResponseWriter, request *http.Request){
	data := getAssigment(db2)
	vars := mux.Vars(request)
	tmpl := template.Must(template.ParseFiles("../../docs/layout.html"))
	for i:=0 ; i<len(data); i++{
		if vars["Uname"] ==  data[i].To{
			b := Assignment_Form{IssueNumber: data[i].IssueNumber ,From: data[i].From,To: data[i].To ,DueDate: data[i].DueDate, Problem: data[i].Problem, Comment: data[i].Comment,Status: data[i].Status}
			no := fmt.Sprintf("#Issue Number %d",data[i].IssueNumber)
        	variable := TodoPageData{
				PageTitle: no,
				Todos: []Todo{
					{Title: "From: ", Value:  b.From, Done: true},
					{Title: "Due Date: ", Value: b.DueDate, Done: true},
					{Title: "Issue: ", Value: b.Problem, Done: true},
					{Title: "Comment: ", Value: b.Comment, Done: true},
					{Title: "Status: ", Value: b.Status, Done: true},
				},
       		}
        	tmpl.Execute(response, variable)
		}
	}

	// add proper code for looping html.
}

// update status

func updateAssHandler(response http.ResponseWriter, request *http.Request){
	data := getAssigment(db2)
	number := request.FormValue("issueNumber")
	intNum, _ := strconv.Atoi(number)
	status := request.FormValue("status")

	for k := range data {
		if data[k].IssueNumber == intNum{
			fmt.Println(data[k].IssueNumber)
			fmt.Println(data[k].To)
			fmt.Println(data[k].From)
			updateAssignment(db2,status,intNum)
			redirectTarget = fmt.Sprintf("/denizcamalan.github.io/%s", data[k].To)
			fmt.Println(redirectTarget)
			break
		}else{
			redirectTarget = fmt.Sprintf("/denizcamalan.github.io/%s", data[k].To)
		}
	}
	http.Redirect(response, request, redirectTarget, http.StatusFound)
	fmt.Println(getUserName(request),90)
}
func updateAssPageHandler(response http.ResponseWriter, request *http.Request){
	http.ServeFile(response, request, "../../docs/updass.html")
	fmt.Println(getUserName(request),70)
}

// add new user
func signinPage(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "../../docs/signin.html")
	fmt.Println(getUserName(request),10)
}

func signinHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("newName")
	pass := request.FormValue("newPassword")

		if name != "" && pass != ""{
			redirectTarget = "/denizcamalan.github.io/loginpage"
			addNew(db,name,pass)
			fmt.Println(getUserName(request),1)
			fmt.Println("added new user")
		}else {
			redirectTarget = "/denizcamalan.github.io/signinpage"
			fmt.Println(getUserName(request),2)
		}	
	http.Redirect(response, request, redirectTarget, http.StatusFound)
}

// password Uupdated handler

func updatepassHandler(response http.ResponseWriter, request *http.Request){
	pass := request.FormValue("password")
	vars := mux.Vars(request)
	name := vars["Uname"]
	arrayData := getAccount(db)
	for i:=0 ; i<len(arrayData) ; i++ {
		oneData := arrayData[i]
		if (oneData.Username) == name{
			redirectTarget = "/denizcamalan.github.io/"
			updateAccount(db,oneData.Username,pass,oneData.Number)
			fmt.Println(getUserName(request),20)
		}
	}
	clearSession(response)
	http.Redirect(response, request, redirectTarget, http.StatusFound)

	// if make json all of the name and password than dont need use "for" loop for update.
}

// logout handler

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	vars["Uname"] = ""
	clearSession(response)
	http.Redirect(response, request, "/", http.StatusFound)
	fmt.Println(getUserName(request),3)
}

// delete account

func deleteHandler(response http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	name := vars["Uname"]
	arrayData := getAccount(db)

	for i:=0 ; i<len(arrayData) ; i++ {
		oneData := arrayData[i]
		if (oneData.Username) == name{
			redirectTarget = "/denizcamalan.github.io/"
			deleteAccount(db,oneData.Number)
			fmt.Println(getUserName(request),20)
		}
	}
	clearSession(response)
	http.Redirect(response, request, redirectTarget, http.StatusFound)
}