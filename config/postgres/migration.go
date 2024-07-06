package postgres

import (
	"github.com/kevinnaserwan/coursphere-api/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.UserAchievement{},
		&model.Achievement{},
		&model.Mentor{},
		&model.Course{},
		&model.Video{},
		&model.CategoryBook{},
		&model.Book{},
		&model.CategoryCourse{},
	)
	if err != nil {
		panic(err)
	}
}
