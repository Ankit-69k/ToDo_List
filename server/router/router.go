package router

import(
	"server/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api/task", controllers.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", controllers.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", controllers.TaskComplete ).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", controllers.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", controllers.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks", controllers.DeleteAllTasks).Methods("DELETE", "OPTIONS")
	return router
}