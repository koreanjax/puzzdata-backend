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

func main() {
    // Create a new Log for all packages to use
    l := log.New(os.Stdout, "padge-api ", log.LstdFlags)

    // Create a full connection string for mysql
    connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQLUSER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQLHOST"), os.Getenv("MYSQLPORT"), os.Getenv("MYSQL_DATABASE"))

    // Connect to the database using the connection string
    db, err := sqlx.Open("mysql", connString)
    log.Println(db)
    if err != nil {
        log.Println(err.Error())
        l.Println("Connection Failed!")
    }
    defer db.Close()

    // Create the handlers
    ch := handlers.NewCards(l, db)
    sh := handlers.NewSkills(l, db)
    fh := handlers.NewFilter(l ,db)

    // Create a new serve mux and register the handlers
    r := mux.NewRouter()

    //searchRouter := r.PathPrefix("/search").Subrouter()
    r.HandleFunc("/search", ch.SearchHandler).Methods(http.MethodGet)
    //searchRouter.HandleFunc("/id={id:[0-9]+}", ch.IdHandler).Methods(http.MethodGet)
    //searchRouter.HandleFunc("/name={name}", ch.NameHandler).Methods(http.MethodGet)
    
    cardRouter := r.PathPrefix("/card/").Subrouter()
    cardRouter.HandleFunc("/id={id:[0-9]+}", ch.CardHandler).Methods(http.MethodGet)
     
    skillRouter := r.PathPrefix("/skill/").Subrouter()
    skillRouter.HandleFunc("/id={id:[0-9]+}", sh.SkillHandler).Methods(http.MethodGet)

    filterRouter := r.PathPrefix("/filter").Subrouter()
    filterRouter.HandleFunc("/settings", fh.FilterHandler).Methods(http.MethodGet)

    handler := cors.Default().Handler(r)

    err = http.ListenAndServe("[::]:8080", handler)
    if err != nil {
        log.Fatalf("ListenAndServe failed: %v", err)
    }
}
