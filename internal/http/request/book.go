package request

type CreateBookRequest struct {
	Title          string `json:"title" binding:"required"`
	Language       string `json:"language" binding:"required"`
	Rank           int    `json:"rank" binding:"required"`
	ReadingTime    int    `json:"reading_time" binding:"required"`
	Likes          int    `json:"likes" binding:"required"`
	BookFile       string `json:"book_file" binding:"required"`
	Overview       string `json:"overview" binding:"required"`
	Writer         string `json:"writer" binding:"required"`
	CategoryBookID string `json:"category_id" binding:"required"`
}
