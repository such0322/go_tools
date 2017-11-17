package misc

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//可重复使用账号
const TYPE_REPEAT_USE = 1

//单批次只可使用一次
const TYPE_ONLY_ONE = 2

//单批次可使用多个
const TYPE_MULTI_USE = 3

const CODE_LEN = 10
const STATUS_OPEN = 1

const QUANTITY_DEFAULT = 1

func GetRandomCode() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	rs := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < CODE_LEN; i++ {
		rs = append(rs, bytes[r.Intn(len(bytes))])
	}
	return string(rs)
}

func CreateGiftCodes(data map[string]interface{}) {
	switch data["type"] {
	case TYPE_REPEAT_USE:
		m := GiftCode{}
		m.createRepeat(data)
	case TYPE_ONLY_ONE:
		fallthrough
	case TYPE_MULTI_USE:
		m := GiftCodes{}
		m.createBatch(data)
	}
}

type GiftCode struct {
	ID          int
	Code        string
	Batch       int
	Channel     string
	Type        int
	Quantity    int
	Package     string
	Status      int
	StartDate   string `db:"start_date"`
	EndDate     string `db:"end_date"`
	LastModDate string `db:"last_mod_date"`
	InsDate     string `db:"ins_date"`
}

func (m *GiftCode) GetOne() {
	err := db.Get(m, "select * from gift_code limit 1")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (m *GiftCode) createRepeat(data map[string]interface{}) int64 {
	if v, ok := data["code"]; !ok || v == "" {
		data["code"] = GetRandomCode()
	}
	raNum, err := m.create(data)
	if err != nil {
		log.Fatal(err)
	}
	return raNum
}

func (m *GiftCode) create(data map[string]interface{}) (int64, error) {
	sql := fmt.Sprintf("insert into gift_code (code,batch,channel,type,quantity,package,status,last_mod_date,ins_date)values (:code,:batch,:channel,:type,:quantity,:package,:status,:last_mod_date,:ins_date)")

	t := time.Now()
	if v, ok := data["package"]; !ok || v == "" {
		data["package"] = "[]"
	}
	if data["type"] != TYPE_ONLY_ONE {
		data["quantity"] = QUANTITY_DEFAULT
	}
	data["batch"] = t.Unix()
	data["status"] = STATUS_OPEN
	data["last_mod_date"] = t.Format("2006-01-02 15:04:05")
	data["ins_date"] = t.Format("2006-01-02 15:04:05")
	fmt.Println(t)

	rs, err := db.NamedExec(sql, data)
	if err != nil {
		log.Fatal(err)
	}
	return rs.RowsAffected()
}

type GiftCodes struct {
	Data []GiftCode
}

func (m *GiftCodes) GetAll() (data []GiftCode) {
	err := db.Select(&data, "select * from gift_code")
	if err != nil {
		log.Fatal(err)
	}
	m.Data = data
	return
}

func (m *GiftCodes) createBatch(data map[string]interface{}) int64 {
	var codes []string
	// 生成codes[]
	for i := 0; i < data["quantity"].(int); i++ {
		code := GetRandomCode()
		if !inCodes(codes, code) {
			codes = append(codes, code)
		}
	}

	//排除已存在的codes
	query, args, err := sqlx.In("SELECT * FROM gift_code WHERE code IN (?);", codes)
	if err != nil {
		log.Println(err)
	}
	db.Select(&m.Data, query, args...)

	for _, row := range m.Data {
		codes = codesUnset(codes, row.Code)
	}

	//数据库处理
	sql := "insert into gift_code (code,batch,channel,type,quantity,package,status,last_mod_date,ins_date) values "
	t := time.Now()
	if v, ok := data["package"]; !ok || v == "" {
		data["package"] = "[]"
	}
	if data["type"] != TYPE_ONLY_ONE {
		data["quantity"] = QUANTITY_DEFAULT
	}
	data["batch"] = t.Unix()
	data["status"] = STATUS_OPEN
	data["last_mod_date"] = t.Format("2006-01-02 15:04:05")
	data["ins_date"] = t.Format("2006-01-02 15:04:05")
	var allArgs []interface{}
	for _, code := range codes {
		data["code"] = code
		query, args, err := sqlx.Named("(:code,:batch,:channel,:type,:quantity,:package,:status,:last_mod_date,:ins_date)", data)
		if err != nil {
			log.Fatal(err)
		}
		sql = sql + query + ", "
		allArgs = append(allArgs, args...)
	}
	sqlr := []rune(sql)
	len := len(sqlr)
	sql = string(sqlr[:len-2])
	rs := db.MustExec(sql, allArgs...)
	raNum, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return raNum
}

func codesUnset(codes []string, code string) []string {
	var i int
	for k, v := range codes {
		if v == code {
			i = k
			break
		}
	}
	return append(codes[:i], codes[i+1:]...)

}

func inCodes(codes []string, code string) bool {
	for _, v := range codes {
		if v == code {
			return true
		}
	}
	return false
}
