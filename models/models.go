package models

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	// _ "github.com/lib/pq"
)

var db *sql.DB
var Lists StaticLists

//ConnectDB - Connect to the database initially.  The connection remains open.
// func ConnectDB() {
// 	dbinfo := fmt.Sprintf("port=5432 user=%s password=%s dbname=%s sslmode=disable", "postgres", "andy", "postgres")
// 	var err error
// 	db, err = sql.Open("postgres", dbinfo)
// 	if err != nil {
// 		fmt.Printf("err: %+v ", err)
// 	}
// 	Lists.LoadLists()
// }

func ConnectDB() {
	var port *int = flag.Int("port", 1433, "the database port")
	var server = flag.String("server", "goirisdev.database.windows.net", "the database server")
	var user = flag.String("user", "doronesk", "the database user")
	var database = flag.String("database", "myfp", "The database")
	var password = flag.String("password", "1Superman", "the password")
	var connString = fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", *server, *database, *user, *password, *port)

	var err error
	db, err = sql.Open("mssql", connString)
	if err != nil {
		fmt.Printf("err: %+v ", err)
	}
	Lists.LoadLists()
}

//StaticLists - Read in from Database and will get refreshed if changed
type StaticLists struct {
	TransCategories []TransCategory
	TransTypes      []TransType
}

// TransType - Transaction Category structure
type TransType struct {
	ID   string `json:"ID"`
	Type string `json:"type"`
}

func (T *StaticLists) GetTransTypes() ([]TransType, error) {
	stmt, err := db.Prepare(`select id, name, type_id from myfp.trans_category`)
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to get type list: %v", err)
	}

	var typ []TransType
	for rows.Next() {
		var t TransType
		err := rows.Scan(&t.ID, &t.Type)
		if err != nil {
			fmt.Println("err: ", err)
			return typ, fmt.Errorf("failed to get transactions list: %v", err)
		}
		typ = append(typ, t)
	}

	return typ, nil
}

// TransCategory - Transaction Category structure
type TransCategory struct {
	ID     string `json:"ID"`
	Name   string `json:"name"`
	TypeID string `json:"type_ID"`
}

func (T *StaticLists) GetTransCategories() ([]TransCategory, error) {
	stmt, err := db.Prepare(`select id, name from trans_category`)
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to get categories list: %v", err)
	}

	var cat []TransCategory
	for rows.Next() {
		var t TransCategory
		err := rows.Scan(&t.ID, &t.Name)
		if err != nil {
			fmt.Println("err: ", err)
			return cat, fmt.Errorf("failed to get transactions list: %v", err)
		}
		cat = append(cat, t)
	}

	return cat, nil
}
func (T *StaticLists) LoadLists() {
	T.TransCategories, _ = T.GetTransCategories()
	T.TransTypes, _ = T.GetTransTypes()
}

//MoneyTrans - Money transaction structure
type MoneyTrans struct {
	TransID          string `json:"trans-id"`
	TransDate        string `json:"trans-date"`
	TransAmount      string `json:"trans-amount"`
	TransCategory    string `json:"trans_category"`
	TransType        string `json:"trans_type"`
	TransDescription string `json:"trans_description"`
}

//GetAllTrans - Retreieve all transactions for the last n days.
func (T *MoneyTrans) GetAllTrans() ([]MoneyTrans, error) {
	// stmt, err := db.Prepare(`select tm.id, to_char(tm.date,'MM/DD/YYYY'), tm.amount, tm.category_id, tm.type_id, tm.description
	stmt, err := db.Prepare(`select tm.id, convert(varchar(12),tm.date,9) , tm.amount, tm.category_id,  tm.description from trans_money tm`)
	rows, err := stmt.Query()

	if err != nil {
		return nil, fmt.Errorf("failed to get transactions list: %v", err)
	}

	var trans []MoneyTrans
	for rows.Next() {
		p, err := T.readTrans(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to get transaction list: %v", err)
		}
		trans = append(trans, p)
	}

	return trans, nil
}

func (T *MoneyTrans) readTrans(rows *sql.Rows) (MoneyTrans, error) {
	var t MoneyTrans
	err := rows.Scan(&t.TransID, &t.TransDate, &t.TransAmount, &t.TransCategory, &t.TransDescription)
	if err != nil {
		fmt.Println("err: ", err)
		return MoneyTrans{}, fmt.Errorf("failed to get transactions list: %v", err)
	}

	return t, nil
}
