package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid()"`
	Username        string            `gorm:"unique;not null;column:username"`
	Email           string            `gorm:"unique;not null;column:email"`
	Password        string            `gorm:"not null;column:password"`
	Profession      string            `gorm:"column:profession;default:null"`
	AuthCode        string            `gorm:"column:auth_code"`
	IsVerified      bool              `gorm:"column:is_verified;default:false"`
	UserAchievement []UserAchievement `gorm:"many2many:user_achievement"`
}

type UserAchievement struct {
	gorm.Model
	ID          uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid()"`
	UserID      uuid.UUID     `gorm:"column:user_id;type:uuid;default:gen_random_uuid()"`
	User        User          `gorm:"foreignKey:UserID;references:ID"`
	Achievement []Achievement `gorm:"many2many:achievement"`
}

type Achievement struct {
	gorm.Model
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	UserAchievementID uuid.UUID `gorm:"column:user_achievement_id;type:uuid;default:gen_random_uuid()"`
	Name              string    `gorm:"column:name"`
	Description       string    `gorm:"column:description"`
}

type Mentor struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name       string    `gorm:"column:name"`
	Image      string    `gorm:"column:image"`
	Experience string    `gorm:"column:experience"`
	Course     []Course  `gorm:"many2many:course"`
}

type CategoryCourse struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name   string    `gorm:"column:name"`
	Course []Course  `gorm:"many2many:course"`
}

type Course struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()"`
	BannerImage string         `gorm:"column:banner_image"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
	MentorID    uuid.UUID      `gorm:"column:id;type:uuid;default:gen_random_uuid()"`
	Videos      []Video        `gorm:"many2many:video"`
	Star        int            `gorm:"column:star"`
	Price       int            `gorm:"column:price"`
	Lessons     int            `gorm:"column:lessons"`
	CategoryID  uuid.UUID      `gorm:"column:category_id;type:uuid;default:gen_random_uuid()"`
	Category    CategoryCourse `gorm:"foreignKey:CategoryID;references:ID"`
	Mentor      Mentor         `gorm:"foreignKey:MentorID;references:ID"`
}

type Video struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CourseID uuid.UUID `gorm:"column:course_id;type:uuid;default:gen_random_uuid()"`
	Name     string    `gorm:"column:name"`
	URL      string    `gorm:"column:url"`
	Time     string    `gorm:"column:time"`
}

type CategoryBook struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name string    `gorm:"column:name"`
	Book []Book    `gorm:"many2many:book"`
}

type Book struct {
	gorm.Model
	ID             uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid()"`
	Title          string       `gorm:"column:title"`
	Language       string       `gorm:"column:languange"`
	Rank           int          `gorm:"column:rank"`
	ReadingTime    int          `gorm:"column:reading_time"`
	Likes          int          `gorm:"column:likes"`
	BookFile       string       `gorm:"column:book_file"`
	Overview       string       `gorm:"column:overview"`
	Writer         string       `gorm:"column:writer"`
	CategoryBookID uuid.UUID    `gorm:"column:category_id;type:uuid;default:gen_random_uuid()"`
	CategoryBook   CategoryBook `gorm:"foreignKey:CategoryBookID;references:ID"`
}
