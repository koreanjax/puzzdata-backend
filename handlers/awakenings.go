package handlers

import (
    "log"
    "net/http"
    "encoding/json"
    "strings"
    "puzzdata-backend/models"
    "github.com/gorilla/mux"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
)

type Awakenings struct {
    l *log.Logger
    db *sqlx.DB
}

func NewAwakenings(l *log.Logger, db *sqlx.DB) *Awakenings {
    return &Awakenings{l, db}
}

func (a *Awakenings) BaseHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var cards = make([]models.Card, 0)

    vars := mux.Vars(r)
    awakenings := vars["awkns"]

    awakenings_split := strings.Split(awakenings, "|")


    q := "SELECT id, name FROM awkn_base WHERE name NOT LIKE \"%%*****%%\" AND"

    for _, condition := range awakenings_split {
        condition_split := strings.Split(condition, ",")
        q+=" `"
        q+=condition_split[0]
        q+="`>="
        q+=condition_split[1]
        q+=" AND"
    }

    q+=" id < 100000"

    results, err := a.db.Query(q)
    if err != nil {
        a.l.Println("failed to fetch information from DB", err)
    }

    for results.Next() {
        var card models.Card

        err = results.Scan(&card.ID, &card.NAME)
        
        if err != nil {
            http.Error(w, "Cards cannot be reached with name parameters", http.StatusInternalServerError)
            return
        }

        cards = append(cards, card)
    }
    defer results.Close()

    json.NewEncoder(w).Encode(cards)
}
