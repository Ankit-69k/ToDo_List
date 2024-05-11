package controllers

import (
	"fmt"
	"server/db"
	"server/models"

	"cloud.google.com/go/firestore"
)

type DB interface {
	GetAllTasks(id string) (db.TaskResponse, error)
	CreateTask(task models.ToDoItem)
	DeleteTask(id string)
}

func GetAllTasks(id string, db *db.Store) (*[]db.TaskResponse, error) {

	tasks,err := db.GetAllTasks(id)

	if err != nil {
		return nil, err
	}

	return &tasks, nil
}

func CreateTask(task models.ToDoItem,db *db.Store) (*firestore.DocumentRef,error) {
	docRef, err := db.CreateTask(task);
	
	if err != nil {
		fmt.Println("Error while creating task")
		return nil,err
	}

	fmt.Printf("Task created successfully with ID: %s\n", docRef.ID)

	return docRef,err
}

func DeleteTask(id string,db *db.Store) (*firestore.WriteResult, error){
	doc, err := db.DeleteTask(id);

	if err != nil {
		fmt.Println("Error while deleting task")
		return nil,err
	}

	fmt.Printf("Task deleted successfully at: %s\n", doc.UpdateTime)

	return doc,err
}

// func DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	// Initialize Firestore client
// 	firestoreClient, err := db.InitFirestore()
// 	if err != nil {
// 		http.Error(w, "Firestore client not initialized", http.StatusInternalServerError)
// 		log.Println("Error: Firestore client not initialized")
// 		return
// 	}

// 	// Get task ID from URL
// 	taskID := r.URL.Path[len("/api/deleteTask/"):]
// 	if taskID == "" {
// 		http.Error(w, "Task ID not provided", http.StatusBadRequest)
// 		log.Println("Error: Task ID not provided")
// 		return
// 	}

// 	// Delete task from Firestore
// 	_, err = firestoreClient.Collection("Tasks").Doc(taskID).Delete(r.Context())
// 	if err != nil {
// 		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
// 		log.Println("Error deleting task from Firestore:", err)
// 		return
// 	}

// 	// Set response status
// 	w.WriteHeader(http.StatusNoContent)
// 	log.Println("Task with ID", taskID, "deleted successfully")
// }
