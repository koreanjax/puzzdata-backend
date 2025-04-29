package handlers

import (
    "log"
    "net/http"
    "fmt"
    "encoding/json"
    "strings"
    // "strconv"
    "puzzdata-backend/models"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
)

type Filter struct {
    l *log.Logger
    db *sqlx.DB
}

func NewFilter(l *log.Logger, db *sqlx.DB) *Filter {
    return &Filter{l, db}
}

const noFilter = "-1"

func CorrectAttribute(attribute string, attrPosition int, index int) string {
    if attrPosition == 0 {
        if index == 5 {
            return "6"
        } else {
            return attribute
        }
    } else {
        if index == 5 {
            return "-1"
        } else {
            return attribute
        }
    }
}

var frontTypeToBackType = map[string]string {
    "0": "5",
    "1": "4",
    "2": "7",
    "3": "8",
    "4": "1",
    "5": "6",
    "6": "2",
    "7": "3",
    "8": "0",
    "9": "12",
    "10": "14",
    "11": "15",
}

var frontAwknToBackAwkn = map[string]string {
    "1": "1",
    "2": "2",
    "3": "3",
    "4": "4",
    "5": "5",
    "6": "6",
    "7": "7",
    "8": "8",
    "9": "9",
    "10": "10",
    "11": "11",
    "12": "12",
    "13": "13",
    "14": "14",
    "15": "15",
    "16": "16",
    "17": "17",
    "18": "18",
    "19": "19",
    "20": "20",
    "21": "21",
    "22": "22",
    "23": "23",
    "24": "24",
    "25": "25",
    "26": "26",
    "27": "27",
    "28": "28",
    "29": "29",
    "30": "30",
    "31": "31",
    "32": "32",
    "33": "33",
    "34": "34",
    "35": "35",
    "36": "36",
    "37": "37",
    "38": "38",
    "39": "39",
    "40": "40",
    "41": "41",
    "42": "42",
    "43": "43",
    "44": "44",
    "45": "45",
    "46": "46",
    "47": "47",
    "48": "48",
    "49": "49",
    "50": "50",
    "51": "51",
    "52": "54",
    "53": "55",
    "54": "57",
    "55": "58",
    "56": "59",
    "57": "60",
    "58": "61",
    "59": "62",
    "60": "63",
    "61": "64",
    "62": "65",
    "63": "66",
    "64": "67",
    "65": "71",
    "66": "72",
    "67": "73",
    "68": "74",
    "69": "75",
    "70": "76",
    "71": "77",
    "72": "78",
    "73": "79",
    "74": "80",
    "75": "81",
    "76": "82",
    "77": "83",
    "78": "84",
    "79": "85",
    "80": "86",
    "81": "87",
    "82": "88",
    "83": "89",
    "84": "90",
    "85": "91",
    "86": "92",
    "87": "93",
    "88": "94",
    "89": "95",
    "90": "104",
    "91": "105",
    "92": "106",
    "93": "126",
    "94": "127",
    "95": "128",
    "96": "129",
    "97": "130",
    "98": "131",
}

func ConvertType(typeNumber string) string {
    return(frontTypeToBackType[typeNumber])
}

func JoinStringsNonEmpty(queryStrings []string, delimiter string) string {
    tempSlice := make([]string, 0)
    for _, queryString := range queryStrings {
        if len(queryString) > 0 {
            tempSlice = append(tempSlice, queryString)
        }
    }
    return strings.Join(tempSlice, delimiter)
}

