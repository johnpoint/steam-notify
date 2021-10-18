package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"steam-notify/initHelper"
	"steam-notify/initHelper/depend"
)

var httpServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		i := initHelper.Helper{}
		i.AddDepend(new(depend.Api))
		err := i.Init(ctx)
		if err != nil {
			panic(err)
			return
		}

		forever := make(chan struct{})
		<-forever
	},
}
