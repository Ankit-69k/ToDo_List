package models

type ToDoItem struct{
	ID			string			`json:"id"`		
	Task		string			`json:"title"`
	Status		string			`json:"status"`
}