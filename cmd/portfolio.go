/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"

	"github.com/Gituser143/cryptgo/pkg/api"
	"github.com/Gituser143/cryptgo/pkg/display/portfolio"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// portfolioCmd represents the portfolio command
var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Track your portfolio",
	Long:  `The portfolio command helps track your own portfolio in real time`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Context and errgroup used to manage routines
		eg, ctx := errgroup.WithContext(context.Background())
		dataChannel := make(chan api.PortfolioData)

		// Fetch coin data
		eg.Go(func() error {
			return nil
		})

		// Display UI for coins
		eg.Go(func() error {
			return portfolio.DisplayPortfolio(ctx, dataChannel)
		})

		if err := eg.Wait(); err != nil {
			if err.Error() != "UI Closed" {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portfolioCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portfolioCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
