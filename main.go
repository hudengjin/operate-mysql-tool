package main

import (
	"flag"
	"os"

	"github.com/huprince/operate-mysql-tool/driver"
	log "github.com/huprince/operate-mysql-tool/log"
	"github.com/huprince/operate-mysql-tool/util"

	_ "github.com/huprince/operate-mysql-tool/config"
)

var q string
var p string
var o bool
var path string
var split string
var f string

const (
	opCreate = "create"
	opInsert = "insert"
	opUpdate = "update"
	opDelete = "delete"
	opQuery  = "query"
)

func init() {
	flag.StringVar(&q, "sql", "", "-sql=?, sql clause option: select, update, insert, delete, create")
	flag.StringVar(&p, "op", "insert", "-op=?, operations: create, insert, update, delete, query")
	flag.BoolVar(&o, "output", false, "-output=?, result output: true, false, 1, 0")
	flag.StringVar(&path, "path", "./output.txt", "-path=?, select clause output file path: ./output.txt ")
	flag.StringVar(&split, "split", ",", "-split=?, output file split char: , "+" \\t"+" ")
	flag.StringVar(&f, "f", "", "-f=?, sql clause as input file: ./test.sql")
}

func main() {
	flag.Parse()

	if q == "" && f == "" {
		log.Logger.Info("Please input query or ddl sql clause!")
		os.Exit(1)
	}
	var sql string
	var err error
	if f != "" {
		sql, err = util.ReadText(f)
		if err != nil {
			log.Logger.Error(err.Error())
			os.Exit(1)
		}
	} else {
		sql = q
	}

	if split == "t" {
		split = "\t"
	}

	log.Logger.Sugar().Infof("Execute operate: %s ", p)

	log.Logger.Info(sql)

	switch p {
	case opCreate, opDelete, opInsert, opUpdate:
		driver.ExecuteDDL(sql)
		driver.CloseConnection()
	case opQuery:
		driver.ExcecuteQuery(sql, o, path, split)
		driver.CloseConnection()
	default:
		log.Logger.Info("Invalid operate, please execute -f option looking for help!")
		os.Exit(1)
	}
}
