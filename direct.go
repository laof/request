package request

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Direct() string {

	url := []string{
		"aHR0cHM6Ly9naXRodWIuY29tL0FsdmluOTk5O",
		"S9uZXctcGFjL3dpa2kvc3MlRTUlODUlOEQlRT",
		"glQjQlQjklRTglQjQlQTYlRTUlOEYlQjc=",
	}

	host, _ := Decode(strings.Join(url, ""))
	resp, err := http.Get(string(host))

	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return ""
	}

	// target := []string{"p", "label", "span", "strong", "i", "em"}
	var nodes []string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		txt := s.Text()
		if hasPrefix(txt) {

			aquery := s.Find("a")
			if len(aquery.Nodes) >= 0 {
				a, _ := goquery.OuterHtml(aquery)
				txt = strings.Replace(txt, a, aquery.Text(), 1)
			}

			nodes = append(nodes, txt)
		}
	})

	return strings.Join(nodes, "\n")

}

func hasPrefix(text string) bool {

	if strings.HasPrefix(text, "ssr://") || strings.HasPrefix(text, "ss://") {
		return true
	}
	return false

	// buf := string(body)

	// reg := regexp.MustCompile(`</?.+?/?>`)

	// html := reg.ReplaceAllStringFunc(buf, func(str string) string {
	// 	return "\n" + str + "\n"
	// })

	// ssr
	// reg = regexp.MustCompile(`ssr://(.*?)(\s|[\r\n]$)`)

	// txt := reg.FindAllString(html, -1)
	// var nodes []string
	// for _, text := range txt {
	// 	nodes = append(nodes, text)
	// }

	// // ss
	// reg = regexp.MustCompile(`ss://(.*?)(\s|[\r\n]$)`)

	// sstxt := reg.FindAllString(html, -1)
	// for _, txt := range sstxt {
	// 	nodes = append(nodes, txt)
	// }
}
