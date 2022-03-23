/*
Copyright © 2022 AREK OUZOUNIAN arek@arekouzounian.com

*/
package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
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
		Scrape(out)
	},
}

/*
	NOTE FOR FUTURE:
		maybe when you're grabbing from pastes, if the paste isn't found,
		check the wayback machine for the paste?
*/

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
	//scrapeCmd.Flags().BoolP("aggregate-data", "a", false, "Stores the entire text of the paste, rather than just the key")
}

func exists(fname string) bool {
	_, err := os.Stat(fname)
	return !errors.Is(err, os.ErrNotExist)
}
