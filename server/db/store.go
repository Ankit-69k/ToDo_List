package db

import (
	"context"
	"fmt"
	"log"
	"server/models"

	"cloud.google.com/go/firestore"
)

type Store struct {
	client *firestore.Client
}

type TaskResponse struct {
	ID   string        `json:"id"`   
	Task models.ToDoItem `json:"task"` 
}

func (s *Store) GetAllTasks(id string) ([]TaskResponse, error) {
	ctx := context.Background()
	tasks := make([]TaskResponse, 0);

	data, err := s.client.Collection("Tasks").Documents(ctx).GetAll()
	if err != nil {
		return tasks, err
	}
	for _, doc := range data {
		var task models.ToDoItem
		if doc.Data()["Owner"] != id {
			continue
		}
		
		doc.DataTo(&task)
		
		taskResponse := TaskResponse{
			ID: doc.Ref.ID,
			Task: task,
		}

		tasks = append(tasks, taskResponse)
	}
	
	
	return tasks, nil
}

func (s *Store) CreateTask(task models.ToDoItem) ( *firestore.DocumentRef,error) {
	ctx := context.Background()

	var newTask = models.ToDoItem{
		Owner: task.Owner,
		Title: task.Title,
		Status: task.Status,
	}

	log.Println(newTask)

	docRef, _, err := s.client.Collection("Tasks").Add(ctx, newTask)

	
	

	if err != nil {
		fmt.Println("Error while creating task")
	}

	fmt.Printf("Task created successfully with ID: %s\n", docRef.ID)
	return docRef ,err
}

func (s *Store) DeleteTask(id string) (*firestore.WriteResult, error){
	ctx := context.Background()

	doc, err := s.client.Collection("Tasks").Doc(id).Delete(ctx)

	if err != nil {
		fmt.Println("Error while deleting task")
	}

	fmt.Printf("Task created successfully at: %s\n", doc.UpdateTime)
	return doc, err
}