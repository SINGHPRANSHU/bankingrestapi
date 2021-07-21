package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"restapi/helper"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)


		
var collection  = helper.ConnectDB()	
var transaction =helper.ConnectTransactionDb()

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
        // if we failed to get the absolute path respond with a 400 bad request
        // and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, r.URL.Path)

    // check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
        // if we got an error (that wasn't that the file doesn't exist) stating the
        // file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    // otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
func main()  {
    //init mux
	router:=mux.NewRouter()
	r := router.Methods(http.MethodGet,http.MethodPost).Subrouter()
	//routes
	r.HandleFunc("/api/customers",GetCustomers).Methods("GET")
	r.HandleFunc("/api/customer",GetCustomer).Methods("GET")
	r.HandleFunc("/api/customer/update", Transfer).Methods("POST")
	r.HandleFunc("/api/alltransaction",Getalltransaction).Methods("GET")
	r.Use(MiddlewareValidateUser)
	
	postR := router.Methods(http.MethodGet, http.MethodPost).Subrouter()
	postR.HandleFunc("/signup", Signup).Methods("POST")
	postR.HandleFunc("/login", Login).Methods("POST")
	
    spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	// router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("build/"))))
	// router.PathPrefix("/").Handler(http.StripPrefix("/",spa))
	router.PathPrefix("/").Handler(spa)
	

    headers:=handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization"})
    methods:=handlers.AllowedMethods([]string{"GET","POST"})
	origins:=handlers.AllowedOrigins([]string{"*"})
	


	err:=http.ListenAndServe(":"+os.Getenv("PORT"),handlers.CORS(headers,methods,origins)(router))
	log.Fatal(err)
	
}











