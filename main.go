package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/munenari/nenga/model"
)

var (
	db *sqlx.DB
)

const nengaTpl = "resources/nenga.tpl"

func main() {
	initDB()

	nenga := model.Nenga{}.MustGet(db)

	funcMap := template.FuncMap{
		"noescape": func(str string) template.HTML {
			return template.HTML(str)
		},
	}
	tpl, err := template.New(filepath.Base(nengaTpl)).Funcs(funcMap).ParseFiles(nengaTpl)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nenga)
	if err != nil {
		log.Fatalln(err)
	}
}

func initDB() {
	dbHost := flag.String("dbhost", "localhost", "db host name (default: localhost)")
	dbPort := flag.Int("dbport", 5432, "db port (default: 5432)")
	dbUser := flag.String("dbuser", "postgres", "db username (default: postgres)")
	dbPassword := flag.String("dbpassword", "", "db password")
	dbName := flag.String("dbname", "nenga", "database name (default: nenga)")
	dbTable := flag.String("dbtable", "ab_atena", "database table name (default: ab_atena)")
	dbOptions := flag.String("dboptions", "", "db connection options (key=value), separated with space (default: none)")
	flag.Parse()

	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s",
		*dbHost,
		*dbPort,
		*dbUser,
		*dbPassword,
		*dbName,
		*dbOptions)
	db = sqlx.MustConnect("postgres", dataSource)
	model.AtenaTableName = *dbTable
}
