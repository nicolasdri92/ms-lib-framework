package framework

type Error struct {
	Codigo    string `json:"codigo"`
	Error     string `json:"error"`
	Timestamp string `json:"timestamp"`
	Code      string `json:"code"`
	Status    int    `json:"status"`
	Type      string `json:"type"`
	Trace     string `json:"trace"`
	Title     string `json:"title"`
	Message   string `json:"message"`
}
