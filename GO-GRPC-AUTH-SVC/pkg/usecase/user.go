package usecase

import (
	"errors"
	"fmt"

	"github.com/akhi9550/auth-svc/pkg/config"
	"github.com/akhi9550/auth-svc/pkg/helper"
	interfaces "github.com/akhi9550/auth-svc/pkg/repository/interface"
	services "github.com/akhi9550/auth-svc/pkg/usecase/interface"
	"github.com/akhi9550/auth-svc/pkg/utils/models"
	"github.com/google/uuid"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository interfaces.UserRepository
}

func NewUserUseCase(repository interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepository: repository,
	}
}
func (ur *userUseCase) UserSignUp(user models.UserSignUpRequest) (*models.ReponseWithToken, error) {
	username, err := ur.userRepository.CheckUserExistsByUsername(user.Username)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("error with server")
	}
	if username != nil {
		return &models.ReponseWithToken{}, errors.New("user with this username is already exists")
	}
	email, err := ur.userRepository.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("error with server")
	}
	if email != nil {
		return &models.ReponseWithToken{}, errors.New("user with this email is already exists")
	}

	phone, err := ur.userRepository.CheckUserExistsByPhone(user.Phone)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("error with server")
	}
	if phone != nil {
		return &models.ReponseWithToken{}, errors.New("user with this phone is already exists")
	}

	hashPassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("error in hashing password")
	}
	user.Password = hashPassword
	userData, err := ur.userRepository.UserSignUp(user)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("could not add the user")
	}
	accessToken, err := helper.GenerateAccessTokenUser(userData)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("couldn't create access token due to error")
	}
	RefreshToken, err := helper.GenerateRefreshTokenUser(userData)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("couldn't create access token due to error")
	}
	return &models.ReponseWithToken{
		Users:        userData,
		AccessToken:  accessToken,
		RefreshToken: RefreshToken,
	}, nil
}

func (ur *userUseCase) UserLogin(user models.UserLoginRequest) (*models.ReponseWithToken, error) {
	email, err := ur.userRepository.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("error with server")
	}
	if email == nil {
		return &models.ReponseWithToken{}, errors.New("email doesn't exist")
	}
	userdeatils, err := ur.userRepository.FindUserByEmail(user)
	if err != nil {
		return &models.ReponseWithToken{}, err
	}
	ok, err := ur.userRepository.FindUserBlockorNot(user.Email)
	if ok {
		return &models.ReponseWithToken{}, errors.New("user is blocked")
	}
	if err != nil {
		return &models.ReponseWithToken{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userdeatils.Password), []byte(user.Password))
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("password not matching")
	}
	var userData models.UserResponse
	err = copier.Copy(&userData, &userdeatils)
	if err != nil {
		return &models.ReponseWithToken{}, err
	}
	accessToken, err := helper.GenerateAccessTokenUser(userData)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("couldn't create access token due to error")
	}
	RefreshToken, err := helper.GenerateRefreshTokenUser(userData)
	if err != nil {
		return &models.ReponseWithToken{}, errors.New("couldn't create access token due to error")
	}
	return &models.ReponseWithToken{
		Users:        userData,
		AccessToken:  accessToken,
		RefreshToken: RefreshToken,
	}, nil
}

func (ur *userUseCase) ForgotPassword(phone string) error {
	cfg, _ := config.LoadConfig()
	ok := ur.userRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	helper.TwilioSetup(cfg.ACCOUNTSID, cfg.AUTHTOKEN)
	_, err := helper.TwilioSendOTP(phone, cfg.SERVICESSID)
	if err != nil {
		return errors.New("error ocurred while generating OTP")
	}
	return nil
}

func (ur *userUseCase) ForgotPasswordVerifyAndChange(model models.ForgotVerify) error {
	cfg, _ := config.LoadConfig()
	phone, err := ur.userRepository.FindIdFromPhone(model.Phone)
	if err != nil {
		return errors.New("cannot find user from mobile number")
	}
	helper.TwilioSetup(cfg.ACCOUNTSID, cfg.AUTHTOKEN)
	err = helper.TwilioVerifyOTP(cfg.SERVICESSID, model.Otp, model.Phone)
	if err != nil {
		return errors.New("error while verifying")
	}
	newpassword, err := helper.PasswordHashing(model.NewPassword)
	if err != nil {
		return errors.New("error in hashing password")
	}

	if err := ur.userRepository.ChangePassword(phone, newpassword); err != nil {
		return errors.New("could not change password")
	}

	return nil
}

