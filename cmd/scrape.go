/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
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

		out, err := cmd.Flags().GetString("file")
		if err != nil {
			panic(err)
		}
		if out == DEFAULT_FILE && len(args) >= 1 {
			fmt.Println("Too many arguments!")
			return
		}

		lst, err := runScrape()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		m, err := json.Marshal(ArchiveLinkEntry{
			Number: 0,
			List:   lst,
		})

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		cast := string(m)
		if exists(out) {
			cast = "," + cast
		}

		f, err := os.OpenFile(out, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		if _, err := f.WriteString(cast); err != nil {
			f.Close()
			panic(err)
		}
		if err := f.Close(); err != nil {
			panic(err)
		}

		fmt.Println("output saved to", out)

	},
}

func runScrape() ([]ArchiveLink, error) {
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
	/*
		m, err := json.Marshal(lst)
		if err != nil {
			panic(err)
		}

		return m, nil
	*/
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
	scrapeCmd.Flags().StringP("file", "f", DEFAULT_FILE, "Output file for the scrape to be saved to.")
	scrapeCmd.Flags().BoolP("loop", "l", false, "Sets the scraper to loop, executing once every minute.")
}

func exists(fname string) bool {
	_, err := os.Stat(fname)
	return !errors.Is(err, os.ErrNotExist)
}
