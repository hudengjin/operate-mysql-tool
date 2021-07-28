package driver

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/huprince/operate-mysql-tool/config"
	logger "github.com/huprince/operate-mysql-tool/log"
	"github.com/huprince/operate-mysql-tool/util"
)

var db *sql.DB


func init() {
	env := config.GetEnv()
	host := env.MySqlHost
	port := env.MySqlPort
	username := env.MySqlUsername
	password := env.MySqlPassword
	dbName := env.MySqlDataBase
	var err error
	db, err = sql.Open("mysql", username + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + dbName + "?charset=utf8mb4")
	if err != nil {
		logger.Logger.Panic("Fetch mysql connection error:" + err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func ExecuteDDL(ddl string) {
	_, err := db.Exec(ddl)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
}

func ExcecuteQuery(querySql string, output bool, outputFile, split string) {
	stmt, err := db.Prepare(querySql)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	columns, err := rows.Columns()
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			logger.Logger.Error(err.Error())
			continue
		}
		var v string
		rowValue := make([]string, 0)
		for _, col := range values {
			if col == nil {
				v = "NULL"
			} else {
				v = string(col)
			}
			rowValue = append(rowValue, v)
			// fmt.Println(columns[i], ":", v)
		}
		if output {
			content := strings.Join(rowValue, split) + "\n"
			util.WriteText(outputFile, content, true)
		} else {
			logger.Logger.Sugar().Infof("row data: %s", strings.Join(rowValue, split))
		}

	}
	if err = rows.Err(); err != nil {
		logger.Logger.Fatal(err.Error())
	}
}


func CloseConnection() {
	db.Close()
}