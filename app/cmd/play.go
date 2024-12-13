package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/ip2region"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {

	region := ip2region.Search("183.129.226.96")

	fmt.Println(region)

}