func (ur *userUseCase) SpecificUserDetails(userID int) (models.UsersDetails, error) {
	return ur.userRepository.SpecificUserDetails(userID)
}

func (ur *userUseCase) UserDetails(userID int) (models.UsersDetails, error) {
	return ur.userRepository.SpecificUserDetails(userID)
}

func (ur *userUseCase) UpdateUserDetails(userDetails models.UsersProfileDetail, file []byte, userID int) (models.UsersProfileDetails, error) {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return models.UsersProfileDetails{}, errors.New("user doesn't exist")
	}
	if userDetails.Firstname != "" {
		ur.userRepository.UpdateFirstName(userDetails.Firstname, userID)
	}
	if userDetails.Lastname != "" {
		ur.userRepository.UpdateLastName(userDetails.Lastname, userID)
	}
	if userDetails.Username != "" {
		ok := ur.userRepository.ExistUsername(userDetails.Username)
		if ok {
			return models.UsersProfileDetails{}, errors.New("username already exist")
		}
		ur.userRepository.UpdateUserName(userDetails.Username, userID)
	}
	if userDetails.Dob != "" {
		ur.userRepository.UpdateDOB(userDetails.Dob, userID)
	}
	if userDetails.Gender != "" {
		ur.userRepository.UpdateGender(userDetails.Gender, userID)
	}
	if userDetails.Phone != "" {
		ok := ur.userRepository.ExistPhone(userDetails.Phone)
		if ok {
			return models.UsersProfileDetails{}, errors.New("phone already exist")
		}
		ur.userRepository.UpdateUserPhone(userDetails.Phone, userID)
	}
	if userDetails.Email != "" {
		ok := ur.userRepository.ExistEmail(userDetails.Email)
		if ok {
			return models.UsersProfileDetails{}, errors.New("email already exist")
		}
		ur.userRepository.UpdateUserEmail(userDetails.Email, userID)
	}
	if userDetails.Bio != "" {
		ur.userRepository.UpdateBIO(userDetails.Bio, userID)
	}
	fileUID := uuid.New()
	fileName := fileUID.String()
	s3Path := userDetails.Username + fileName
	if file != nil {
		url, err := helper.AddImageToAwsS3(file, s3Path)
		if err != nil {
			return models.UsersProfileDetails{}, errors.New("passing aws")
		}
		if string(file) != "" {
			ur.userRepository.UpdatePhoto(url, userID)
		}
	}
	return ur.userRepository.UserDetails(userID)
}

func (ur *userUseCase) ChangePassword(id int, change models.ChangePassword) error {
	userPassword, err := ur.userRepository.GetPassword(id)
	if err != nil {
		return errors.New("internal error")
	}
	err = helper.CompareHashAndPassword(userPassword, change.Oldpassword)
	if err != nil {
		return errors.New("password incorrect")
	}
	if change.Password != change.Repassword {
		return errors.New("password doesn't match")
	}
	newpassword, err := helper.PasswordHash(change.Password)
	if err != nil {
		return errors.New("error in hashing password")
	}
	return ur.userRepository.Changepassword(id, string(newpassword))
}

func (ur *userUseCase) CheckUserAvalilabilityWithUserID(userID int) (bool, error) {
	ok, _ := ur.userRepository.CheckUserAvalilabilityWithUserID(userID)
	return ok, nil
}

func (ur *userUseCase) UserData(userID int) (models.UserData, error) {
	data, err := ur.userRepository.UserData(userID)
	if err != nil {
		return models.UserData{}, err
	}
	return data, nil
}

func (ur *userUseCase) CheckUserAvalilabilityWithTagUserID(users []models.Tag) (bool, error) {
	ok, _ := ur.userRepository.CheckUserAvalilabilityWithTagUserID(users)
	return ok, nil
}

func (ur *userUseCase) GetUserNameWithTagUserID(users []models.Tag) ([]models.UserTag, error) {
	data, err := ur.userRepository.GetUserNameWithTagUserID(users)
	if err != nil {
		return []models.UserTag{}, err
	}
	return data, nil
}

func (ur *userUseCase) GetFollowingUsers(userID int) ([]models.FollowUsers, error) {
	data, err := ur.userRepository.GetFollowingUsers(userID)
	if err != nil {
		return []models.FollowUsers{}, err
	}
	return data, nil
}

