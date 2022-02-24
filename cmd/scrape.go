/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

const (
	URL = "http://pastebin.com/archive"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Scrape,
}

func Scrape(cmd *cobra.Command, args []string) {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rows := doc.Find("td")
	lst := make([]ArchiveLink, 0)
	currKey := ""
	rows.Each(func(i int, s *goquery.Selection) {
		key, exists := s.Find("a").Attr("href")
		if exists {
			if key != "" {
				if !strings.Contains(key, "archive") {
					currKey = key
				} else {
					lst = append(lst, ArchiveLink{
						Key:            currKey,
						GroupDirectory: key,
					})
				}
			}
		}
	})

	m, err := json.Marshal(lst)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("scrape.json", m, 0666); err != nil {
		panic(err)
	}
	fmt.Printf("Output saved to 'scrape.json'\n")
}

func init() {
	rootCmd.AddCommand(scrapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
