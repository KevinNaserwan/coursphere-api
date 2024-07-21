package request

type VideoRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Time string `json:"time"`
}
