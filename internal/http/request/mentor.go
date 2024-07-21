package request

type CreateMentorRequest struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	Experience string `json:"experience"`
}

type UpdateMentorRequest struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Experience string `json:"experience"`
}
