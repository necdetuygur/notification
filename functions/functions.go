package functions

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func WriteFile(filename string, content string) {
	mydata := []byte(content)
	err := ioutil.WriteFile(filename, mydata, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func Pr(a string) {
	fmt.Fprintln(os.Stdout, a)
}

func Get(addr string) string {
	resp, err := http.Get(addr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	str := string(body)
	return str + ""
}

func Post(addr string, postBody map[string][]string) string {
	resp, err := http.PostForm(addr, postBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	str := string(body)
	return str + ""
}

func Pause() {
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func Parse(str string, rgx string, key int, clr string) string {
	r, _ := regexp.Compile(rgx)
	arr := r.FindStringSubmatch(str)
	if len(arr) == 0 {
		return "0"
	}
	pri := strings.ReplaceAll(strings.TrimSpace(arr[key]), clr, "")
	return pri
}

func ParseMultiple(str string, rgx string, key int, key2 int, clr string) string {
	r, _ := regexp.Compile(rgx)
	arr := r.FindAllStringSubmatch(str, 999)
	if len(arr) == 0 {
		return "0"
	}
	pri := strings.ReplaceAll(strings.TrimSpace(arr[key][key2]), clr, "")
	return pri
}

func NotificationSend(title string, body string, link string, topic string) string {
	Pr("NotificationSend:")
	Pr("\ttitle: " + title)
	Pr("\tbody: " + body)
	Pr("\tlink: " + link)
	Pr("\ttopic: " + topic)
	url := "https://fcm.googleapis.com/fcm/send"
	notificationIcon := "https://avatars.githubusercontent.com/u/26275074"
	authKey := "AAAAX8bkM-4:APA91bFjL2mYl0BQZCrvNdPtFMaMFz03GtKnT_sli-CP9EpcgFEGWTnDRaUPAnVa8zxRIzf_XBB7gH_whZ1vnWcbdPhQPsy1NuHNsNc3yoSawlQqqnp5r_E9a50MkrnqQpioyJzBeooS"
	var jsonStr = []byte(`
    {
        "to" :  "/topics/` + topic + `",
        "data": {
            "title": "` + title + `",
            "body": "` + body + `",
            "icon": "` + notificationIcon + `",
            "link": "` + link + `"
        }
    }
    `)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "key="+authKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	str := string(bd)
	return str + ""
}

func Minifier(a string) string {
	i := 9
	for i > 0 {
		a = regexp.MustCompile(`(?ms)<!--.*?-->`).ReplaceAllString(a, ` `)
		a = regexp.MustCompile(`  `).ReplaceAllString(a, ` `)
		a = regexp.MustCompile(`\t`).ReplaceAllString(a, ` `)
		a = regexp.MustCompile(`\n`).ReplaceAllString(a, ` `)
		a = regexp.MustCompile(`\r`).ReplaceAllString(a, ` `)
		i--
	}
	return a
}

func ParseInteger(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		return int(0)
	}
	return int(i)
}
