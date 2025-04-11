package handlers

import (
    "log"
    "database/sql"
	"encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    models "puzzdata-backend/models"
    _ "github.com/go-sql-driver/mysql"
)

type Cards struct {
	l *log.Logger
	db *sql.DB
}

func NewCards(l *log.Logger, db *sql.DB) *Cards {
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

    q := `SELECT id, name, rarity, cost, official FROM card WHERE id=?;`

    results, err := c.db.Query(q, id)
    if err != nil  {
        c.l.Println("failed to fetch from DB", err)
        return
    }
    defer results.Close()

    for results.Next() {
        var card models.Card

        err = results.Scan(&card.ID, &card.NAME, &card.RARITY, &card.COST, &card.OFFICIAL)
        
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

    q := `SELECT id, name FROM card WHERE name=?;`

    results, err := c.db.Query(q, name);

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

    json.NewEncoder(w).Encode(cards)
}