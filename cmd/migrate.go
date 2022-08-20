/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/setting"
	"github.com/golang-migrate/migrate"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Long:  `Run migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")

		if _, err := os.Stat("./database/migrations"); os.IsNotExist(err) {
			log.Fatal("No migration directory database/migrations found")
		}

		setting.Setup()
		database.Setup()

		m, err := migrate.New("file://database/migrations", database.DSN())

		if err != nil {
			log.Fatal(err)
		}

		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
