package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/controllers"
	"server/db"
	"server/middleware"
	"server/models"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

type Server struct {
	port int
	db   *db.Store
}

func NewServer(db *db.Store) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("port:",NewServer.port);
	return server
}

func (s *Server) RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	
	router.Use(middleware.CorsMiddleware)
	
	router.HandleFunc("/api/tasks/{id}", s.GetAllTasks).Methods("GET")
	router.HandleFunc("/api/createTask", s.CreateTask).Methods("POST")
	router.HandleFunc("/api/deleteTask/{id}", s.DeleteTask).Methods("DELETE")
	
	return router
}

func (s *Server) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	
	
	id := mux.Vars(r)["id"]
	fmt.Println("id:",id)
	tasks, err := controllers.GetAllTasks(id, s.db)

	if err != nil {
		fmt.Println("error in finding the tasks")
	}

	fmt.Println(tasks)
	
	json.NewEncoder(w).Encode(tasks)
}

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {

	var task models.ToDoItem;
	
	// if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
	// 	http.Error(w, "Invalid JSON data", http.StatusBadRequest)
	// 	return
	// }

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&task); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		log.Println("Error decoding JSON data:", err)
		return
	}


	docRef, err := controllers.CreateTask(task, s.db);
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		log.Println("Error creating new task in Firestore:", err)
		return
	}
	
	

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Println("Error encoding response:", err)
	}
	fmt.Println("Task created successfully with ID:", docRef.ID)
}

func (s *Server) DeleteTask(w http.ResponseWriter, r *http.Request){
	taskID := r.URL.Path[len("/api/deleteTask/"):]
	doc, err := controllers.DeleteTask(taskID, s.db);

	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		log.Println("Error deleting task from Firestore:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Task deleted successfully at:", doc.UpdateTime)
}
