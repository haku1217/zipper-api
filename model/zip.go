package model

// package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Zip struct {
	ID                 int    `json:"id" db:"id"`
	LocalGovemmentCode int    `json:"local_government_code" db:"local_government_code"`
	PrefectureCode     int    `json:"prefecture_code" db:"prefecture_code"`
	ZipCode            int    `json:"zip_code" db:"zip_code"`
	PrefectureKana     string `json:"prefecture_kana" db:"prefecture_kana"`
	Prefecture         string `json:"prefecture" db:"prefecture"`
	CityKana           string `json:"city_kana" db:"city_kana"`
	City               string `json:"city" db:"city"`
	TownKana           string `json:"town_kana" db:"town_kana"`
	Town               string `json:"town" db:"town"`
}

type ZipList []Zip

type ZipRepository interface {
	findAll() ZipList
	findByZipCode(code int) Zip
	store(zip Zip) error
}

type ZipDaoImpl struct {
	Zip Zip
	DB  *sqlx.DB
}

// DSN is [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func NewMySQL(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	const MYSQL_HOST = "127.0.0.1"
	const MYSQL_TCP_PORT = "13322"
	const MYSQL_PWD = "root"
	const MYSQL_USER = "user"
	const DB_NAME = "sample_db"
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USER, MYSQL_PWD, MYSQL_HOST, MYSQL_TCP_PORT, DB_NAME)
	fmt.Println(DSN)

	db, err := NewMySQL(DSN)
	if err != nil {
		log.Fatal(err)
	}
	dao, _ := NewZipRepository(db)
	fmt.Println(dao.findByZipCode(2300000))
	zip := Zip{ID: 1, LocalGovemmentCode: 14101, PrefectureCode: 230, ZipCode: 2300000, PrefectureKana: "ｶﾅｶﾞﾜｹﾝ", Prefecture: "神奈川県", CityKana: "ﾖｺﾊﾏｼﾂﾙﾐｸ", City: "横浜市鶴見区", TownKana: "ｲｶﾆｹｲｻｲｶﾞﾅｲﾊﾞｱｲ", Town: "以下に掲載がない場合"}
	dao.store(zip)
}

func NewZipRepository(db *sqlx.DB) (ZipRepository, error) {
	return &ZipDaoImpl{DB: db}, nil
}

func (dao ZipDaoImpl) findAll() ZipList {
	var zipList ZipList
	err := dao.DB.Select(&zipList, "SELECT * FROM zip")
	if err != nil {
		log.Fatal(err)
	}
	return zipList
}

func (dao ZipDaoImpl) findByZipCode(code int) Zip {
	err := dao.DB.QueryRowx("SELECT * FROM zip WHERE zip_code=? LIMIT 1", code).StructScan(&dao.Zip)
	if err != nil {
		log.Fatal(err)
	}
	return dao.Zip
}

func (dao ZipDaoImpl) store(zip Zip) (err error) {
	zipState := `INSERT INTO zip (local_government_code,prefecture_code,zip_code,prefecture_kana,prefecture,city_kana,city,town_kana,town) 
				VALUES(:local_government_code,:prefecture_code,:zip_code,:prefecture_kana,:prefecture,:city_kana,:city,:town_kana,:town)`
	_, err = dao.DB.NamedExec(zipState, zip)
	return err
}
