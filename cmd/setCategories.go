package main

import (
	"github.com/antchfx/xmlquery"
	"strings"
	"github.com/laurabcn/gobcn/Domain"
	"github.com/satori/go.uuid"
	"github.com/laurabcn/gobcn/Application"
	"os"
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
		var category Domain.Category
		var uniqesCategories map[string]bool
		uniqesCategories = make(map[string]bool)
		for i := 0; i < len(categories); i++ {
			label := categories[i].SelectAttr("label")

			if "" != label && !uniqesCategories[label] && !strings.Contains(label, "Level") {
				uniqesCategories[label] = true
				category = Domain.Category{uuid.Must(uuid.NewV4()), label, language, true}
				Application.Add(&category)
			}
		}

		var site Domain.Site

		for _, n:= range xmlquery.Find(doc, "//row"){
			var district string
			var phone string
			var web string
			var title string
			var barri string
			var content string

			if xmlquery.FindOne(n, "//district") != nil {
				district =  xmlquery.FindOne(n, "//district").InnerText()
			}else {
				district = "0"
			}

			if xmlquery.FindOne(n, "//phonenumber") != nil {
				phone =  xmlquery.FindOne(n, "//phonenumber").InnerText()
			}else {
				phone = "0"
			}


			if xmlquery.FindOne(n, "//code_url") != nil {
				web =  xmlquery.FindOne(n, "//code_url").InnerText()
			}else {
				web = "0"
			}

			if xmlquery.FindOne(n, "//title") != nil {
				title =  xmlquery.FindOne(n, "//title").InnerText()
			}else {
				title = "0"
			}

			if xmlquery.FindOne(n, "//content") != nil {
				content =  xmlquery.FindOne(n, "//content").InnerText()
			}else {
				content = "0"
			}

			var excerpt = xmlquery.FindOne(n, "//excerpt").InnerText()

			var long = xmlquery.FindOne(n, "//gmapx").InnerText()

			var lat = xmlquery.FindOne(n, "//gmapy").InnerText()

			var typeSite = xmlquery.FindOne(n, "//type").InnerText()


			if xmlquery.FindOne(n, "//addresses//item//barri") != nil {
				barri =  xmlquery.FindOne(n, "//addresses//item//barri").InnerText()
			}else {
				barri = "0"
			}

			var address = xmlquery.FindOne(n, "//addresses//item//address").InnerText()
			var position = xmlquery.FindOne(n, "//pos").InnerText()

			site = Domain.Site{
				uuid.Must(uuid.NewV4()),
				title,
				language,
				true,
				district,
				phone,
				web,
				content,
				excerpt,
				lat,
				long,
				typeSite,
				barri,
				address,
				position,
			}

			Application.AddSite(&site)

			for _, x := range xmlquery.Find(n, "//code2//item"){
				var categorySite = Domain.Category{uuid.Must(uuid.NewV4()), x.SelectAttr("label"), language, true}
				Application.AddSiteCategory(&site, &categorySite)
			}
		}
	}
}
