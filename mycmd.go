package main

import (
    "os"
    "github.com/codegangsta/cli" // cli.goをインポート
)

func main() {
    app := cli.NewApp()
    app.Name = "mycmd"            // ヘルプを表示する際に使用される
    app.Usage = "print arguments" // ヘルプを表示する際に使用される
    app.Version = "1.2.3"         // ヘルプを表示する際に使用される
    app.Action = func(c *cli.Context) { // コマンド実行時の処理
        println("Number of arguments is", len(c.Args()))
        println("Fisrt argument is", c.Args()[0])
        println("Second argument is", c.Args()[1])
    }
    app.Run(os.Args)
}
