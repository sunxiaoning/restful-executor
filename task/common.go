package task

const (
	EnvLogFile  = "ENV_LOG_FILE"
	EnvDataFile = "ENV_DATA_FILE"
)

type RestResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	CodeSucceed = 0
)
