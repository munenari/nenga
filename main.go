package main

import (
	"html/template"
	"log"
	"os"

	"github.com/munenari/nenga/model"
)

func main() {

	nenga := model.Nenga{
		Destinations: []model.Destination{
			{
				Names: []string{
					"山田　太郎　様",
					"　　　花子　様",
				},
				Postcode: "1234567",
				Address:  "鹿児島県鹿児島市○○町一―一<br>桜島マンション一〇一号",
			},
		},
		Sender: model.Sender{
			Names: []string{
				"山田　一郎",
				"　　　二郎",
			},
			Postcode: "1234568",
			Address:  "鹿児島市○○町一―二<br>霧島コーポ二〇二",
		},
	}

	tpl, err := template.ParseFiles("resources/nenga.tpl")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nenga)
	if err != nil {
		log.Fatalln(err)
	}
}