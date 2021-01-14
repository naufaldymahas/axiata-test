package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID     int64     `json:"user_id" gorm:"primaryKey;unique"`
	UserName   string    `json:"user_name" gorm:"unique;not null"`
	Password   string    `json:"-" gorm:"not null"`
	EmployeeID string    `json:"employee_id" gorm:"unique;not null"`
	BirthPlace string    `json:"birth_place" gorm:"not null"`
	BirthDate  time.Time `json:"birh_date" gorm:"type:DATE;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"not null;autoUpdateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var lastYear int
	if err := tx.Raw("SELECT last_value FROM employee_year_seq").Scan(&lastYear).Error; err != nil {
		return err
	}
	nowYear := time.Now().Year()
	if lastYear < nowYear {
		if err := tx.Exec("SELECT setval('employee_year_seq', ?, true)", nowYear).Error; err != nil {
			return err
		}

		if err := tx.Exec("SELECT setval('employee_id_seq', (select cast((to_char(to_date(now()::text, 'YYYY'), 'YY') || '0001') as integer)), true)", nowYear).Error; err != nil {
			return err
		}
	}

	var eID string
	if err := tx.Raw("SELECT cast(nextval('employee_id_seq') AS varchar)").Scan(&eID).Error; err != nil {
		return err
	}

	u.EmployeeID = eID
	return nil
}
