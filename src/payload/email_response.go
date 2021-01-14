package payload

type EmailRespose struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"errormsg"`
	Result   string `json:"result"`
}
