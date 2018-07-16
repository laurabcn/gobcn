package sites

import (
	"github.com/antchfx/xmlquery"
	"strings"
	"awesomeProject2/Domain"
	"github.com/satori/go.uuid"
	"awesomeProject2/Application"
)

func GetCategoriesEnglish() {
	url := "http://opendata-ajuntament.barcelona.cat/data/dataset/462e7ea8-aa84-4892-b93f-3bc9ab8e5b4b/resource/8511595f-9d15-42ec-8604-876ca777311f/download"
	doc, err := xmlquery.LoadURL(url)

	if err != nil {
		panic("That's embarrassing...")
	}
	total := doc.FirstChild.SelectAttr("total")

	if strings.Compare("0",total) == 0 {
		return
	}

	//row := xmlquery.Find(doc, "//row")
	labels := xmlquery.Find(doc,"//code2//item")

	var category Domain.Category
	var uniqesCategories map[string]bool
	uniqesCategories = make(map[string]bool)

	for i := 0; i < len(labels); i++ {
		label := labels[i].SelectAttr("label")

		if "" != label && !uniqesCategories[label] && !strings.Contains(label, "Level") {
			uniqesCategories[label] = true

			category = Domain.Category{uuid.Must(uuid.NewV4()), label, "en", true}
			application.AddCategory(&category)
		}
	}
}