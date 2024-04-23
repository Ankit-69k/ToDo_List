package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

var firestoreClient *firestore.Client

func GetAllTasks(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	userUID, ok := ctx.Value("userUID").(string)
	if !ok {
		http.Error(w, "User UID not found", http.StatusUnauthorized)
		return
	}

	docs, err := firestoreClient.Collection("todos").Doc(userUID).Collection("items").Documents(ctx).GetAll()
	if err != nil {
		http.Error(w, "Error fetching to-do items", http.StatusInternalServerError)
		return
	}

	todoItems := []models.ToDoItem{}
	for _, doc := range docs {
		item := models.ToDoItem{
			ID:     doc.Ref.ID,
			Task:   doc.Data()["task"].(string),
			Status: doc.Data()["status"].(string),
		}
		todoItems = append(todoItems, item)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoItems)
}

func CreateTask(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	userUID, ok := ctx.Value("userUID").(string)
	if !ok {
		http.Error(w, "User UID not found", http.StatusUnauthorized)
		return
	}

	var todoItem models.ToDoItem
	if err := json.NewDecoder(r.Body).Decode(&todoItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the new item
	docRef := firestoreClient.Collection("todos").Doc(userUID).Collection("items").NewDoc()
	todoItem.ID = docRef.ID

	_, err := docRef.Set(ctx, map[string]interface{}{
		"task":   todoItem.Task,
		"status": todoItem.Status,
	})
	if err != nil {
		http.Error(w, "Error adding to-do item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todoItem)
}

func TaskUpdate(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	userUID, ok := ctx.Value("userUID").(string)
	if !ok {
		http.Error(w, "User UID not found", http.StatusUnauthorized)
		return
	}

	// Extract the to-do item ID from the request
	vars := mux.Vars(r)
	itemID := vars["id"]
	if itemID == "" {
		http.Error(w, "To-do item ID is required", http.StatusBadRequest)
		return
	}

	var todoItem models.ToDoItem
	if err := json.NewDecoder(r.Body).Decode(&todoItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	docRef := firestoreClient.Collection("todos").Doc(userUID).Collection("items").Doc(itemID)
	

	// Update the task and status of the existing to-do item
	_, err := docRef.Update(ctx, []firestore.Update{
		{
			Path:  "task",
			Value: todoItem.Task,
		},
		{
			Path:  "status",
			Value: todoItem.Status,
		},
	})
	if err != nil {
		http.Error(w, "Error updating to-do item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todoItem)
}



func DeleteTask(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	userUID, ok := ctx.Value("userUID").(string)
	if !ok {
		http.Error(w, "User UID not found", http.StatusUnauthorized)
		return
	}

	// Extract the to-do item ID from the request
	vars := mux.Vars(r)
	itemID := vars["id"]
	if itemID == "" {
		http.Error(w, "To-do item ID is required", http.StatusBadRequest)
		return
	}

	docRef := firestoreClient.Collection("todos").Doc(userUID).Collection("items").Doc(itemID)

	// Delete the to-do item
	_, err := docRef.Delete(ctx)
	if err != nil {
		http.Error(w, "Error deleting to-do item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "To-do item deleted successfully")
}

	

