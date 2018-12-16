package service

import (
	"github.com/cong5/persimmon/app/db"
	"github.com/cong5/persimmon/app/models"
	"github.com/revel/revel"
	"strings"
)

type UserService struct{}

func (this *UserService) GetUserByUid(uid int) (*models.Users, error) {
	users := &models.Users{Id: uid}
	_, err := db.MasterDB.Get(users)
	if err != nil {
		revel.AppLog.Errorf("GetUserByUid failed : %s", err)
		return nil, err
	}

	return users, nil
}

func (this *UserService) GetUserByEmail(email string) (*models.Users, error) {
	email = strings.ToLower(email)
	users := &models.Users{Email: email}
	_, err := db.MasterDB.Get(users)
	if err != nil {
		revel.AppLog.Errorf("GetUserByEmail failed : %s", err)
		return nil, err
	}

	return users, nil
}

func (this *UserService) UpdatePassword(id int, password string) (bool, error) {
	newPassword, hashErr := authService.HashPassword(password)
	if hashErr != nil {
		//revel.AppLog.Errorf("Update hash password failed: %s", hashErr)
		return false, hashErr
	}

	user := models.Users{Password: newPassword}
	_, err := db.MasterDB.Id(id).Cols("password").Update(user)
	if err != nil {
		//revel.AppLog.Errorf("Update password failed: %s", err)
		return false, hashErr
	}

	return true, nil
}

func (this *UserService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
