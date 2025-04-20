package main

import(
    "os"
    "fmt"
    "log"
    "puzzdata-backend/handlers"
    "net/http"
    "github.com/rs/cors"
    "github.com/gorilla/mux"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
)


const (
    USER = "root"
    PASSWORD = "*****"
    DBNAME = "test2"
)

func main() {
    // Create a new Log for all packages to use
    l := log.New(os.Stdout, "puzzdata-api ", log.LstdFlags)

    // Create a full connection string for mysql
    connString := fmt.Sprintf("%s:%s@/%s", USER, PASSWORD, DBNAME)

    // Connect to the database using the connection string
    db, err := sqlx.Open("mysql", connString)
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
    
    cardRouter := r.PathPrefix("/card/").Subrouter()
    cardRouter.HandleFunc("/id={id:[0-9]+}", ch.CardHandler).Methods(http.MethodGet)

    handler := cors.Default().Handler(r)

    log.Fatal(http.ListenAndServe(":8080", handler))
}