func (ur *userUseCase) ReportUser(userID int, req models.ReportRequest) error {
	ReportuserExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !ReportuserExist {
		return errors.New("user doesn't exist")
	}
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(int(req.UserID))
	if !userExist {
		return errors.New("user doesn't exist")
	}
	Isreport := ur.userRepository.AlreadyReported(userID, int(req.UserID))
	if Isreport {
		return errors.New("already reported")
	}
	err := ur.userRepository.ReportUser(userID, req)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userUseCase) FollowREQ(userID, FollowingUserID int) error {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return errors.New("user doesn't exist")
	}
	FollowuserExist := ur.userRepository.CheckUserAvailabilityWithUserID(FollowingUserID)
	if !FollowuserExist {
		return errors.New("user doesn't exist")
	}
	exist := ur.userRepository.ExistFollowreq(userID, FollowingUserID)
	if exist {
		return errors.New("already send following request")
	}
	err := ur.userRepository.FollowREQ(userID, FollowingUserID)
	if err != nil {
		return err
	}
	UserName, err := ur.UserData(userID)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("%s Requested to Follow You", UserName.Username)
	helper.SendNotification(models.Notification{
		UserID:   FollowingUserID,
		SenderID: userID,
	}, []byte(msg))

	return nil
}

func (ur *userUseCase) ShowFollowREQ(userID int) ([]models.FollowingRequests, error) {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return nil, errors.New("user doesn't exist")
	}
	data, err := ur.userRepository.ShowFollowREQ(userID)
	if err != nil {
		return nil, err
	}
	var response []models.FollowingRequests
	for _, follower := range data {
		userData, err := ur.userRepository.UserData(int(follower.FollowingUserID))
		if err != nil {
			return nil, err
		}
		details := models.FollowingRequests{
			FollowingUserID: follower.FollowingUserID,
			FollowingUser:   userData.Username,
			Profile:         userData.Profile,
			CreatedAt:       follower.CreatedAt,
		}
		response = append(response, details)
	}
	return response, nil
}

func (ur *userUseCase) AcceptFollowREQ(userID, FollowingUserID int) error {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return errors.New("user doesn't exist")
	}
	FollowuserExist := ur.userRepository.CheckUserAvailabilityWithUserID(FollowingUserID)
	if !FollowuserExist {
		return errors.New("user doesn't exist")
	}
	req := ur.userRepository.CheckRequest(userID, FollowingUserID)
	if !req {
		return errors.New("no req available")
	}
	alreadyfollow := ur.userRepository.AlreadyAccepted(userID, FollowingUserID)
	if alreadyfollow {
		return errors.New("already exist")
	}
	err := ur.userRepository.AcceptFollowREQ(userID, FollowingUserID)
	if err != nil {
		return err
	}

	UserName, err := ur.UserData(userID)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("%s Started Follow You", UserName.Username)
	helper.SendNotification(models.Notification{
		UserID:   FollowingUserID,
		SenderID: userID,
	}, []byte(msg))

	return nil
}

func (ur *userUseCase) UnFollow(userID, UnFollowUserID int) error {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return errors.New("user doesn't exist")
	}
	FollowuserExist := ur.userRepository.CheckUserAvailabilityWithUserID(UnFollowUserID)
	if !FollowuserExist {
		return errors.New("user doesn't exist")
	}
	err := ur.userRepository.UnFollow(userID, UnFollowUserID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userUseCase) Following(userID int) ([]models.FollowingResponse, error) {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return []models.FollowingResponse{}, errors.New("user doesn't exist")
	}
	data, err := ur.userRepository.Following(userID)
	if err != nil {
		return []models.FollowingResponse{}, err
	}
	var response []models.FollowingResponse
	for _, follow := range data {
		userData, err := ur.userRepository.UserData(int(follow.FollowingUserID))
		if err != nil {
			return nil, err
		}
		details := models.FollowingResponse{
			FollowingUser: userData.Username,
			Profile:       userData.Profile,
		}
		response = append(response, details)
	}
	return response, nil
}

func (ur *userUseCase) Follower(userID int) ([]models.FollowingResponse, error) {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return []models.FollowingResponse{}, errors.New("user doesn't exist")
	}
	data, err := ur.userRepository.Follower(userID)
	if err != nil {
		return []models.FollowingResponse{}, err
	}
	var response []models.FollowingResponse
	for _, follow := range data {
		userData, err := ur.userRepository.UserData(int(follow.FollowingUserID))
		if err != nil {
			return nil, err
		}
		details := models.FollowingResponse{
			FollowingUser: userData.Username,
			Profile:       userData.Profile,
		}
		response = append(response, details)
	}
	return response, nil
}

