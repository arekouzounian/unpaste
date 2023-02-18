/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// loopCmd represents the loop command
var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		// create an infinite loop that calls scrape
		// must have a timer to only call on a set interval
		// are channels necessary?
		out, err := cmd.Flags().GetString("output")
		if err != nil {
			panic(err)
		}
		if out == DEFAULT_FILE && len(args) >= 1 {
			out = args[0]
		}

		intv, err := cmd.Flags().GetUint("interval")
		if err != nil {
			panic(err)
		}
		if intv <= 1 {
			fmt.Println("Interval too small! Cannot scrape with this much frequency.")
			return
		}

		fmt.Println("Starting to loop...")

		for {
			fmt.Println("Scraping data...")
			Scrape(out)
			fmt.Println("Next scrape in ", intv, " minutes.")
			time.Sleep(time.Duration(intv) * time.Minute)
		}
	},
}

func init() {
	rootCmd.AddCommand(loopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loopCmd.Flags().StringP("output", "o", DEFAULT_FILE, "Output file for the scrape ot save to.")
	loopCmd.Flags().UintP("interval", "i", 5, "The time interval in which pastebin is scraped continously. Cannot be less than one minute.")
}
