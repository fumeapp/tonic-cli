package main

import (
	"github.com/fumeapp/tonic-cli/cmd"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cmd.Execute()
}
