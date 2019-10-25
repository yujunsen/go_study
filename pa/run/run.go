package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	savefile = "BoolList.txt"
	//workNum    = 1  //工作任务
	controlnum = 1  //控制进程
	buffersize = 10 //通道缓存
)

var (
	wg sync.WaitGroup
	wo sync.WaitGroup
	//m  sync.Mutex
)

func init() {
	if _, err := os.Stat(savefile); os.IsNotExist(err) {
		_, err := os.Create(savefile)
		if err != nil {
			log.Fatalln("os.Create err", err)
		}
	}
	rand.Seed(time.Now().UnixNano())
}

func main1() {
	c := make(chan map[string]interface{})
	url := `https://www.bukebook.cn/wp-content/plugins/ordown/down.php?id=9861`
	//url := `https://www.bukebook.cn/wp-content/plugins/ordown/down.php?id=12773`
	getresult(url, c)
}

func main() {

	var start, end int
	fmt.Print("起始页(>=1):")
	fmt.Scan(&start)
	fmt.Print("终止页(>=起始页[适度爬取,太多小心IP被封哦]):")
	fmt.Scan(&end)
	if start < 1 || end < 1 || start > end {
		fmt.Println("please inpput start > 1 end > 1 end > 1")
		return
	}
	workNum := end - start + 1
	wo.Add(workNum)
	wg.Add(controlnum)
	n := make(chan map[string]interface{}, buffersize)
	fmt.Println("开始爬虫")
	for i := start; i <= end; i++ {
		go work(n, i)
	}
	for i := 0; i < controlnum; i++ {
		go write(n)
	}
	wo.Wait()
	close(n)
	wg.Wait()
	//close(n)
	fmt.Println("爬取完毕,马上闪开!!!")
}

func work(c chan map[string]interface{}, index int) {
	//defer close(c)
	defer wo.Done()

	url := "https://www.bukebook.cn/page/" + strconv.Itoa(index)
	result, err := httpGet(url)
	if err != nil {
		fmt.Println("检查网络,或者IP被封了...")
		return
	}
	//正则处理信息获得bookID
	//<a href="https://www.bukebook.cn/12748.html" rel="bookmark">
	bookIDRule := `<a href="https://www.bukebook.cn/([0-9]+).html" rel="bookmark">`
	//<a style="color:#999" href="https://www.bukebook.cn/12770.html">《皇帝圆舞曲--从启蒙到日落的欧洲》 高林（作者）azw3+epub+mobi</a>
	//<a class="ordown-button" href="https://tc5.us/file/17456484-403840784" target="_blank">城通网盘</a>
	//bookNameRule := `.html">《(?s:(.*?))</a></h2>`
	allID := regexpData(result, bookIDRule)
	//https://www.bukebook.cn/wp-content/plugins/ordown/down.php?id=12770
	for _, tmpID := range allID {
		bookID := tmpID[1]
		dlUrl := "https://www.bukebook.cn/wp-content/plugins/ordown/down.php?id=" + bookID
		// dlResult, err := httpGet(dlUrl)
		// if err != nil {
		// 	fmt.Println("dl httpGet err", err)
		// 	return
		// }
		//处理数据获取书名,下载地址及密码

		//break
		//println(dlUrl)
		getresult(dlUrl, c)
	}
}

func write(c chan map[string]interface{}) {
	defer wg.Done()
	for {
		resp, ok := <-c
		if !ok {
			return
		}
		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			log.Fatalln("err :", err)
		}

		fd, err := os.OpenFile(savefile, os.O_APPEND, 0666)
		if err != nil {
			return
		}

		fd.Write(data)
		fd.WriteString("\n")
		fd.Close()
	}
}

func httpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if err != nil {
			err = err2
			return
		}
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}
	return
}

func regexpData(data, rule string) [][]string {
	reg := regexp.MustCompile(rule)
	return reg.FindAllStringSubmatch(data, -1)
}

func getresult(url string, c chan map[string]interface{}) {
	resp, err := httpGet(url)
	if err != nil {
		log.Printf("httpGet url %s err %d\n", url, err)
		return
	}
	bookNameRule := `.html">(《.*?)</a></h2>`
	//panNameRule := `href="(.*?)">百度云盘</a>`

	keyNameRule := `提取秘钥： </strong>(.*?)</br>`
	result := regexpData(resp, bookNameRule)
	if len(result) == 0 {
		return
	}
	//fmt.Println(result[0][1])
	file := make(map[string]interface{})
	bookName := result[0][1]
	file["name"] = bookName

	if strings.Contains(resp, "百度云盘") {
		panNameRule := `ordown-button" href="(.*?)">百度云盘`
		result = regexpData(resp, panNameRule)
		file["url"] = `https://www.bukebook.cn/wp-content/plugins/ordown/` + result[0][1]
		file["pan"] = "百度云盘"
	} else {

		panNameRule := `ordown-button" href="(.*?)".*target="_blank">(.*?)</a>`
		result = regexpData(resp, panNameRule)
		file["pan"] = result[0][2]
		file["url"] = result[0][1]

	}
	result = regexpData(resp, keyNameRule)
	file["key"] = result[0][1]
	c <- file
	//fmt.Println(file)
}

//target="_blank">(.*?)</a>
//<a class="ordown-button" href="https://tc5.us/file/17456484-403564518" target="_blank">城通网盘</a>
//<a class="ordown-button" href="download1.php?id=9861">百度云盘</a>
// //删除函数
// func remove(s []string, i int) []string {
// 	return append(s[:i], s[i+1:]...)
// }
