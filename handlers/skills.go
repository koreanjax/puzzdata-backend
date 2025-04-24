package handlers

import (
    "log"
    "encoding/json"
    "net/http"
    "puzzdata-backend/models"
    "github.com/gorilla/mux"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
)

type Skills struct {
    l *log.Logger
    db *sqlx.DB
}

func NewSkills(l *log.Logger, db *sqlx.DB) *Skills {
    return &Skills{l, db}
}

func (s *Skills) HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("API Homepage"))
}

func (s *Skills) SkillHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    vars := mux.Vars(r)
    id := vars["id"]

    q := `SELECT * FROM skill WHERE id=?`

    var skillResults models.SkillResults

    err := s.db.Get(&skillResults, q, id)
    if err != nil {
        s.l.Println("failed to fetch Skills from DB", err)
        return
    }

    json.NewEncoder(w).Encode(skillResults)
}
