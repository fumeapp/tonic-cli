package main

import (
	"github.com/fumeapp/tonic-cli/cmd"
	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/setting"
)

func init() {
	setting.Setup()
	database.Setup()
}

func main() {
	cmd.Execute()
}
