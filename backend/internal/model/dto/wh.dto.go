package dto

type WebhookPath struct {
	Path string `json:"path"`
}

type LineWebhook struct {
	Destination string             `json:"destination" `
	Event       []LineWebhookEvent `json:"events"`
}

type LineWebhookEvent struct {
	Type            string                     `json:"type" bson:"type"`
	Mode            string                     `json:"mode" bson:"mode"`
	Timestamp       int64                      `json:"timestamp" bson:"timestamp"`
	Source          *LineWebhookSource         `json:"source" bson:"source"`
	WebhookEventID  string                     `json:"webhookEventId" bson:"webhookEventId"`
	DeliveryContext LineWebhookDeliveryContext `json:"deliveryContext" bson:"deliveryContext"`
	ReplyToken      string                     `json:"replyToken" bson:"replyToken"`
	Message         *LineWebhookMessage        `json:"message" bson:"message"`
}

type LineWebhookSource struct {
	Type    string  `json:"type" bson:"type"`
	UserId  *string `json:"userId" bson:"userId"`
	GroupId *string `json:"groupId" bson:"groupId"`
	RoomId  *string `json:"roomId" bson:"roomId"`
}

type LineWebhookDeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery" bson:"isRedelivery"`
}

type LineWebhookMessage struct {
	Id              *string                          `json:"id" bson:"id"`
	Type            string                           `json:"type" bson:"type"`
	QuoteToken      *string                          `json:"quoteToken" bson:"quoteToken"`
	QuotedMessageId *string                          `json:"quotedMessageId" bson:"quotedMessageId"`
	MarkAsReadToken *string                          `json:"markAsReadToken" bson:"markAsReadToken"`
	Text            *string                          `json:"text" bson:"text"`
	Emojis          []LineMessageEmoji               `json:"emojis" bson:"emojis"`
	Mention         *LineMessageMention              `json:"mention" bson:"mention"`
	ContentProvider *LineMessageImageContentProvider `json:"contentProvider" bson:"contentProvider"`
	ImageSet        *LineMessageImageSet             `json:"imageSet" bson:"imageSet"`
}

type LineMessageEmoji struct {
	Index     int    `json:"index" bson:"index"`
	Length    int    `json:"length" bson:"length"`
	ProductId string `json:"productId" bson:"productId"`
	EmojiId   string `json:"emojiId" bson:"emojiId"`
}

type LineMessageMention struct {
	Mentionees []LineMessageMentionees `json:"mentionees" bson:"mentionees"`
}

type LineMessageMentionees struct {
	Index  int     `json:"index" bson:"index"`
	Length int     `json:"length" bson:"length"`
	Type   string  `json:"type" bson:"type"`
	UserId *string `json:"userId" bson:"userId"`
	IsSelf *bool   `json:"isSelf" bson:"isSelf"`
}

type LineMessageImageContentProvider struct {
	Type               string  `json:"type" bson:"type"`
	OriginalContentUrl *string `json:"originalContentUrl" bson:"originalContentUrl"`
	PreviewImageUrl    *string `json:"previewImageUrl" bson:"previewImageUrl"`
}

type LineMessageImageSet struct {
	Id    *string `json:"id" bson:"id"`
	Index *int    `json:"index" bson:"index"`
	Total *int    `json:"total" bson:"total"`
}

type LinePushMessage struct {
	To      string               `json:"to"`
	Message []LineWebhookMessage `json:"messages"`
}

type Message struct {
	Message struct {
		ID         string `bson:"id"`
		QuoteToken string `bson:"quoteToken"`
	} `bson:"message"`
}
