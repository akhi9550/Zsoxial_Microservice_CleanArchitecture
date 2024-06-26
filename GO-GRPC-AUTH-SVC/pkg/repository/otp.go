package repository

import (
	"errors"

	"github.com/akhi9550/auth-svc/pkg/domain"
	interfaces "github.com/akhi9550/auth-svc/pkg/repository/interface"
	"github.com/akhi9550/auth-svc/pkg/utils/models"
	"gorm.io/gorm"
)

type otpRepository struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpRepository{
		DB: DB,
	}
}

func (op *otpRepository) FindUserByPhoneNumber(phone string) (*domain.User, error) {
	var user domain.User
	res := op.DB.Where(&domain.User{Phone: phone}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error

	}
	return &user, nil
}
func (op *otpRepository) UserDetailsUsingPhone(phone string) (models.UserResponse, error) {
	var userDeatils models.UserResponse
	if err := op.DB.Raw("SELECT id,username,email,imageurl,isadmin FROM users WHERE phone = ?", phone).Scan(&userDeatils).Error; err != nil {
		return models.UserResponse{}, err
	}
	return userDeatils, nil
}
func (op *otpRepository) FindUsersByEmail(email string) (bool, error) {
	var count int
	if err := op.DB.Raw("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
func (op *otpRepository) GetUserPhoneByEmail(email string) (string, error) {
	var phone string
	if err := op.DB.Raw("SELECT phone FROM users WHERE email = ?", email).Scan(&phone).Error; err != nil {
		return "", nil
	}
	return phone, nil
}
