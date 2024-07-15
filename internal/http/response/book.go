package response

type BookResponse struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title" binding:"required"`
	Language    string                 `json:"language" binding:"required"`
	Rank        int                    `json:"rank" binding:"required"`
	ReadingTime int                    `json:"reading_time" binding:"required"`
	Likes       int                    `json:"likes" binding:"required"`
	BookFile    string                 `json:"book_file" binding:"required"`
	Overview    string                 `json:"overview" binding:"required"`
	Writer      string                 `json:"writer" binding:"required"`
	Category    []BookCategoryResponse `json:"category" binding:"required"`
}
