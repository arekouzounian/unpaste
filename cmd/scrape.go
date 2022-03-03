/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

const (
	URL          = "http://pastebin.com/archive"
	DEFAULT_FILE = "scrape.json"
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
	Run: func(cmd *cobra.Command, args []string) {

		out, err := cmd.Flags().GetString("output")
		if err != nil {
			panic(err)
		}
		if out == DEFAULT_FILE && len(args) >= 1 {
			out = args[0]
		}

		lst, err := runScrape()
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

		werr := os.WriteFile(out, m, 0644)
		if werr != nil {
			fmt.Printf("Error writing to file '%s': %s\n", out, werr.Error())
		}

		fmt.Println("output saved to", out)
	},
}

func runScrape() (ArchiveList, error) {
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

func init() {
	rootCmd.AddCommand(scrapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	scrapeCmd.Flags().StringP("output", "o", DEFAULT_FILE, "Output file for the scrape to be saved to.")
	scrapeCmd.Flags().BoolP("loop", "l", false, "Sets the scraper to loop, executing once every minute.")
	scrapeCmd.Flags().BoolP("aggregate-data", "a", false, "Stores the entire text of the paste, rather than just the key")
}

func exists(fname string) bool {
	_, err := os.Stat(fname)
	return !errors.Is(err, os.ErrNotExist)
}
