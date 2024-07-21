package request

type AddCourseVideosRequest struct {
	CourseID string `json:"course_id"`
	VideoID  string `json:"video_id"`
}
