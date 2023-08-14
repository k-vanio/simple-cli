/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"os"

	"github.com/k-vanio/simple-cli/internal/db/models"
	"github.com/spf13/cobra"
)

func newListCmd(categoryModel *models.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples`,
		RunE:  runList(categoryModel),
	}
}

func runList(categoryModel *models.Category) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		rows, err := categoryModel.FindAll()
		if err != nil {
			return err
		}

		json.NewEncoder(os.Stdout).Encode(rows)
		return nil
	}
}

func init() {
	categoryCmd.AddCommand(newListCmd(models.NewCategory(GetDb())))
}
