package handlers

import (
    "log"
    "net/http"
    "fmt"
    "encoding/json"
    "strings"
    "strconv"
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

// All the strings used in the filter function. Might need to split up some
// tasks to make it look clean.
const noFilter string = "-1"
const whereClause string = "WHERE "
const attrTableBase string = "(SELECT monster_id FROM attr "
const typeBase string = "types LIKE "
const awknsTableBase string = "(SELECT DISTINCT monster_id, name FROM awkn "
const dbQueryBase string = "SELECT awkn.monster_id, awkn.name FROM %s awkn INNER JOIN %s attr ON awkn.monster_id=attr.monster_id WHERE awkn.name NOT LIKE \"%%*****%%\" AND awkn.name NOT LIKE \"%%????%%\""
const skillTableBase string = "SELECT monster_id, name, active_skill FROM card c INNER JOIN (%s) s ON c.active_skill = s.skill_id"
const skillTypeBase string = "SELECT DISTINCT skill_id FROM skill s1 INNER JOIN (%s) s2 ON s1.skill_id = s2.root_id"
const skillTypeUnionBase string = "SELECT root_id FROM skill_type_%s"
const skillTypeJoin string = "INNER JOIN (%s) s%s ON s%s.root_id = s%s.root_id"
const multiSkillQualifyBase string = "SELECT * FROM skill WHERE %s"
const paramInBase string = "%s IN (%s)"
const paramColunmn string = "param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8"
const cardFromSkillBase string = "SELECT c.monster_id, c.name FROM card c INNER JOIN (%s) s ON c.active_skill=s.skill_id WHERE c.name NOT LIKE \"%%*****%%\" AND c.name NOT LIKE \"%%????%%\""
const multiSkillBase string = "SELECT c1.monster_id, c1.name FROM card c1 INNER JOIN (SELECT skill_id FROM skill WHERE (%s)) s1 ON c1.active_skill=s1.skill_id WHERE c1.name NOT LIKE \"%%*****%%\" AND c1.name NOT LIKE \"%%????%%\""
const singularMultiMerge string = "SELECT monster_id, name FROM ((%s) UNION (%s)) as c2 ORDER BY monster_id"
const combinedBase string = "SELECT cs.monster_id, cs.name FROM (%s) cs INNER JOIN (%s) ataw ON cs.monster_id=ataw.monster_id"
const multiSkillType string = `multi_skills RLIKE '\\|[0-9]+,%s\\|'`

// For reasons that only God knows, No-Attr has a value of 6 for main attribute.
// Sub and Third attributes have a value of -1. Maybe it's related to how they
// want to sort?
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

// Dictionaries for conversion but I'm sure there is a better way to handle
// these types of operations?
var frontTypeToBackType = map[string]string {
    "0": "5", "1": "4", "2": "7", "3": "8", "4": "1", "5": "6",
    "6": "2", "7": "3", "8": "0", "9": "12", "10": "14", "11": "15",
}

var frontAwknToBackAwkn = map[string]string {
    "1": "1", "2": "2","3": "3", "4": "4", "5": "5",
    "6": "6", "7": "7", "8": "8", "9": "9", "10": "10",
    "11": "11", "12": "12", "13": "13", "14": "14", "15": "15",
    "16": "16", "17": "17", "18": "18", "19": "19", "20": "20",
    "21": "21", "22": "22", "23": "23", "24": "24", "25": "25",
    "26": "26", "27": "27", "28": "28", "29": "29", "30": "30",
    "31": "31", "32": "32", "33": "33", "34": "34", "35": "35",
    "36": "36", "37": "37", "38": "38", "39": "39", "40": "40",
    "41": "41", "42": "42", "43": "43", "44": "44", "45": "45",
    "46": "46", "47": "47", "48": "48", "49": "49", "50": "50",
    "51": "51", "52": "54", "53": "55", "54": "57", "55": "58",
    "56": "59", "57": "60", "58": "61", "59": "62", "60": "63",
    "61": "64", "62": "65", "63": "66", "64": "67", "65": "71",
    "66": "72", "67": "73", "68": "74", "69": "75", "70": "76",
    "71": "77", "72": "78", "73": "79", "74": "80", "75": "81",
    "76": "82", "77": "83", "78": "84", "79": "85", "80": "86",
    "81": "87", "82": "88", "83": "89", "84": "90", "85": "91",
    "86": "92", "87": "93", "88": "94", "89": "95", "90": "105",
    "91": "106", "92": "126", "93": "127","94": "128", "95": "129",
    "96": "130", "97": "131",
}

//  Manual conversion of frontend int to backend int for types
func ConvertType(typeNumber string) string {
    return(frontTypeToBackType[typeNumber])
}

//  Manual conversion of frontend int to backend int for awakenings
func ConvertAwakening(awknNumber string) string {
    return(frontAwknToBackAwkn[awknNumber])
}

// strings.Join can return an empty string, causing unnecessary Joins on
// empty query parts. This checks if each string is empty and chooses not
// to join those
//
func JoinStringsNonEmpty(queryStrings []string, delimiter string) string {
    tempSlice := make([]string, 0)
    for _, queryString := range queryStrings {
        if len(queryString) > 0 {
            tempSlice = append(tempSlice, queryString)
        }
    }
    return strings.Join(tempSlice, delimiter)
}


func AddParanthesis(text string) string {
    return "(" + text + ")"
}

func findString(arr []string, target string) (bool, string, string) {
    for _, value := range arr {
        if strings.Contains(value, target) {
            multiSkillSplit := strings.Split(strings.Trim(value, "|"), ",")
            return false, multiSkillSplit[0], multiSkillSplit[1] // skill_id and skill_type
        }
    }
    return true, "", "" // Return empty strings if not found
}

func removeLast(slice []string) []string {
    i := len(slice)
    if (i > 0) {
        return slice[:i-1]
    }
    return slice
}

/*
Main filter handler for now. It uses 3 different tables to search thoroughly.

Attrs table has the information for attributes, types, and rarity.

Awkns table has the information for awakenings, including all their variations.
It has a column to check if it includes SA so that non-SA variations can be searched
later on.

Skills table has the information for skills, including all their parameters that one
needs to check.
 */
func (f *Filter) FilterHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    cards := make([]models.Card, 0)

    query := r.URL.Query()

    // Check the query from frontend
    f.l.Println(query)

    attrQuery := make([]string, 0)

    attributes := query.Get("attrs")
    attributeSplit := strings.Split(attributes, "|")

    // Outer for loop is for each type of attribute (main, sub, third)
    for attrPosition, condition := range attributeSplit {

        // If there is no filter, skip to the next type of attribute
        if condition != noFilter {
            filterArray := make([]string, 0)
            attributeCondition := strings.Split(condition, ",")

            // Inner for loop is for each color of attribute (r,b,g,l,d) 
            for index, attribute := range attributeCondition {

                // If there is no filter, skip to the next attribute color
                if attribute != noFilter {
                    attrNumber := fmt.Sprintf("%d", attrPosition+1)

                    // No attribute is a different value for main attribute only
                    filter := "attr_" + attrNumber + "=" + CorrectAttribute(attribute,attrPosition,index)
                    filterArray = append(filterArray, filter)                    
                }
            }
            singleAttrQuery := AddParanthesis(strings.Join(filterArray, queryOr))
            attrQuery = append(attrQuery, singleAttrQuery)
        }
    }

    attrQueryFinal := strings.Join(attrQuery, queryAnd)

    // If more than one condition, it needs to be wrapped around with
    // paranthesis to keep the condition unique
    if len(attrQueryFinal) > 0 {
        attrQueryFinal = AddParanthesis(attrQueryFinal)
    }

    typeQuery := ""
    cardTypes := query.Get("types")

    // Because types can be in 6 possible combinations, check for all 6 combinations
    // if type filtering is specified. If only one type is requests, it only needs to
    // check 3 combinations since 3 instead of 3x2. That will be for later when this
    // filtering reached an o k point.
    if cardTypes != noFilter {
        filterArray := make([]string, 6)
        cardTypeSplit := strings.Split(cardTypes, ",")

        // Types are initially wildcarded
        type1 := `%`
        type2 := `%`
        type3 := `%`

        // Specify type 1 if requested
        if cardTypeSplit[0] != noFilter {
            type1 = frontTypeToBackType[cardTypeSplit[0]]
        }

        // Specify type 2 if requested
        if len(cardTypeSplit) > 1 {
            if cardTypeSplit[1] != noFilter {
                type2 = frontTypeToBackType[cardTypeSplit[1]]
            }
        }

        // Specify type 3 if requested
        if len(cardTypeSplit) > 2 {
            if cardTypeSplit[2] != noFilter {
                type3 = frontTypeToBackType[cardTypeSplit[2]]
            }
        }

        filterArray[0] = AddParanthesis(typeBase + "'" + type1 + "|" + type2 + "|" + type3 + "'")
        filterArray[1] = AddParanthesis(typeBase + "'" + type1 + "|" + type3 + "|" + type2 + "'")
        filterArray[2] = AddParanthesis(typeBase + "'" + type2 + "|" + type1 + "|" + type3 + "'")
        filterArray[3] = AddParanthesis(typeBase + "'" + type2 + "|" + type3 + "|" + type1 + "'")
        filterArray[4] = AddParanthesis(typeBase + "'" + type3 + "|" + type1 + "|" + type2 + "'")
        filterArray[5] = AddParanthesis(typeBase + "'" + type3 + "|" + type2 + "|" + type1 + "'")

        typeQuery = strings.Join(filterArray, queryOr)
    }

    typesQueryFinal := typeQuery

    // If more than one condition, it needs to be wrapped around with
    // paranthesis to keep the condition unique
    if len(typesQueryFinal) > 0 {
        typesQueryFinal =  AddParanthesis(typesQueryFinal)
    }

    rarityFilter := make([]string, 0)

    rarity := query.Get("rarity")
    raritySplit := strings.Split(rarity, ",")

    // Unlike types, rarity is always an OR condition since one card cannot
    // be of multiple rarities. Join the conditions by OR.
    for _, rarityNumber := range raritySplit {
        if rarityNumber != noFilter {
            filter := "rarity=" + rarityNumber
            rarityFilter = append(rarityFilter, filter)
        }
    }
    rarityQueryFinal :=strings.Join(rarityFilter, queryOr)

    // If more than one condition, it needs to be wrapped around with
    // paranthesis to keep the condition unique
    if len(rarityQueryFinal) > 0 {
        rarityQueryFinal =  AddParanthesis(rarityQueryFinal)
    }
    
    attrTableQuery := JoinStringsNonEmpty([]string {attrQueryFinal, typesQueryFinal, rarityQueryFinal}, queryAnd)

    attrTableQueryFinal := attrTableBase

    // If attrs table query is empty, skip the WHERE clause.
    if len(attrTableQuery) > 0 {
        attrTableQueryFinal = attrTableQueryFinal + whereClause
    }
    attrTableQueryFinal = attrTableQueryFinal + attrTableQuery + ")"

    awknsQuery := make([]string, 0)
    awakenings := query.Get("awkns")

    // Query contains the specific awakening in frontend number that needs to be converted
    // to backend number. Query also contains the count for each awakening so this
    // filtering is very easy to look about it.
    if awakenings != noFilter {
        awakeningSplit := strings.Split(awakenings, "|")

        // For each awakening, check how many a card needs to have and add that to the list
        for _, condition := range awakeningSplit {
            awakeningCondition := strings.Split(condition, ",")
            filter := "awkn_" + ConvertAwakening(awakeningCondition[0]) + ">=" + awakeningCondition[1]
            awknsQuery = append(awknsQuery, filter)
        }
    }

    awknsTableQuery := strings.Join(awknsQuery, queryAnd)
    awknsTableQueryFinal := awknsTableBase

    // If awkns table query is empty, skip the WHERE clause.
    if len(awknsTableQuery) > 0 {
        awknsTableQueryFinal = awknsTableQueryFinal + whereClause
    }
    awknsTableQueryFinal = awknsTableQueryFinal + awknsTableQuery + ")"

    dbQueryFinal := ""
    // This does not need to be long if any of the queries are empty.
    // TODO: Make this not happen if some or all are empty.
    // Now empty if no filters.
    if (attributes != noFilter || cardTypes != noFilter || awakenings != noFilter) {
        dbQueryFinal = fmt.Sprintf(dbQueryBase, awknsTableQueryFinal, attrTableQueryFinal)
    }
    
    skillQuery := make([]string, 0)
    skillType := query.Get("skill")
    skillFilter := ""

    // Very annoying to deal with due to skill combos and skill evos.
    // For now, it deals with singular types only.
    // TODO: implement a way to filter multiple skill types.
    // Initial pass through is needed to parse out the parameters for multi_skills?
    // It may not be necessary but for now it stays.
    if skillType != noFilter {

        // Split by . to get each skill part
        skillSplit := strings.Split(skillType, ".")
        if len(skillSplit) > 0 {
            for _, skill := range skillSplit {
                skillPart := make([]string, 0)
                // Split by | to get each skill type that can represent this
                // skill part. ie. Bind Recovery can be 117 or 179
                skillTypeParam := strings.Split(skill, "|")
                for _, skillType := range skillTypeParam {

                    // Split by - to separate the part from its parameters
                    skillCondition := strings.Split(skillType, "-")
                    paramArray := make([]string, 0)
                    correctedSkillType, _ := strconv.Atoi(skillCondition[0])
                    if (correctedSkillType > 1000) {
                        correctedSkillType -= 1000
                    }
                    typeFilter := fmt.Sprintf(skillTypeUnionBase, strconv.Itoa(correctedSkillType))

                    // If skill part has parameters, add that as a condition for the query
                    if len(skillCondition) > 1 {

                        paramSplit := strings.Split(skillCondition[1], ",")
                        for index, param := range paramSplit {
                            if (index > 7) {
                                continue;
                            }
                            // If *, it can be any so skip to the next parameter to check
                            if param == "*" {
                                continue;
                            }

                            if (skillCondition[0] == "1249") {
                                if (index == 1) {
                                    paramFilter := "param_2 & " + param + " = " + param
                                    paramArray = append(paramArray, paramFilter)
                                    continue;
                                }
                            }

                            if (skillCondition[0] == "1152" || skillCondition[0] == "1190" || skillCondition[0] == "1140") {
                                if (index == 0) {
                                    paramFilter := "param_1 & " + param + " = " + param
                                    paramArray = append(paramArray,paramFilter)
                                    continue;
                                }
                            }

                            if (skillCondition[0] == "127" || skillCondition[0] == "128" || skillCondition[0] == "176") {
                                if (index == 5) {
                                    paramFilter := "param_" + strconv.Itoa(index+1) + param
                                    paramArray = append(paramArray,paramFilter)
                                    continue;
                                }
                                if (skillCondition[0] == "127") {
                                    paramFilter := "param_" + strconv.Itoa(index+1) + " & " + param
                                    paramArray = append(paramArray,paramFilter)
                                    continue;
                                } else {
                                    paramFilter := "param_" + strconv.Itoa(index+1) + " & " + param + " = " + param
                                    paramArray = append(paramArray,paramFilter)
                                    continue;
                                }
                            }

                            if (skillCondition[0] == "1071") {
                                paramFilter := fmt.Sprintf(paramInBase, param, paramColunmn)
                                paramArray = append(paramArray, paramFilter)
                                continue;
                            }

                            if (correctedSkillType == 141) {
                                if (index == 1) {
                                    paramFilter := "param_2 & " + param
                                    paramArray = append(paramArray, paramFilter)
                                    continue;
                                }
                                if (index == 2) {
                                    if (skillCondition[0] == "1141")  {
                                        paramFilter := "NOT param_3 & " + param
                                        paramArray = append(paramArray, paramFilter)
                                        continue;
                                    } else {
                                        paramFilter := "param_3 & " + param
                                        paramArray = append(paramArray, paramFilter)
                                        continue;
                                    }
                                }
                            }

                            if (correctedSkillType == 208) {
                                if (index == 1 || index == 4) {
                                    paramFilter := "param_2 & " + param
                                    paramArray = append(paramArray, paramFilter)
                                    continue;
                                }
                                if (index == 2 || index == 5) {
                                    if (skillCondition[0] == "1208") {
                                        paramFilter := "NOT param_" + strconv.Itoa(index+1) + " & " + param
                                        paramArray = append(paramArray, paramFilter)
                                        continue;
                                    } else {
                                        paramFilter := "param_" + strconv.Itoa(index+1) + " & " + param
                                        paramArray = append(paramArray, paramFilter)
                                        continue;
                                    }
                                }
                            }

                            if (skillCondition[0] == "154") {
                                paramFilter := "param_" + strconv.Itoa(index+1) + " & " + param + " >= " + param
                                paramArray = append(paramArray, paramFilter)
                                continue;
                            }

                            if (index == 2 && skillCondition[0] == "90") {
                                param_2 := paramArray[len(paramArray)-1]
                                paramArray = removeLast(paramArray)
                                paramFilter := "param_" + strconv.Itoa(index+1) + param
                                paramArray = append(paramArray, param_2 + queryOr + paramFilter)
                                continue;
                            }
                            if (index ==2 && skillCondition[0] == "258") {
                                paramFilter := "param_3 & " + param
                                paramArray = append(paramArray, paramFilter)
                                continue;
                            }
                            paramFilter := "param_" + strconv.Itoa(index+1) + param
                            paramArray = append(paramArray, paramFilter)
                        }
                        if (skillCondition[0] == "71") {
                            if (len(paramSplit) < 8) {
                                paramFilter := "param_" + strconv.Itoa(len(paramSplit) + 1) + "=-1"
                                paramArray = append(paramArray, paramFilter)
                            }
                        }

                        if (correctedSkillType == 141) {
                            if len(paramArray) < 2 {
                                if (skillCondition[0] == "1141")  {
                                    paramFilter := "NOT param_2 & param_3"
                                    paramArray = append(paramArray, paramFilter)
                                } else {
                                    paramFilter := "param_2 & param_3"
                                    paramArray = append(paramArray, paramFilter)
                                }
                            }
                        }
                        if (correctedSkillType == 208) {
                            if len(paramArray) < 2 {
                                if (skillCondition[0] == "1208")  {
                                    paramFilter := "NOT param_2 & param_3"
                                    paramArray = append(paramArray, paramFilter)
                                    paramFilter = "NOT param_5 & param_6"
                                    paramArray = append(paramArray, paramFilter)
                                } else {
                                    paramFilter := "param_2 & param_3"
                                    paramArray = append(paramArray, paramFilter)
                                    paramFilter = "param_5 & param_6"
                                    paramArray = append(paramArray, paramFilter)
                                }
                            }
                        }
                    }

                    if (skillCondition[0] == "0") {
                        paramFilter := "param_2>0"
                        paramArray = append(paramArray, paramFilter)
                    }

                    if len(paramArray) > 0 {
                        typeFilter += queryWhere
                        typeFilter += strings.Join(paramArray, queryAnd)
                    }
                    skillPart = append(skillPart, typeFilter)
                }
                skillPartFinal := strings.Join(skillPart, queryUnion)
                skillQuery = append(skillQuery, skillPartFinal)
            }
        }
        // TODO: Consider skill combo skills
        // Done?
    }

    for index, skillQueryPart := range skillQuery {
        if index == 0 {
            skillFilter += fmt.Sprintf(skillTypeBase, skillQueryPart)
        } else {
            skillFilter += " "
            skillFilter += fmt.Sprintf(skillTypeJoin, skillQueryPart, strconv.Itoa(index+2), strconv.Itoa(index+1), strconv.Itoa(index+2))
        }
    }

    skillFinal := ""
    if len(skillFilter) > 0 {
        skillFinal = fmt.Sprintf(skillTableBase, skillFilter)
    }

    combinedFinal := ""
    // Same problem with above, this does not need to be combined if either or
    // is empty to save query time.
    // I think it's okay now?
    if len(skillFinal) > 0 && len(dbQueryFinal) > 0 {
        combinedFinal = fmt.Sprintf(combinedBase, skillFinal, dbQueryFinal)
    } else if len(skillFinal) > 0 {
        combinedFinal = skillFinal
    } else if len(dbQueryFinal) > 0 {
        combinedFinal = dbQueryFinal
    } else {
        json.NewEncoder(w).Encode(cards)
        return
    }

    f.l.Println(combinedFinal)

    // skillResults := make([]models.SkillResults,0)

    results, err := f.db.Query(combinedFinal)
    if err != nil {
        f.l.Println("failed to fetch information from DB", err)
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
