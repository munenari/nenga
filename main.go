package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/munenari/nenga/model"
	"golang.org/x/text/unicode/norm"
)

var (
	db       *sqlx.DB
	replacer = map[string]string{
		// "1":  "一",
		// "2":  "二",
		// "3":  "三",
		// "4":  "四",
		// "5":  "五",
		// "6":  "六",
		// "7":  "七",
		// "8":  "八",
		// "9":  "九",
		// "0":  "〇",
		"1":  "１",
		"2":  "２",
		"3":  "３",
		"4":  "４",
		"5":  "５",
		"6":  "６",
		"7":  "７",
		"8":  "８",
		"9":  "９",
		"0":  "０",
		"一":  "１",
		"二":  "２",
		"三":  "３",
		"四":  "４",
		"五":  "５",
		"六":  "６",
		"七":  "７",
		"八":  "８",
		"九":  "９",
		"〇":  "０",
		" ":  "　",
		"-":  "―",
		"−":  "―",
		"\n": "<br>　",
	}
	replacerAlphabet = map[string]string{}
)

const nengaTpl = "resources/nenga.tpl"

func init() {
	for s := 'A'; s <= 'Z'; s++ {
		replacerAlphabet[string(s)] = string('Ａ' + (s - 'A'))
	}
}

func main() {
	initDB()

	nenga := model.Nenga{}.MustGet(db)

	funcMap := template.FuncMap{
		"noescape": func(str string) template.HTML {
			return template.HTML(str)
		},
		"normtategaki": replaceKanSuji,
		"parseaddress": func(str string) template.HTML {
			return template.HTML(replaceKanSuji(str))
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

func replaceKanSuji(str string) string {
	str = norm.NFKC.String(str)
	for old, new := range replacer {
		str = strings.ReplaceAll(str, old, new)
	}
	for old, new := range replacerAlphabet {
		str = strings.ReplaceAll(str, old, new)
	}
	return str
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
