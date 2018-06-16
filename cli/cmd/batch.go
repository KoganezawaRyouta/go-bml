package cmd

import (
	"github.com/KoganezawaRyouta/go-bml/batche"
	"github.com/spf13/cobra"
)

var (
	speak string
)
var speakCmd = &cobra.Command{
	Use:   "speak_world",
	Short: "Golangを思い出す",
	Long:  "Golangを思い出す",
	Run: func(cmd *cobra.Command, args []string) {
		batche.NewHello(&conf, speak).Run()
	},
}

// フラグの値を変数にセットする場合
// 第1引数: 変数のポインタ
// 第2引数: フラグ名
// 第3引数: デフォルト値
// 第4引数: 説明
func init() {
	// サブコマンドのフラグ定義
	speakCmd.PersistentFlags().StringVar(&speak, "speak", "デフォルト値", "your speaking")
	RootCmd.AddCommand(speakCmd)
}
