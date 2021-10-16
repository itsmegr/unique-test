package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func HomeHandler(res http.ResponseWriter, req *http.Request){
	v :=  req.FormValue("v")

	if v == "" {
		http.Error(res, "V is empty", http.StatusBadRequest)
		return;
	}

	//then v should be chagable to int

	intV , err := strconv.Atoi(v);

	if err!=nil {
		http.Error(res, "V should be Integer", http.StatusBadRequest)
		return;
	}

	fmt.Fprintf(res, strconv.Itoa(intV*2))
}

func Handler() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/double",HomeHandler)
	return r;
}

func main() {
	http.ListenAndServe(":3000", Handler())
}