package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL          = "http://pastebin.com/archive"
	DEFAULT_FILE = "scrape.json"
)

func Scrape(out string) {
	lst, err := apiCallScrape()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if exists(out) {
		old_data, err := ioutil.ReadFile(out)
		if err != nil {
			fmt.Printf("Error reading file '%s': %s\n", out, err.Error())
			return
		}

		var prev ArchiveList
		jerr := json.Unmarshal(old_data, &prev)
		if jerr != nil {
			panic(err)
		}

		for idx, item := range prev {
			if lst[idx] == item {
				continue
			}
			lst = append(lst, item)
		}

	}

	m, err := json.Marshal(lst)
	if err != nil {
		panic(err)
	}

	werr := os.WriteFile(out, m, 644)
	if werr != nil {
		fmt.Printf("Error writing to file '%s': %s\n", out, werr.Error())
	}

	fmt.Println("output saved to", out)
}

func apiCallScrape() (ArchiveList, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	rows := doc.Find("td")
	lst := make([]ArchiveLink, 0)
	currKey := ""
	rows.Each(func(i int, s *goquery.Selection) {
		key, exists := s.Find("a").Attr("href")
		if exists && key != "" {
			if !strings.Contains(key, "archive") {
				currKey = key
			} else {
				lst = append(lst, ArchiveLink{
					Key:            currKey,
					GroupDirectory: key,
				})
			}
		}
	})

	return lst, nil
}
