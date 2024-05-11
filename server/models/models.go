package models

type ToDoItem struct{
	Owner		string			`json:"owner"`
	Title		string			`json:"title"`
	Status		string			`json:"status"`
}