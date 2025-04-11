package models

type Card struct {
    ID              int `json:"id"`
    NAME            string `json:"name"`
    RARITY          int `json:"rarity"`
    COST            int `json:"cost"`
	OFFICIAL        int `json:"official"`
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
    ASSIST          int
    EXPAND          int
    BASE_ID         int
    GROUP_ID        int
    MP              int
    COLLAB_ID       int
    KEYWORDS        string
    ORB_BGM_ID      int
}