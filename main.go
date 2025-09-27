package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"nodebus/cli"
	"nodebus/configm"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := cli.Init()

	cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		config_manager := configm.GetManager()
		cfgCenterServer := *cli.UseCfgCenter

		if cfgCenterServer == "" {
			config_manager.LoadJSON()

		} else {
			client := &http.Client{Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}}

			req, err := http.NewRequest("GET", fmt.Sprintf("https://%s", cfgCenterServer), nil)
			if err != nil {
				panic(fmt.Errorf("不能创建请求: %v", err))
			}

			req.Header.Add("Auth", *cli.CfgCenterAuth)
			resp, err := client.Do(req)
			if err != nil {
				req.URL.Scheme = "http"

				resp, err = client.Do(req)
				if err != nil {
					panic(fmt.Errorf("不能请求 cfgcenter 服务器: %v", err))
				}
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					panic(fmt.Errorf("不能读取 cfgcenter 响应"))
				}

				panic(fmt.Errorf("cfgcenter 拒绝了获取请求: %s", string(body)))
			}

			if err := config_manager.LoadJSONFromReader(resp.Body); err != nil {
				panic(fmt.Errorf("不能解析 cfgcenter 响应: %v", err))
			}

		}
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
