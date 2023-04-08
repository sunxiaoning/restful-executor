package task

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	url    string
	method string
	query  string
	data   string
)
var (
	timeout uint
)

var (
	methods = map[string]bool{http.MethodGet: true, http.MethodPost: true, http.MethodPut: true, http.MethodDelete: true}
)

func Run() {
	flag.Parse()
	if url == "" {
		log.Fatal("task url not specified!")
	}
	if method == "" {
		log.Fatal("task method not specified!")
	}
	if !methods[strings.ToUpper(method)] {
		log.Fatal("task method is invalid!")
	}
	if (method == http.MethodPost || method == http.MethodPut) && data == "" {
		log.Fatal("post,put method should with data!")
	}
	logFileEnv, dataFileEnv := os.Getenv(EnvLogFile), os.Getenv(EnvDataFile)
	if logFileEnv == "" {
		log.Fatal("logFile not specified!")
	}
	if dataFileEnv == "" {
		log.Fatal("dataFile not specified!")
	}
	logFile, err := os.OpenFile(logFileEnv, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("open log file: " + err.Error())
	}
	defer logFile.Close()
	dataFile, err := os.OpenFile(dataFileEnv, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("open data file: " + err.Error())
	}
	defer dataFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println("request info: ")
	logger.Println("task: ", url)
	logger.Println("method: ", method)
	logger.Println("data: ", data)
	logger.Println("timeout: ", timeout)
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}
	var req *http.Request
	httpMethod := strings.ToUpper(method)
	if method != http.MethodPost && method != http.MethodPut {
		req, err = http.NewRequest(httpMethod, url, nil)
	} else {
		req, err = http.NewRequest(httpMethod, url, strings.NewReader(data))
	}
	if err != nil {
		logger.Fatal("prepare request: " + err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal("do http request: " + err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Println("read http resp body: " + err.Error())
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logger.Printf("http request failed, http_code=%d, resp_body=%s", resp.StatusCode, string(body))
	}
	var rs RestResponse
	err = json.Unmarshal(body, &rs)
	if err != nil {
		logger.Fatalf("parse resp body: " + err.Error())
	}
	if rs.Code != CodeSucceed {
		logger.Printf("http request failed,code=%d, resp=%s", rs.Code, rs.Msg)
	}
	if rs.Data != nil {
		outputBytes, err := json.Marshal(rs.Data)
		if err != nil {
			logger.Fatal("parse rest resp data: " + err.Error())
		}
		output := string(outputBytes)
		_, err = io.WriteString(dataFile, output)
		if err != nil {
			logger.Fatal("save output failed: " + err.Error())
		}
	}
}

func init() {
	flag.StringVar(&url, "url", "", "resource task")
	flag.StringVar(&method, "method", "", "request method.")
	flag.StringVar(&data, "data", "", "request data")
	flag.UintVar(&timeout, "timeout", 1, "request timeout seconds")
}
