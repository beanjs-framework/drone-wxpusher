package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	// _ "github.com/joho/godotenv/autoload"
)

type JValue struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type JData struct {
	First    *JValue `json:"first"`
	Keyword1 *JValue `json:"keyword1"`
	Keyword2 *JValue `json:"keyword2"`
	Remark   *JValue `json:"remark"`
}

type JBody struct {
	UserIds    []string `json:"userIds"`
	TemplateId string   `json:"template_id"`
	Url        string   `json:"url"`
	Data       JData    `json:"data"`
}

func main() {

	e_remark := os.Getenv("PLUGIN_REMARK")
	e_uids := os.Getenv("PLUGIN_IDS")
	e_subtitle := os.Getenv("PLUGIN_SUBTITLE_TEXT")
	e_subcolor := os.Getenv("PLUGIN_SUBTITLE_COLOR")

	cstZone := time.FixedZone("CST", 8*3600)

	if e_remark == "" || e_uids == "" {
		log.Fatalln("key or text is required")
	}

	if e_subtitle == "" {
		e_subtitle = "简单提醒"
	}

	if strings.Contains(e_subtitle, "警告") {
		e_subcolor = "#FF4500"
	} else if strings.Contains(e_subtitle, "发布") {
		e_subcolor = "#07C160"
	} else if e_subcolor == "" {
		e_subcolor = "#C0C0C0"
	}

	body := &JBody{}
	body.UserIds = strings.Split(e_uids, ",")
	body.TemplateId = "6gxAPiVQ5fruViymf5azMurvNJ8wZZqwcnprI2RKpGY"
	body.Url = os.Getenv("PLUGIN_URL")
	body.Data.First = &JValue{
		Value: os.Getenv("PLUGIN_TITLE"),
		Color: "#000000",
	}
	body.Data.Remark = &JValue{
		Value: e_remark,
		Color: "#000000",
	}
	body.Data.Keyword2 = &JValue{
		Value: time.Now().In(cstZone).Format("2006-01-02 15:04:05"),
		Color: "#C0C0C0",
	}
	body.Data.Keyword1 = &JValue{
		Value: e_subtitle,
		Color: e_subcolor,
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		log.Fatalln("key or text is required")
	}

	reqURL := "http://wxmsg.dingliqc.com/send"
	_, err = http.Post(reqURL, "application/json", strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatalln("post error: ", err)
	}
}
