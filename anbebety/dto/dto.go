package dto

type ReleaseDto struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Aim      string `json:"aim"`
	Time     string `json:"time"`
	Location string `json:"location"`
	Require  string `json:"require"`
	Number   int    `json:"number"`
}
type ApplyDto struct {
	Name       string `json:"name"`
	GroupTitle string `json:"GroupTitle"`
	Reason     string `json:"reason"`
	Advantage  string `json:"advantage"`
}
type CheckDto struct {
	Total   int             `json:"total"`
	Records []CheckResponse `json:"records"`
}
type CheckResponse struct {
	SenderName string `json:"sender_name"`
	Content    string `json:"content"`
}
