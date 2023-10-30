/*
练习 4.13： 使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/
上电影的名字和对应的海报图像。编写一个poster工具，通过命令行输入的电影名字，下载对
应的海报。
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
}

/*
YourApiKey
由于没有注册apiKey，代码请求失败
要获得自己的开放电影数据库的 API 密钥，你需要按照以下步骤进行：

	访问 https://www.omdbapi.com/ 并注册一个账户。
	登录后，你可以在 https://www.omdbapi.com/apikey.aspx 页面上找到你的 API 密钥。
	复制你的 API 密钥。
*/
var YourApiKey = "xxxxx"

func main() {
	// 从命令行获取电影名称
	movieName := flag.String("movie", "", "电影名称")
	flag.Parse()

	if *movieName == "" {
		fmt.Println("请输入电影名称")
		return
	}

	// 格式化电影名称，将空格替换为加号
	formattedMovieName := strings.ReplaceAll(*movieName, " ", "+")

	// 构建请求URL
	url := fmt.Sprintf("http://www.omdbapi.com/?t=%s&apikey=%s", formattedMovieName, YourApiKey)

	// 发送HTTP GET请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("请求失败：%v\n", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("search url:", url)
	// 解析JSON响应
	var movie Movie
	err = json.NewDecoder(response.Body).Decode(&movie)
	if err != nil {
		fmt.Printf("解析JSON失败：%v\n", err)
		return
	}

	fmt.Printf("result movie:%v", movie)

	// 下载海报图像
	err = downloadPoster(movie.Poster, movie.Title+".jpg")
	if err != nil {
		fmt.Printf("下载海报失败：%v\n", err)
		return
	}

	fmt.Println("海报下载完成")
}

func downloadPoster(url, filename string) error {
	// 发送HTTP GET请求
	response, err := http.Get(url)
	fmt.Println("poster url:", url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应体写入文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
