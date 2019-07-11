package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	//  _ "github.com/joho/godotenv/autoload"
)


func main() {

	e_remark := os.Getenv("PLUGIN_REMARK")
	e_uids := os.Getenv("PLUGIN_IDS")
	e_title:=os.Getenv("PLUGIN_TITLE")

	if e_remark == "" || e_uids == "" {
		log.Fatalln("key or text is required")
	}

	reqMsg := &url.Values{
		"title": []string{e_title},
		"msg": []string{e_remark},
		"userIds": []string{e_uids},
	}

	reqURL := "http://wxmsg.dingliqc.com/send?"

	_, err := http.Get(reqURL+reqMsg.Encode())
	if err != nil {
		log.Fatalln("post error: ", err)
	}
}
