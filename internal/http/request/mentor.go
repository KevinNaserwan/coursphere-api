package request

type CreateMentorRequest struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	Experience string `json:"experience"`
}
