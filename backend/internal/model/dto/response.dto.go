package dto

type ResponseBody struct {
	Message string `json:"message"`
}

type PushMessageResponse struct {
	SentMessages []SentMessage `json:"sentMessages"`
}

type SentMessage struct {
	Id         string `json:"id"`
	QuoteToken string `json:"quoteToken"`
}
