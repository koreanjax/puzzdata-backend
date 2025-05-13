package handlers

import (
    "log"
    "encoding/json"
    "net/http"
    "strings"
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

const EVO_SKILL int = 232
const EVO_SKILL_LOOP int = 233
const queryAnd string = " AND "
const queryOr string = " OR "
const queryWhere string = " WHERE "
const queryUnion string = " UNION "


func (s *Skills) HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("API Homepage"))
}

func (s *Skills) SkillHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    vars := mux.Vars(r)
    id := vars["id"]

    q := `SELECT * FROM skill WHERE skill_id=?`

    var skillResult models.SkillResults

    err := s.db.Get(&skillResult, q, id)
    if err != nil {
        s.l.Println("failed to fetch Skills from DB", err)
        return
    }

    skillType := skillResult.SKILL_TYPE
    var skillResults = make([]models.SkillResults, 0)
    if skillType == EVO_SKILL || skillType == EVO_SKILL_LOOP {
        queryEvo := "SELECT * FROM skill WHERE "
        parametersSplit := strings.Split(skillResult.PARAMETERS, "|")
        evoSkillId := make([]string, 0)
        evoSkillId = append(evoSkillId, "skill_id=" + id)
        for _, parameter := range parametersSplit {
            evoSkillId = append(evoSkillId, "skill_id=" + parameter)
        }
        evoSkillQuery := strings.Join(evoSkillId, queryOr)
        queryEvoFinal := queryEvo + evoSkillQuery
        err = s.db.Select(&skillResults, queryEvoFinal)
    }

    if len(skillResults) > 0 {
        json.NewEncoder(w).Encode(skillResults)
    } else {
        json.NewEncoder(w).Encode(skillResult)
    }
}
