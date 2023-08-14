/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/k-vanio/simple-cli/internal/db/models"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryModel *models.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples`,
		RunE:  runCreate(categoryModel),
	}
}

func runCreate(categoryModel *models.Category) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		if _, err := categoryModel.Create(name, "-"); err != nil {
			return err
		}
		return nil
	}
}

func init() {
	categoryCmd.AddCommand(newCreateCmd(models.NewCategory(GetDb())))
}
