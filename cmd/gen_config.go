package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"steam-notify/config"
)

var genConfigCommand = &cobra.Command{
	Use:   "genConfig",
	Short: "Generate config",
	Run: func(cmd *cobra.Command, args []string) {
		newConfig := config.ServiceConfig{
			MongoDBConfig: &config.MongoDBConfig{},
		}
		jsonByte, _ := json.Marshal(&newConfig)
		fmt.Println(string(jsonByte))
	},
}
