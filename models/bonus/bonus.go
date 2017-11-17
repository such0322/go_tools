package bonus

import (
	"log"

	"github.com/jmoiron/sqlx"
)

const TypeArtifact = "artifact"
const TypeEP = "ep"
const TypeEXP = "exp"
const TypeGachaTicket = "gacha_ticket"
const TypeGuildCoin = "guild_coin"
const TypeGuildEXP = "guild_exp"
const TypeItem = "item"
const TypeItemBundle = "item_bundle"
const TypeLUPI = "lupi"
const TypeMap = "map"
const TypeMileage = "mileage"
const TypeORB = "orb"
const TypeRandomBox = "random_box"
const TypeRune = "rune"
const TypeToken = "token"
const TypeVIP = "vip"
const TypeMissionItem = "mission_item"
const TypeLetter = "letter"
const TypeChest = "chest"

var RewardType = map[string]string{
	TypeORB:         "水晶",
	TypeLUPI:        "金币",
	TypeMileage:     "稀有勋章",
	TypeArtifact:    "装备",
	TypeRune:        "符石",
	TypeMap:         "地图",
	TypeGachaTicket: "扭蛋券",
	TypeItem:        "道具",
	TypeItemBundle:  "道具包",
	TypeEXP:         "经验值",
}
var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "root:123456@tcp(192.168.7.120:3306)/twod_enish_master_2?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
}

type Bonus struct {
	Type        string
	ID          int
	Quantity    int
	AddContents string
}
