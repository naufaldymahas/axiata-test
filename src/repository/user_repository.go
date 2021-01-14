package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service-user-management/src/entity"
	"service-user-management/src/payload"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func ProvideUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (ur *UserRepository) InsertUser(user *entity.User) error {
	tx := ur.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var count int64
	tx.Model(&entity.User{}).Where("user_name = ?", user.UserName).Count(&count)
	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("User already exists")
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	requestBody, _ := json.Marshal(map[string]interface{}{
		"emailaddress": user.UserName,
		"password":     user.Password,
		"space":        user.UserID,
	})

	responseBody := bytes.NewBuffer(requestBody)

	resp, err := http.Post("http://localhost:8081/api/mail", "application/json", responseBody)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tx.Rollback()
		return err
	}

	var er payload.EmailRespose
	if err := json.Unmarshal(body, &er); err != nil {
		return err
	}

	if er.Status >= 400 {
		return fmt.Errorf(er.ErrorMsg)
	}

	return tx.Commit().Error
}

func (ur *UserRepository) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	if err := ur.db.Where("user_name = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) FindByUsernameOrEmployeeIDOrBirthPlace(search string) []entity.User {
	var users []entity.User
	ur.db.Where("lower(user_name) LIKE ?", search).Or("employee_id LIKE ?", search).Or("lower(birth_place) LIKE ?", search).Find(&users)
	return users
}
