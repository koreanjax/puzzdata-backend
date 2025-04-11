package main

import(
    "os"
    "database/sql"
    "fmt"
    "log"
    handlers "puzzdata-backend/handlers"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)


const (
    USER = "root"
    PASSWORD = "*****"
    DBNAME = "test"
)

func main() {
    // Create a new Log for all packages to use
    l := log.New(os.Stdout, "puzzdata-api ", log.LstdFlags)

    // Create a full connection string for mysql
    connString := fmt.Sprintf("%s:%s@/%s", USER, PASSWORD, DBNAME)

    // Connect to the database using the connection string
    db, err := sql.Open("mysql", connString)
    log.Println(db)
    if err != nil {
        log.Println(err.Error())
    }
    defer db.Close()

    l.Println("Connection Successful!")

    // Create the handlers
    ch := handlers.NewCards(l, db)

    // Create a new serve mux and register the handlers
    r := mux.NewRouter()

    r.HandleFunc("/", ch.HomeHandler)

    searchRouter := r.PathPrefix("/search").Subrouter()
    searchRouter.HandleFunc("/id={id:[0-9]+}", ch.IdHandler).Methods(http.MethodGet)
    searchRouter.HandleFunc("/name={name}", ch.NameHandler).Methods(http.MethodGet)

    log.Fatal(http.ListenAndServe(":8080", r))
}