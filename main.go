package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	url_title   string
	url_msg     string
	url_userIds string
	url_url     string
)

func init() {
	url_title = os.Getenv("PLUGIN_TITLE")
	url_msg = os.Getenv("PLUGIN_MSG")
	url_userIds = os.Getenv("PLUGIN_IDS")
	url_url = os.Getenv("PLUGIN_URL")
}

func main() {
	if url_msg == "" || url_userIds == "" {
		log.Fatalln("key or text is required")
	}

	reqMsg := &url.Values{
		"title":   []string{url_title},
		"msg":     []string{url_msg},
		"userIds": []string{url_userIds},
		"url":     []string{url_url},
	}

	reqURL := fmt.Sprintf("http://wxmsg.dingliqc.com/send?%s", reqMsg.Encode())
	fmt.Println(reqURL)
	_, err := http.Get(reqURL)
	if err != nil {
		log.Fatalln("post error: ", err)
	}
}
