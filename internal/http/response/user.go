package response

type UserResponse struct {
	ID              string                    `json:"id"`
	Username        string                    `json:"username"`
	Email           string                    `json:"email"`
	Profession      string                    `json:"profession"`
	UserAchievement []UserAchievementResponse `json:"user_achievement"`
}

type UserAchievementResponse struct {
	ID          string                `json:"id"`
	UserID      string                `json:"user_id"`
	Achievement []AchievementResponse `json:"achievement"`
}

type AchievementResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
