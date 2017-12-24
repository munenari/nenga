package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/munenari/nenga/model"
)

var db *sqlx.DB

func main() {
	initDB()
	sender, err := model.Sender{}.Get(db)
	if err != nil {
		log.Println("failed to load sender information")
		log.Fatalln(err)
	}
	destinations, err := model.Destination{}.All(db)
	if err != nil {
		log.Println("failed to load destinations information")
		log.Fatalln(err)
	}
	nenga := model.Nenga{
		Sender:       *sender,
		Destinations: *destinations,
	}

	funcMap := template.FuncMap{
		"noescape": func(str string) template.HTML {
			return template.HTML(str)
		},
	}
	tpl, err := template.New("nenga.tpl").Funcs(funcMap).ParseFiles("resources/nenga.tpl")
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
}
