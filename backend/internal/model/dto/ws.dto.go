package dto

type ReqWebSocket struct {
	Type           string           `json:"type"`
	QuoteMessageId *string          `json:"quotedMessageId"`
	Timestamp      int64            `json:"timestamp"`
	Data           ReqWebSocketData `json:"data"`
}

type ReqWebSocketData struct {
	Message   *string `json:"message"`
	MessageId *string `json:"messageId"`
	Stage     *string `json:"stage"`
	Status    *string `json:"status"`
}
