package interfaces

import (
	"github.com/akhi9550/auth-svc/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserSignUpRequest) (*models.ReponseWithToken, error)
	UserLogin(user models.UserLoginRequest) (*models.ReponseWithToken, error)
	ForgotPassword(phone string) error
	ForgotPasswordVerifyAndChange(model models.ForgotVerify) error
	SpecificUserDetails(userID int) (models.UsersDetails, error)
	UserDetails(userID int) (models.UsersDetails, error)
	UpdateUserDetails(userDetails models.UsersProfileDetail, file []byte, userID int) (models.UsersProfileDetails, error)
	ChangePassword(id int, change models.ChangePassword) error
	CheckUserAvalilabilityWithUserID(userID int) (bool, error)
	UserData(userID int) (models.UserData, error)
	CheckUserAvalilabilityWithTagUserID(users []models.Tag) (bool, error)
	GetUserNameWithTagUserID(users []models.Tag) ([]models.UserTag, error)
	GetFollowingUsers(userID int)([]models.FollowUsers,error)
	ReportUser(userID int, req models.ReportRequest) error
	FollowREQ(userID, FollowingUserID int) error
	ShowFollowREQ(userID int) ([]models.FollowingRequests, error)
	AcceptFollowREQ(userID, FollowingUserID int) error
	UnFollow(userID, UnFollowUserID int) error
	Following(userID int) ([]models.FollowingResponse, error)
	Follower(userID int) ([]models.FollowingResponse, error)
	SearchUser(req models.SearchUser) ([]models.Users, error)
	VideoCallKey(userID, oppositeUser int) (string, error)
	CreateGroup(userID int, data models.GroupReq, users []models.Tag, file []byte)error
	ExitFormGroup(userID,groupID int)error
	ShowGroups(userID int)([]models.Groups,error)
	ShowGroupMembers(userID,groupID int)([]models.Mebmers,error)
}
