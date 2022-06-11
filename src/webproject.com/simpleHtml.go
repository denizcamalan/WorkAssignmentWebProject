package main

type Todo struct {
    Title string
	Value string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

type AccountName struct{
    PageName string
    WserName string
    ActionLog  string
    ActionCreate string
    ActionGetAss string
    ActionUpdate string
    ActionDelete string
}