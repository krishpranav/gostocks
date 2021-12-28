package main

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/krishpranav/gostocks/api"
	"github.com/krishpranav/gostocks/graph"
	"github.com/piquette/finance-go/datetime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	interval,
	save,
	remove,
	name *string
	year    *bool
	ytd     *bool
	week    *bool
	version *bool
	extra   *bool
	width   *int
	theme   *string
	days    *int

	configPath string
)

func setDefaults() {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		panic("No HOME env var set")
	}
	configPath = fmt.Sprintf("%s/.config", home)
	viper.AddConfigPath(configPath)
	viper.SetConfigName("stonks.yml")
	viper.SetConfigType("yaml")

	viper.SetDefault("favourites", map[string]interface{}{})
	
	viper.SetDefault("config.standalone_height", 12)
	viper.SetDefault("config.favourites_height", 12)
	viper.SetDefault("config.default_theme", graph.LineTheme)

	viper.ReadInConfig()
}

func main() {
	setDefaults()
}