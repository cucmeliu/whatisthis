package utils

import (

	//	"bytes"
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func httpGet(url string) (response string) {

	resp, err := http.Get(url)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		// handle error

	}

	return string(body)

}

func HttpPost(urlstr string, params map[string]string) (response string, err error) {

	paramstr := ""
	for k, v := range params {
		paramstr += k + "=" + url.QueryEscape(v) + "&"
	}

	// 去掉最后一个 & 符号
	paramstr = paramstr[:len(paramstr)-1]

	// Go里一个巨大的坑，特殊字符UrlEncode后，%后的是大写的，不是小写的。。。。
	c := []byte(paramstr)
	for i := 0; i < len(c); i++ {
		if rune(c[i]) == '%' {
			if c[i+2] >= 65 && c[i+2] <= 90 {
				c[i+2] = c[i+2] + 32
			}
		}
	}

	str := string(c)

	resp, err := http.Post(urlstr, "application/x-www-form-urlencoded", strings.NewReader(str))

	if err != nil {
		fmt.Println("post err", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	return string(body), err

}

func httpPostForm() {

	resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {

		// handle error

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		// handle error

	}

	fmt.Println(string(body))

}

//复杂的请求
//有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo() {

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))

	if err != nil {

		// handle error

	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		// handle error

	}

	fmt.Println(string(body))

}

//同上面的post请求，必须要设定Content-Type为application/x-www-form-urlencoded，post参数才可正常传递。
//如果要发起head请求可以直接使用http client的head方法，比较简单，这里就不再说明。
//Tips：使用这个方法的话，第二个参数要设置成”application/x-www-form-urlencoded”，否则post参数无法传递。

func main() {
	fmt.Println("Hello World!")
	httpGet("http://www.ikohoo.com/")
}
