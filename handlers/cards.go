package handlers

import (
    "log"
    "fmt"
    "encoding/json"
    "net/http"
    "puzzdata-backend/models"
    "github.com/gorilla/mux"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
)

type Cards struct {
    l *log.Logger
    db *sqlx.DB
}

const redacted = `"%*****%"`

func NewCards(l *log.Logger, db *sqlx.DB) *Cards {
    return &Cards{l, db}
}

func (c *Cards) HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("API Homepage"))
}

func (c *Cards) IdHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var cards = make([]models.Card, 0)
    vars := mux.Vars(r)

    id := vars["id"]

    q := fmt.Sprintf(`SELECT id, name FROM card WHERE id < 100000 AND official=100 AND name NOT LIKE %s AND id=%s;`, redacted, id)

    results, err := c.db.Query(q)
    if err != nil  {
        c.l.Println("failed to fetch from DB", err)
        return
    }
    defer results.Close()

    for results.Next() {
        var card models.Card

        err = results.Scan(&card.ID, &card.NAME)
        
        if err != nil {
            http.Error(w, "Card cannot be reached with id: " + id, http.StatusInternalServerError)
            return
        }

        cards = append(cards, card)
    }

    json.NewEncoder(w).Encode(cards)
}

func (c *Cards) NameHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var cards = make([]models.Card, 0)

    vars := mux.Vars(r)
    name := vars["name"]

    queryFormat := `"%` + name + `%"`
    q := fmt.Sprintf(`SELECT id, name FROM card WHERE id < 100000 AND official=100 AND name NOT LIKE %s AND (name LIKE %s OR search_keywords LIKE %s);`, redacted, queryFormat, queryFormat)

    results, err := c.db.Query(q)

    if err != nil  {
        c.l.Println("failed to fetch from DB", err)
        return
    }

    for results.Next() {
        var card models.Card

        err = results.Scan(&card.ID, &card.NAME)
        
        if err != nil {
        	http.Error(w, "Card cannot be reached with name: " + name, http.StatusInternalServerError)
        	return
        }

        cards = append(cards, card)
    }
    defer results.Close()

    json.NewEncoder(w).Encode(cards)
}

func (c *Cards) CardHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)
    id := vars["id"]

    var queryResults models.QueryResults

    q := `SELECT * FROM card WHERE id < 100000 AND id=?`

    err := c.db.Get(&queryResults, q, id)
    if err != nil {
        c.l.Println("failed to fetch information from DB", err)
    }

    json.NewEncoder(w).Encode(queryResults)
}
