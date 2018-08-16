package main

import (
	xmlquery "github.com/antchfx/xmlquery"
	"strings"
	"github.com/laurabcn/gobcn/Domain"
	"github.com/satori/go.uuid"
	"github.com/laurabcn/gobcn/Application"
	"os"
	//"fmt"
)

func selectUrl(language string) string {
	var url = ""

	switch language {
		case "en":
			url = "http://opendata-ajuntament.barcelona.cat/data/dataset/462e7ea8-aa84-4892-b93f-3bc9ab8e5b4b/resource/8511595f-9d15-42ec-8604-876ca777311f/download"
		case "es":
			url = "http://opendata-ajuntament.barcelona.cat/data/dataset/462e7ea8-aa84-4892-b93f-3bc9ab8e5b4b/resource/554094ce-7c08-4e5e-97aa-2bb39b12ced1/download"
		case "ca":
			url = "http://opendata-ajuntament.barcelona.cat/data/dataset/462e7ea8-aa84-4892-b93f-3bc9ab8e5b4b/resource/69a48d9b-a606-4c31-859b-67724b7377f2/download"
		case "fr":
			url = "http://www.bcn.cat/tercerlloc/pits_opendata_fr.xml"

	}
	return url
}

func main() {

	languages := os.Args[1:]

	for i, language := range languages {
		url := selectUrl(language)

		doc, err := xmlquery.LoadURL(url)
		i++
		if err != nil {
			panic("That's embarrassing...")
		}
		total := doc.FirstChild.SelectAttr("total")

		if strings.Compare("0",total) == 0 {
			return
		}

		categories := xmlquery.Find(doc,"//code2//item")
		setCategories(categories, language)

		row := xmlquery.Find(doc, "//row")


		var site Domain.Site

		for i:=0; i<len(row);i++{
			site = Domain.site(
				uuid.Must(uuid.NewV4()),
				xmlquery.Find(doc,"//title"),
				language,
				true,
				xmlquery.Find(doc,"//district"),
				xmlquery.Find(doc, "//phonenumber"),
				xmlquery.Find(doc,"//code_url"),
				xmlquery.Find(doc,"//content"),
				xmlquery.Find(doc,"//excerpt"),
				xmlquery.Find(doc, "//gmapx"),
				xmlquery.Find(doc, "//gmapy"),
				xmlquery.Find(doc, "//type"),
				xmlquery.Find(doc, "//addresses//item//barri"),
				xmlquery.Find(doc, "//addresses//item//address"),
			)

			Application.AddSite(&site)

		}


	}

}

func setCategories(categories, language string) {
	var category Domain.Category
	var uniqesCategories map[string]bool
	uniqesCategories = make(map[string]bool)
	for i := 0; i < len(categories); i++ {
		label := categories[i].SelectAttr("label")

		if "" != label && !uniqesCategories[label] && !strings.Contains(label, "Level") {
			uniqesCategories[label] = true
			category = Domain.Category{uuid.Must(uuid.NewV4()), label, language, true}
			Application.AddCategory(&category)
		}
	}
}
