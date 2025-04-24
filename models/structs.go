package models

type Card struct {
    ID              int `json:"id"`
    NAME            string `json:"name"`
}

type Attrs struct {
    ATTR_1          int
    ATTR_2          int
    ATTR_3          int
}

type Types struct {
    TYPE_1          int
    TYPE_2          int
    TYPE_3          int
}

type Stat struct {
    MAX_LEVEL       int
    HP_INIT         int
    HP_MAX          int
    HP_GROWTH       int
    ATK_INIT        int
    ATK_MAX         int
    ATK_GROWTH      int
    RCV_INIT        int
    RCV_MAX         int
    RCV_GROWTH      int
    LIMIT_PERCENT   int
    MAX_EXP         int
    EXP_GROWTH 		int
    FUSE_EXP        int
    COIN_VALUE      int
}

type Evolution struct {
    IS_ULT          int
    FROM            int
    EMAT_1          int
    EMAT_2          int
    EMAT_3          int
    EMAT_4          int
    EMAT_5          int
    DEMAT_1         int
    DEMAT_2         int
    DEMAT_3         int
    DEMAT_4         int
    DEMAT_5         int
}

type Skills struct {
    ACTIVE          int
    LEADER          int
}

type Awakenings struct {
    AWAKENINGS      int
    A_1             int
    A_2             int
    A_3             int
    A_4             int
    A_5             int
    A_6             int
    A_7             int
    A_8             int
    A_9             int
}

type SuperAwakenings struct {
    S_AWAKENINGS    int
    SA_1            int
    SA_2            int
    SA_3            int
    SA_4            int
    SA_5            int
    SA_6            int
    SA_7            int
    SA_8            int
    SA_9            int
    SA_10           int
}

type KeyValues struct {
    RARITY          int
    COST            int
    OFFICIAL        int
    ASSIST          int
    EXPAND          int
    BASE_ID         int
    GROUP_ID        int
    MP              int
    COLLAB_ID       int
    KEYWORDS        string
    ORB_BGM_ID      int
}

type QueryResults struct {
    ID                          int `json:"id"`
    NAME                        string `json:"name"`
    ATTRS                       string `json:"attrs"`
    TYPES                       string `json:"types"`
    RARITY                      int `json:"rarity"`
    COST                        int `json:"cost"`
    ULT_EVO                     int `json:"ult_evo"`
    MAX_LEVEL                   int `json:"max_level"`
    FUSE_EXP                    int `json:"fuse_exp"`
    COIN_VALUE                  int `json:"coin_value"`
    OFFICIAL                    int `json:"official"`
    HP_VALS                     string `json:"hp_vals"`
    ATK_VALS                    string `json:"atk_vals"`
    RCV_VALS                    string `json:"rcv_vals"`
    EXP_VALS                    string `json:"exp_vals"`
    ACTIVE_SKILL                int `json:"active"`
    LEADER_SKILL                int `json:"leader"`
    EVOLVED_FROM                int `json:"from"`
    EVOLVE_MATS                 string `json:"mats"`
    DEVOLVE_MATS                string `json:"dmats"`
    AWAKENINGS                  string `json:"awkns"`
    SUPER_AWAKENINGS            string `json:"super_awkns"`
    ASSIST                      int `json:"assist"`
    EXPAND                      int `json:"expand"`
    BASE_ID                     int `json:"base"`
    GROUP_ID                    int `json:"group"`
    MONSTER_POINTS              int `json:"mp"`
    COLLAB_ID                   int `json:"collab"`
    SEARCH_KEYWORDS             string `json:"keywords"`
    LIMIT_BREAK_PERCENT         int `json:"limit_break"`
    ORB_BGM_ID                  int `json:"orb_bgm_id"`
    SYNC_AWAKENING              int `json:"sync_awkn"`
    SYNC_MATS                   string `json:"sync_mats"`
}

type SkillResults struct {
    ID              int `json:"id"`
    NAME            string `json:"name"`
    TEXT            string `json:"text"`
    SKILL_TYPE      int `json:"skill_type"`
    SKILL_MAX_LEVEL int `json:"skill_max_level"`
    SKILL_INIT_CD   int `json:"skill_init_cd"`
    UNKNOWN_1       string `json:"unknown_1"`
    PARAMETERS      string `json:"parameters"`
}
