package main

import (
	    "log"
		"net/http"
		"restapi/helper"
		"github.com/gorilla/mux"
		"github.com/gorilla/handlers"
		"os"
		)


		
var collection  = helper.ConnectDB()	
var transaction =helper.ConnectTransactionDb()
func main()  {
	
	//  os.Setenv("PORT","4000")
	 
    //init mux
	router:=mux.NewRouter()
	r := router.Methods(http.MethodGet,http.MethodPost).Subrouter()
	//routes
	r.HandleFunc("/api/customers",GetCustomers).Methods("GET")
	r.HandleFunc("/api/customer",GetCustomer).Methods("GET")
	r.HandleFunc("/api/customer/update", UpdateBalance).Methods("POST")
	r.HandleFunc("/api/alltransaction",Getalltransaction).Methods("GET")
	r.Use(MiddlewareValidateUser)
	
	postR := router.Methods(http.MethodGet, http.MethodPost).Subrouter()
	postR.HandleFunc("/signup", Signup).Methods("POST")
	postR.HandleFunc("/login", Login).Methods("POST")
	

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("build/"))))

	

    headers:=handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization"})
    methods:=handlers.AllowedMethods([]string{"GET","POST"})
	origins:=handlers.AllowedOrigins([]string{"*"})
	


	err:=http.ListenAndServe(":"+os.Getenv("PORT"),handlers.CORS(headers,methods,origins)(router))
	log.Fatal(err)
	
}











