package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter().StrictSlash(true)

	// server.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("app/assets"))))

	server.Router.HandleFunc("/", server.Home).Methods("GET")
	server.Router.HandleFunc("/listpet", server.Listpet).Methods("GET")
	server.Router.HandleFunc("/dashboard", server.Dashboard).Methods("GET")
	server.Router.HandleFunc("/dashboard/listpet", server.DashboardListPet).Methods("GET")
	server.Router.HandleFunc("/dashboard/add", server.AddPet).Methods("GET")
	server.Router.HandleFunc("/dashboard/add", server.CreatePet).Methods("POST")
	server.Router.HandleFunc("/admin/pet/create", server.CreatePet).Methods("POST")
	server.Router.HandleFunc("/pets/{name}", server.PetDetailsByName).Methods("GET")
	server.Router.HandleFunc("/submit-adoption", server.SubmitAdoption).Methods("POST")
	server.Router.HandleFunc("/dashboard/adoptions", server.AdoptionList).Methods("GET")
	server.Router.HandleFunc("/admin/pet/edit/{id}", server.EditPet).Methods("GET")
	server.Router.HandleFunc("/admin/pet/update/{id}", server.UpdatePet).Methods("POST")
	server.Router.HandleFunc("/dashboard/adoptions/{id}/approve", server.ApproveAdoption).Methods("POST")
	server.Router.HandleFunc("/dashboard/adoptions/{id}/reject", server.RejectAdoption).Methods("POST")
	server.Router.HandleFunc("/admin/pet/delete/{id}", server.DeletePet).Methods("POST")
	server.Router.HandleFunc("/login", server.Login).Methods("GET", "POST")
	server.Router.HandleFunc("/logout", server.Logout).Methods("GET", "POST")


	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