func (ur *userUseCase) SearchUser(req models.SearchUser) ([]models.Users, error) {
	data, err := ur.userRepository.SearchUser(req)
	if err != nil {
		return []models.Users{}, err
	}
	return data, nil
}

func (ur *userUseCase) VideoCallKey(userID, oppositeUser int) (string, error) {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return "", errors.New("user doesn't exist")
	}
	FollowuserExist := ur.userRepository.CheckUserAlreadyExistFromFollowers(userID, oppositeUser)
	if !FollowuserExist {
		return "", errors.New("this user not follow you")
	}
	key, err := helper.GenerateVideoCallKey(userID, oppositeUser)
	if err != nil {
		return "", err
	}
	UserName, err := ur.userRepository.UserData(userID)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s Invaited to Join the link :- https://zsoxial.zhooze.shop/index?room=%s", UserName.Username, key)
	helper.SendNotification(models.Notification{
		UserID:   oppositeUser,
		SenderID: userID,
	}, []byte(url))

	return key, nil
}

func (ur *userUseCase) CreateGroup(userID int, data models.GroupReq, users []models.Tag, file []byte) error {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(int(userID))
	if !userExist {
		return errors.New("user doesn't exist")
	}
	fileUID := uuid.New()
	fileName := fileUID.String()
	s3Path := data.Name + fileName
	url, err := helper.AddImageToAwsS3(file, s3Path)
	if err != nil {
		return errors.New("passing aws")
	}

	usersExist, err := ur.userRepository.CheckUserAvalilabilityWithTagUserID(users)
	if !usersExist {
		return errors.New("user doesn't exist")
	}
	if err != nil {
		return errors.New("user doesn't exist")
	}
	err = ur.userRepository.CreateGroup(userID, data.Name, data.Description, users, url)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userUseCase) ExitFormGroup(userID, groupID int) error {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return errors.New("user doesn't exist")
	}
	groupExist := ur.userRepository.CheckGroupAvailabilityWithID(groupID)
	if !groupExist {
		return errors.New("group doesn't exist")
	}
	groupandUser := ur.userRepository.CheckGroupAvailability(userID, groupID)
	if !groupandUser {
		return errors.New("user doesn't exist from the group")
	}
	ownwerormember := ur.userRepository.CheckUserOwnerOrMember(userID, groupID)
	if ownwerormember == 1 {
		err := ur.userRepository.ExitFormGroup(userID)
		if err != nil {
			return errors.New("error From Exit from Group")
		}
	}
	if ownwerormember > 1 {
		err := ur.userRepository.DeleteFormGroup(userID, groupID)
		if err != nil {
			return errors.New("error From delete Group")
		}
	}
	return nil
}

func (ur *userUseCase) ShowGroups(userID int) ([]models.Groups, error) {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return []models.Groups{}, errors.New("user doesn't exist")
	}
	data, err := ur.userRepository.ShowGroups(userID)
	if err != nil {
		return []models.Groups{}, err
	}
	var result []models.Groups
	for _, v := range data {
		groups := models.Groups{
			ID:      v.ID,
			Name:    v.Name,
			Profile: v.Profile,
		}
		result = append(result, groups)
	}
	return result, nil
}

func (ur *userUseCase) ShowGroupMembers(userID, groupID int) ([]models.Mebmers, error) {
	userExist := ur.userRepository.CheckUserAvailabilityWithUserID(userID)
	if !userExist {
		return []models.Mebmers{}, errors.New("user doesn't exist")
	}
	groupExist := ur.userRepository.CheckGroupAvailabilityWithID(groupID)
	if !groupExist {
		return []models.Mebmers{}, errors.New("user doesn't exist")
	}
	groupandUser := ur.userRepository.CheckGroupAvailability(userID, groupID)
	if !groupandUser {
		return []models.Mebmers{}, errors.New("user doesn't exist from the group")
	}
	data, err := ur.userRepository.ShowGroupMembers(groupID)
	if err != nil {
		return []models.Mebmers{}, err
	}
	var response []models.Mebmers
	for _, v := range data {
		userData, err := ur.userRepository.UserData(int(v.ID))
		if err != nil {
			return nil, err
		}
		details := models.Mebmers{
			Username: userData.Username,
			Profile:  userData.Profile,
		}
		response = append(response, details)
	}
	return response, nil
}
