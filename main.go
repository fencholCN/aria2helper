package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	argc := len(os.Args)
	if argc < 2 {
		return
	}

	var params []interface{}

	urlStr := os.Args[1]
	var urlList []string
	urlList = append(urlList, urlStr)
	params = append(params, urlList)

	//option
	if argc > 2 {
		options := make(map[string]interface{})
		cookie := "Cookie:" + os.Args[2]
		headerArr := []string{cookie}
		options["header"] = headerArr
		if argc > 3 {
			for i := 3; i < argc; i++ {
				argStr := os.Args[i]
				index := strings.Index(argStr, ":")
				opKey := argStr[:index]
				opValue := argStr[index+1:]
				if len(opKey) > 0 && len(opValue) > 0 {
					options[opKey] = opValue
				}
			}
		}
		params = append(params, options)
	}

	buf, err := json.Marshal(params)
	if err != nil {
		return
	}
	bs := base64.StdEncoding.EncodeToString(buf)
	fin := url.QueryEscape(bs)

	rid := rand.Intn(999)
	idstr := strconv.Itoa(rid)
	resp, err := http.Get("http://127.0.0.1:6800/jsonrpc?method=aria2.addUri&id=" + idstr + "&params=" + fin)

	if err != nil {
		fmt.Println("err is: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
	fmt.Println(string(body))
}
