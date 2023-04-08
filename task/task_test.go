package task

import (
	"flag"
	"math"
	"os"
	"testing"
)

func TestTask(t *testing.T) {
	//ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	rs := &RestResponse{
	//		Code: 0,
	//		Msg:  "success",
	//		Data: map[string]interface{}{
	//			"id": 1001,
	//		},
	//	}
	//	rsBytes, _ := json.Marshal(rs)
	//	fmt.Fprint(w, string(rsBytes))
	//}))
	//defer ts.Close()
	//t.Log(ts.URL)
	//time.Sleep(time.Hour)
	os.Setenv("ENV_LOG_FILE", os.Stdout.Name())
	os.Setenv("ENV_DATA_FILE", "/Users/william/go/src/github.com/sunxiaoning/restful-executor/data1.json")
	flag.Set("url", "http://localhost:8080/api/v1/demo")
	flag.Set("method", "get")
	Run()
}

func TestDemo(t *testing.T) {
	t.Log(math.MaxInt)
}
