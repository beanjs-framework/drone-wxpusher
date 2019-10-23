package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	// _ "github.com/joho/godotenv/autoload"
)

func main() {

	p := Plugin{}

	// var err error

	p.AppToken = os.Getenv("PLUGIN_APP_TOKEN")
	p.Content = os.Getenv("PLUGIN_CONTENT")
	p.ContentType, _ = strconv.Atoi(os.Getenv("PLUGIN_CONTENT_TYPE"))
	p.Uids = strings.Split(os.Getenv("PLUGIN_UIDS"), ",")
	// p.TopicIds = strings.Split(os.Getenv("PLUGIN_TOPIC_IDS"), ",")
	p.URL = os.Getenv("PLUGIN_URL")

	p.Exec()
}

type Plugin struct {
	AppToken    string   `json:"appToken"`
	Content     string   `json:"content"`
	ContentType int      `json:"contentType"`
	TopicIds    []int    `json:"topicIds,omitempty"`
	Uids        []string `json:"uids"`
	URL         string   `json:"url,omitempty"`
}

func (p *Plugin) Exec() {
	jsonStr, _ := json.Marshal(p)

	fmt.Println(string(jsonStr))

	postUrl := "http://wxpusher.zjiecode.com/api/send/message"

	_, err := http.Post(postUrl, "application/json", bytes.NewBuffer(jsonStr))

	if err != nil {
		os.Exit(-1)
	}

	// fmt.Println(res)
}
