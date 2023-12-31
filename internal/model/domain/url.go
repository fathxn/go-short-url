package domain

import "time"

type URL struct {
	Id        int       `db:"id" gorm:"primaryKey;autoIncrement;type:int"`
	UserId    string    `db:"user_id" gorm:"type:varchar(255)"`
	LongURL   string    `db:"long_url" gorm:"type:varchar(255)"`
	ShortCode string    `db:"short_code" gorm:"type:varchar(255)"`
	CreatedAt time.Time `db:"created_at"`
}
