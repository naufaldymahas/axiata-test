package payload

type responseHandler struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseWithData(d interface{}) responseHandler {
	return responseHandler{Status: 200, Message: "", Data: d}
}

func ResponseWithMessage(s int, m string) responseHandler {
	return responseHandler{Status: s, Message: m, Data: nil}
}
