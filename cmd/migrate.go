/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/setting"

	"github.com/golang-migrate/migrate"

	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Long:  `Run migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("./database/migrations"); os.IsNotExist(err) {
			log.Fatal("No migration directory database/migrations found")
		}

		setting.Setup()

		log.Println("Connecting to : " + database.DURL())

		m, err := migrate.New("file://database/migrations", database.DURL())

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
