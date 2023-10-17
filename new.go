package request

import (
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/laof/proxy"
)

type Data struct {
	DateTime string
	Nodes    []string
}

const target = "https://github.com/Alv" +
	"in9999" +
	"/new-" +
	"pac" +
	"/wiki/ss%E5%85%8D%E8%B4%B9%E8%B4%A6%E5%8F%B7"

func New() (Data, error) {

	data := Data{}

	html, err := load()

	if err != nil {
		log.Println("start proxy channel")
		html = proxy.Get(target)
	}

	if html == "" {
		return data, errors.New("cannot load content")
	}

	data.DateTime = transTime(getTime(html))
	data.Nodes = append(getNodes(html, "ssr://"), getNodes(html, "ss://")...)
	return data, nil
}

func load() (string, error) {
	res, err := http.Get(target)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getNodes(txt, ss string) []string {
	arr := []string{}
	re := regexp.MustCompile(`<p>` + ss + `(.*?)</p>`)
	matches := re.FindAllStringSubmatch(txt, -1)
	for _, match := range matches {
		if len(match) >= 2 {
			arr = append(arr, removeEmail(ss+match[1]))
		}
	}
	return arr
}

func transTime(dateString string) string {
	const dateTimeFormat = "2006-01-02 15:04:05"
	t, err := time.Parse("2006-01-02T15:04:05Z", dateString)
	if err != nil {
		return ""
	}
	return t.Format(dateTimeFormat)
}

func removeEmail(txt string) string {
	// txt := `xxx:<a href="mailto:33333@qq.com">33333@qq.com</a>cc`

	pattern := `<a href="mailto:(.*?)">(.*?)</a>`
	reg := regexp.MustCompile(pattern)
	res := reg.ReplaceAllString(txt, "$1")
	return res
}

func getTime(txt string) string {
	pattern := `(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z)`
	reg := regexp.MustCompile(pattern)
	return reg.FindString(txt)
}
