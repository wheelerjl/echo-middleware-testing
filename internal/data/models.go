package data

const (
	EchoCtxRespKey = "ctx_resp_key"
)

type InternalError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data any
	Code int
}

type WebFruit struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Extra string `json:"extra"`
}

type DBFruit struct {
	ID   int
	Name string
}
