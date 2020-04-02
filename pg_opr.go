package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
create table depth (
	id bigserial primary key,
	market varchar(16) not null,
	coin varchar(16) not null,
	data_type varchar(16) not null,
	buy real not null,
	sell real not null,
	trade real not null,
	time timestamp
);

*/
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "kyoyo"
)

type PG_OPR struct {
}

func (pg *PG_OPR) connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	if db, err := sql.Open("postgres", psqlInfo); err != nil {
		fmt.Println(err)
		return nil
	} else {
		return db
	}
}

func (pg *PG_OPR) Insert(gd *GatherData) {
	if db := pg.connect(); db != nil {
		defer db.Close()
		if rows, err := db.Query("INSERT INTO depth (market,coin,data_type,buy,sell,trade,time) VALUES($1,$2,$3,$4,$5,$6,$7);",
			gd.Market, gd.Coin, gd.DataType, gd.Buy, gd.Sell, gd.Trade, gd.Time); err == nil {
			defer rows.Close()
		} else {
			fmt.Println(err)
			defer rows.Close()
		}
	}
}


