package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/IrisAdminApi/backend/config"
	"github.com/snowlyg/ffmpegTest"
	"time"
)

func main() {

	go func() {
		f := NewLogFile()
		defer f.Close()

		api := NewApp()
		//api.Logger().SetOutput(f) //记录日志

		if config.Config.HTTPS {
			host := fmt.Sprintf("%s:%d", config.Config.Host, 443)
			if err := api.Run(iris.TLS(host, config.Config.Certpath, config.Config.Certkey)); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := api.Run(
				iris.Addr(fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port)),
				iris.WithoutServerError(iris.ErrServerClosed),
				iris.WithOptimizations,
				iris.WithTimeFormat(time.RFC3339),
			); err != nil {
				fmt.Println(err)
			}
		}
	}()

	go func() {
		fmt.Println("hls start")
		//inFilename := "rtsp://183.59.168.27/PLTV/88888905/224/3221227272/10000100000000060000000001030757_0.smil?icip=88888888"
		inFilename := "rtmp://58.200.131.2:1935/livetv/hunantv"
		ffmpegTest.ToHls(inFilename, config.Config.RecordPath, "tcp")
		fmt.Println("hls end")
	}()

	select {}

}
