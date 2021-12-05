package request

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type workExample struct {
	URL string `json:"url"`
}

type jsonTypeData struct {
	workExample `json:"workExample"`
}

func Proxy() string {

	url := []string{
		"aHR0cHM6Ly9jb2Rlc2FuZGJveC5pby9zL2",
		"dpdGh1Yi9sYW9mL3Nzc3NhbmRib3g=",
	}

	host, _ := Decode(strings.Join(url, ""))
	res, fail := http.Get(string(host))

	if fail != nil {
		return ""
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return ""
	}

	scripts := doc.Find("script")
	if len(scripts.Nodes) <= 0 {
		return ""
	}

	var jsonStr string

	scripts.Each(func(i int, s *goquery.Selection) {

		val, ok := s.Attr("type")

		if !ok {
			return
		}

		if val == "application/ld+json" {
			jsonStr = s.Text()
		}

	})

	if jsonStr == "" {
		return ""
	}

	jsonData := jsonTypeData{}
	json.Unmarshal([]byte(jsonStr), &jsonData)

	view := jsonData.workExample.URL

	if !strings.HasPrefix(view, "http") {
		return ""
	}

	nodes := PDBody(view)

	return nodes

}
