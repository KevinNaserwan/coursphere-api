package request

type CreateCourseRequest struct {
	BannerImage      string `json:"banner_image"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	MentorID         string `json:"mentor_id"`
	CategoryCourseID string `json:"category_course_id"`
	Star             int    `json:"star"`
	Price            int    `json:"price"`
	Lessons          int    `json:"lessons"`
	CategoryID       string `json:"category_id"`
}
