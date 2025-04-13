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
    ATTR_1                      int `json:"attr_1"`
    ATTR_2                      int `json:"attr_2"`
    ATTR_3                      int `json:"attr_3"`
    TYPE_1                      int `json:"type_1"`
    TYPE_2                      int `json:"type_2"`
    TYPE_3                      int `json:"type_3"`
    RARITY                      int `json:"rarity"`
    COST                        int `json:"cost"`
    IS_ULT_EVO                  int `json:"ult_evo"`
    MAX_LEVEL                   int `json:"max_level"`
    FUSE_EXP                    int `json:"fuse_exp"`
    COIN_VALUE                  int `json:"coin_value"`
    OFFICIAL                    int `json:"official"`
    HP_INIT                     int `json:"hp_init"`
    HP_MAX                      int `json:"hp_max"`
    HP_GROWTH                   int `json:"hp_growth"`
    ATK_INIT                    int `json:"atk_init"`
    ATK_MAX                     int `json:"atk_max"`
    ATK_GROWTH                  int `json:"atk_growth"`
    RCV_INIT                    int `json:"rcv_init"`
    RCV_MAX                     int `json:"rcv_max"`
    RCV_GROWTH                  int `json:"rcv_growth"`
    MAX_EXP                     int `json:"max_exp"`
    EXP_GROWTH                  int `json:"exp_growth"`
    ACTIVE_SKILL                int `json:"active"`
    LEADER_SKILL                int `json:"leader"`
    EVOLVED_FROM                int `json:"from"`
    EVOLVE_MAT_1                int `json:"mat_1"`
    EVOLVE_MAT_2                int `json:"mat_2"`
    EVOLVE_MAT_3                int `json:"mat_3"`
    EVOLVE_MAT_4                int `json:"mat_4"`
    EVOLVE_MAT_5                int `json:"mat_5"`
    DEVOLVE_MAT_1               int `json:"dmat_1"`
    DEVOLVE_MAT_2               int `json:"dmat_2"`
    DEVOLVE_MAT_3               int `json:"dmat_3"`
    DEVOLVE_MAT_4               int `json:"dmat_4"`
    DEVOLVE_MAT_5               int `json:"dmat_5"`
    AWAKENINGS                  int `json:"awkns"`
    AWAKENING_1                 int `json:"awkn_1"`
    AWAKENING_2                 int `json:"awkn_2"`
    AWAKENING_3                 int `json:"awkn_3"`
    AWAKENING_4                 int `json:"awkn_4"`
    AWAKENING_5                 int `json:"awkn_5"`
    AWAKENING_6                 int `json:"awkn_6"`
    AWAKENING_7                 int `json:"awkn_7"`
    AWAKENING_8                 int `json:"awkn_8"`
    AWAKENING_9                 int `json:"awkn_9"`
    SUPER_AWAKENINGS            int `json:"super_awkns"`
    SUPER_AWAKENING_1           int `json:"s_awkn_1"`
    SUPER_AWAKENING_2           int `json:"s_awkn_2"`
    SUPER_AWAKENING_3           int `json:"s_awkn_3"`
    SUPER_AWAKENING_4           int `json:"s_awkn_4"`
    SUPER_AWAKENING_5           int `json:"s_awkn_5"`
    SUPER_AWAKENING_6           int `json:"s_awkn_6"`
    SUPER_AWAKENING_7           int `json:"s_awkn_7"`
    SUPER_AWAKENING_8           int `json:"s_awkn_8"`
    SUPER_AWAKENING_9           int `json:"s_awkn_9"`
    SUPER_AWAKENING_10          int `json:"s_awkn_10"`
    CAN_ASSIST                  int `json:"assist"`
    CAN_EXPAND                  int `json:"expand"`
    BASE_ID                     int `json:"base"`
    GROUP_ID                    int `json:"group"`
    MONSTER_POINTS              int `json:"mp"`
    COLLAB_ID                   int `json:"collab"`
    SEARCH_KEYWORDS             string `json:"keywords"`
    LIMIT_BREAK_STAT_PERCENT    int `json:"limit_break"`
    ORBSKIN_BGM_ID              int `json:"orb_bgm_id"`
}
