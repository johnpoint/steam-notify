package telegram

type BasicMessage struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type ForwardFrom struct {
	LastName  string `json:"last_name"`
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Type      string `json:"type"`
	Title     string `json:"title"`
}

type ReplyToMessage struct {
	Date      int    `json:"date"`
	Chat      Chat   `json:"chat"`
	MessageId int    `json:"message_id"`
	Text      string `json:"text"`
}

type Message struct {
	Date           int            `json:"date"`
	Chat           Chat           `json:"chat"`
	MessageId      int            `json:"message_id"`
	From           From           `json:"from"`
	Text           string         `json:"text"`
	ForwardFrom    ForwardFrom    `json:"forward_from"`
	ForwardDate    int            `json:"forward_date"`
	ReplyToMessage ReplyToMessage `json:"reply_to_message"`
	EditDate       int            `json:"edit_date"`
	Entities       []*Entities    `json:"entities"`
	Audio          Audio          `json:"audio"`
	Voice          Voice          `json:"voice"`
	Document       Document       `json:"document"`
}

type Voice struct {
	FileId   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

type Audio struct {
	FileId   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
	Title    string `json:"title"`
}

type Document struct {
	FileId   string `json:"file_id"`
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

type Entities struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type Chat struct {
	LastName  string `json:"last_name"`
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type From struct {
	LastName  string `json:"last_name"`
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}
