package screepsws

type ConsoleMessage struct {
	Log     []string `json:"log"`
	Results []string `json:"results"`
}

type MessageUpdate struct {
	ID         string           `json:"_id"`
	OutMessage string           `json:"outMessage"`
	Text       string           `json:"text"`
	Direction  MessageDirection `json:"type"`
	Unread     bool             `json:"unread"`
	User       string           `json:"user"`
	Respondent string           `json:"respondent"`
}

type MessageReadUpdate struct {
	ID     string `json:"_id"`
	Unread bool   `json:"unread"`
}

type NewMessage struct {
	Message MessageUpdate `json:"message"`
}

type MessageRead struct {
	Message MessageReadUpdate `json:"message"`
}
