package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/taverok/proxy-checker-example/service/checker"
	"github.com/taverok/proxy-checker-example/service/checker/proxy/dto"
)

var showContent bool
var proxies []string

var rootCmd = &cobra.Command{
	Use:   "checker [-c] [-p proxies] domain",
	Short: "check target domain under proxy",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := checker.NewApp()
		if err != nil {
			panic(err)
		}
		defer app.Shutdown()

		in := &dto.ProxyRequest{
			Target:      args[0],
			ShowContent: showContent,
			Proxies:     proxies,
		}
		result, err := app.ProxyService.CheckRequest(in)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, r := range result {
			fmt.Println(r.CliRepr())
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showContent, "content", "c", false, "show content")
	rootCmd.Flags().StringSliceVarP(&proxies, "proxy", "p", []string{}, "proxy list")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
