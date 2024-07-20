package response

type CourseResponse struct {
	BannerImage string                   `json:"banner_image"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Star        int                      `json:"star"`
	Price       int                      `json:"price"`
	Lessons     int                      `json:"lessons"`
	Mentor      []MentorResponse         `json:"mentor"`
	Category    []CourseCategoryResponse `json:"category"`
	Videos      []VideoResponse          `json:"videos"`
}