func (f *Filter) FilterHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    cards := make([]models.Card, 0)
    
    attrsTableBase := "(SELECT monster_id FROM attrs "

    query := r.URL.Query()

    queryAnd := " AND "
    queryOr := " OR "

    f.l.Println(query)

    attrQuery := make([]string, 0)

    attributes := query.Get("attrs")
    attributeSplit := strings.Split(attributes, "|")

    for attrPosition, condition := range attributeSplit {
        if condition != noFilter {
            filterArray := make([]string, 0)
            attributeCondition := strings.Split(condition, ",")
            for index, attribute := range attributeCondition {
                if attribute != noFilter {
                    attrNumber := fmt.Sprintf("%d", attrPosition+1)
                    filter := "attr_" + attrNumber + "=" + CorrectAttribute(attribute,attrPosition,index)
                    filterArray = append(filterArray, filter)                    
                }
            }
            singleAttrQuery := "(" + strings.Join(filterArray, queryOr) + ")"
            attrQuery = append(attrQuery, singleAttrQuery)
        }
    }

    attrQueryFinal := (strings.Join(attrQuery, queryAnd))


    typesQuery := make([]string, 0)
    cardTypes := query.Get("types")

    if cardTypes != noFilter {
        cardTypeSplit := strings.Split(cardTypes, ",")
        for i := range 3{
            filterArray := make([]string, 0)
            for _, cardType := range cardTypeSplit {
                if cardType != noFilter {
                    typeNumber := fmt.Sprintf("%d", i+1)
                    filter := "type_" + typeNumber + "=" + ConvertType(cardType)
                    filterArray = append(filterArray, filter)
                }
            }
            typeQuery := "(" + strings.Join(filterArray, queryOr) + ")"
            typesQuery = append(typesQuery, typeQuery)
        }
    }

    typesQueryFinal := (strings.Join(typesQuery, queryOr))

    rarityFilter := make([]string, 0)

    rarity := query.Get("rarity")
    raritySplit := strings.Split(rarity, ",")

    for _, rarityNumber := range raritySplit {
        if rarityNumber != noFilter {
            filter := "rarity=" + rarityNumber
            rarityFilter = append(rarityFilter, filter)
        }
    }
    rarityQueryFinal :=strings.Join(rarityFilter, queryOr)
    if len(rarityQueryFinal) > 0 {
        rarityQueryFinal =  "(" + rarityQueryFinal + ")"
    }
    
    attrsTableQuery := JoinStringsNonEmpty([]string {attrQueryFinal, typesQueryFinal, rarityQueryFinal}, queryAnd)

    attrsTableQueryFinal := attrsTableBase
    if len(attrsTableQuery) > 0 {
        attrsTableQueryFinal = attrsTableQueryFinal + "WHERE "
    }
    attrsTableQueryFinal = attrsTableQueryFinal + attrsTableQuery + ")"

    awknsTableBase := "(SELECT DISTINCT monster_id, name FROM awkn "

    awknsQuery := make([]string, 0)
    awakenings := query.Get("awkns")

    if awakenings != noFilter {
        awakeningSplit := strings.Split(awakenings, "|")
        for _, condition := range awakeningSplit {
            awakeningCondition := strings.Split(condition, ",")
            filter := "awkn_" + frontAwknToBackAwkn[awakeningCondition[0]] + ">=" + awakeningCondition[1]
            awknsQuery = append(awknsQuery, filter)
        }

    }

    awknsTableQuery := strings.Join(awknsQuery, queryAnd)
    awknsTableQueryFinal := awknsTableBase
    if len(awknsTableQuery) > 0 {
        awknsTableQueryFinal = awknsTableQueryFinal + "WHERE "
    }
    awknsTableQueryFinal = awknsTableQueryFinal + awknsTableQuery + ")"

    dbQueryBase := "SELECT awkn.monster_id, awkn.name FROM %s awkn INNER JOIN %s attrs ON awkn.monster_id=attrs.monster_id WHERE awkn.name NOT LIKE \"%%*****%%\" AND awkn.name NOT LIKE \"%%????%%\";"
    dbQueryFinal := fmt.Sprintf(dbQueryBase, awknsTableQueryFinal, attrsTableQueryFinal)

    f.l.Println(dbQueryFinal)

    results, err := f.db.Query(dbQueryFinal)
    if err != nil {
        f.l.Println("ff.led to fetch information from DB", err)
    }

    for results.Next() {
        var card models.Card

        err = results.Scan(&card.MONSTER_ID, &card.NAME)
        
        if err != nil {
            http.Error(w, "Cards cannot be reached with name parameters", http.StatusInternalServerError)
            return
        }

        cards = append(cards, card)
    }
    defer results.Close()

    json.NewEncoder(w).Encode(cards)
}
