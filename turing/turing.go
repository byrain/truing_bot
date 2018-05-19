package turing

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SetTuringBot(apiKey, userID, turingURL string) {
	APIKEY = apiKey
	USERID = userID
	TuringURL = turingURL
}

type TuringMessageReq struct {
	ReqType    int           `json:"reqType"`
	Perception PerceptionReq `json:"perception"`
	UserInfo   UserInfoReq   `json:"userInfo"`
}

type PerceptionReq struct {
	InputText InputTextReq `json:"inputText"`
}

type InputTextReq struct {
	Text string `json:"text"`
}

type UserInfoReq struct {
	ApiKey string `json:"apiKey"`
	UserId string `json:"userId"`
}

func NewTuringMessage(content string) TuringMessageReq {
	userInfo := UserInfoReq{
		ApiKey: APIKEY,
		UserId: USERID,
	}
	perception := PerceptionReq{InputTextReq{content}}
	return TuringMessageReq{
		ReqType:    0,
		Perception: perception,
		UserInfo:   userInfo,
	}
}

type TuringMessageResp struct {
	Intent IntentResp   `json:"intent"`
	Result []ResultResp `json:"results"`
}

type IntentResp struct {
	Code       int                    `json:"code"`
	IntentName string                 `json:"intentName"`
	ActionNaem string                 `json:"ActionName"`
	Parameters map[string]interface{} `json:"Parameters"`
}

type ResultResp struct {
	GroupType  int               `json:"groupType"`
	ResultType string            `json:"resultType"`
	Values     map[string]string `json:"values"`
}

func GetTuringBotResp(message TuringMessageReq) TuringMessageResp {
	r, _ := json.Marshal(message)
	request, err := http.NewRequest("POST", TuringURL, bytes.NewReader(r))
	if err != nil {
		log.Println("http.NewRequest,[err=%s][url=%s]", err, TuringURL)
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, TuringURL)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, TuringURL)
	}

	turingResp := TuringMessageResp{}
	err = json.Unmarshal(b, &turingResp)
	if err != nil {
		log.Println(err)
	}
	return turingResp
}
