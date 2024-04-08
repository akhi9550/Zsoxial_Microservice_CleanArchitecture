// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: pkg/pb/auth/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthService_UserSignUp_FullMethodName                          = "/user.AuthService/UserSignUp"
	AuthService_UserLogin_FullMethodName                           = "/user.AuthService/UserLogin"
	AuthService_SendOtp_FullMethodName                             = "/user.AuthService/SendOtp"
	AuthService_VerifyOtp_FullMethodName                           = "/user.AuthService/VerifyOtp"
	AuthService_ForgotPassword_FullMethodName                      = "/user.AuthService/ForgotPassword"
	AuthService_ForgotPasswordVerifyAndChange_FullMethodName       = "/user.AuthService/ForgotPasswordVerifyAndChange"
	AuthService_UserDetails_FullMethodName                         = "/user.AuthService/UserDetails"
	AuthService_UpdateUserDetails_FullMethodName                   = "/user.AuthService/UpdateUserDetails"
	AuthService_ChangePassword_FullMethodName                      = "/user.AuthService/ChangePassword"
	AuthService_AdminLogin_FullMethodName                          = "/user.AuthService/AdminLogin"
	AuthService_ShowAllUsers_FullMethodName                        = "/user.AuthService/ShowAllUsers"
	AuthService_AdminBlockUser_FullMethodName                      = "/user.AuthService/AdminBlockUser"
	AuthService_AdminUnblockUser_FullMethodName                    = "/user.AuthService/AdminUnblockUser"
	AuthService_CheckUserAvalilabilityWithUserID_FullMethodName    = "/user.AuthService/CheckUserAvalilabilityWithUserID"
	AuthService_UserData_FullMethodName                            = "/user.AuthService/UserData"
	AuthService_CheckUserAvalilabilityWithTagUserID_FullMethodName = "/user.AuthService/CheckUserAvalilabilityWithTagUserID"
	AuthService_GetUserNameWithTagUserID_FullMethodName            = "/user.AuthService/GetUserNameWithTagUserID"
	AuthService_ReportUser_FullMethodName                          = "/user.AuthService/ReportUser"
	AuthService_FollowREQ_FullMethodName                           = "/user.AuthService/FollowREQ"
	AuthService_ShowFollowREQ_FullMethodName                       = "/user.AuthService/ShowFollowREQ"
	AuthService_AcceptFollowREQ_FullMethodName                     = "/user.AuthService/AcceptFollowREQ"
	AuthService_Following_FullMethodName                           = "/user.AuthService/Following"
	AuthService_Follower_FullMethodName                            = "/user.AuthService/Follower"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	UserSignUp(ctx context.Context, in *UserSignUpRequest, opts ...grpc.CallOption) (*UserSignUpResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	SendOtp(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*SendOtpResponse, error)
	VerifyOtp(ctx context.Context, in *VerifyOtpRequest, opts ...grpc.CallOption) (*VerifyOtpResponse, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error)
	ForgotPasswordVerifyAndChange(ctx context.Context, in *ForgotPasswordVerifyAndChangeRequest, opts ...grpc.CallOption) (*ForgotPasswordVerifyAndChangeResponse, error)
	UserDetails(ctx context.Context, in *UserDetailsRequest, opts ...grpc.CallOption) (*UserDetailsResponse, error)
	UpdateUserDetails(ctx context.Context, in *UpdateUserDetailsRequest, opts ...grpc.CallOption) (*UpdateUserDetailsResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	ShowAllUsers(ctx context.Context, in *ShowAllUsersRequest, opts ...grpc.CallOption) (*ShowAllUsersResponse, error)
	AdminBlockUser(ctx context.Context, in *AdminBlockUserRequest, opts ...grpc.CallOption) (*AdminBlockUserResponse, error)
	AdminUnblockUser(ctx context.Context, in *AdminUnblockUserRequest, opts ...grpc.CallOption) (*AdminUnblockUserResponse, error)
	CheckUserAvalilabilityWithUserID(ctx context.Context, in *CheckUserAvalilabilityWithUserIDRequest, opts ...grpc.CallOption) (*CheckUserAvalilabilityWithUserIDResponse, error)
	UserData(ctx context.Context, in *UserDataRequest, opts ...grpc.CallOption) (*UserDataResponse, error)
	CheckUserAvalilabilityWithTagUserID(ctx context.Context, in *CheckUserAvalilabilityWithTagUserIDRequest, opts ...grpc.CallOption) (*CheckUserAvalilabilityWithTagUserIDResponse, error)
	GetUserNameWithTagUserID(ctx context.Context, in *GetUserNameWithTagUserIDRequest, opts ...grpc.CallOption) (*GetUserNameWithTagUserIDResponse, error)
	ReportUser(ctx context.Context, in *ReportUserRequest, opts ...grpc.CallOption) (*ReportUserResponse, error)
	FollowREQ(ctx context.Context, in *FollowREQRequest, opts ...grpc.CallOption) (*FollowREQResponse, error)
	ShowFollowREQ(ctx context.Context, in *ShowREQRequest, opts ...grpc.CallOption) (*ShowREQResponse, error)
	AcceptFollowREQ(ctx context.Context, in *AcceptFollowREQRequest, opts ...grpc.CallOption) (*AcceptFollowREQResponse, error)
	Following(ctx context.Context, in *FollowingRequest, opts ...grpc.CallOption) (*FollowingResponse, error)
	Follower(ctx context.Context, in *FollowerRequest, opts ...grpc.CallOption) (*FollowerResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) UserSignUp(ctx context.Context, in *UserSignUpRequest, opts ...grpc.CallOption) (*UserSignUpResponse, error) {
	out := new(UserSignUpResponse)
	err := c.cc.Invoke(ctx, AuthService_UserSignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, AuthService_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SendOtp(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*SendOtpResponse, error) {
	out := new(SendOtpResponse)
	err := c.cc.Invoke(ctx, AuthService_SendOtp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyOtp(ctx context.Context, in *VerifyOtpRequest, opts ...grpc.CallOption) (*VerifyOtpResponse, error) {
	out := new(VerifyOtpResponse)
	err := c.cc.Invoke(ctx, AuthService_VerifyOtp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	out := new(ForgotPasswordResponse)
	err := c.cc.Invoke(ctx, AuthService_ForgotPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgotPasswordVerifyAndChange(ctx context.Context, in *ForgotPasswordVerifyAndChangeRequest, opts ...grpc.CallOption) (*ForgotPasswordVerifyAndChangeResponse, error) {
	out := new(ForgotPasswordVerifyAndChangeResponse)
	err := c.cc.Invoke(ctx, AuthService_ForgotPasswordVerifyAndChange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserDetails(ctx context.Context, in *UserDetailsRequest, opts ...grpc.CallOption) (*UserDetailsResponse, error) {
	out := new(UserDetailsResponse)
	err := c.cc.Invoke(ctx, AuthService_UserDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUserDetails(ctx context.Context, in *UpdateUserDetailsRequest, opts ...grpc.CallOption) (*UpdateUserDetailsResponse, error) {
	out := new(UpdateUserDetailsResponse)
	err := c.cc.Invoke(ctx, AuthService_UpdateUserDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, AuthService_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, AuthService_AdminLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ShowAllUsers(ctx context.Context, in *ShowAllUsersRequest, opts ...grpc.CallOption) (*ShowAllUsersResponse, error) {
	out := new(ShowAllUsersResponse)
	err := c.cc.Invoke(ctx, AuthService_ShowAllUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AdminBlockUser(ctx context.Context, in *AdminBlockUserRequest, opts ...grpc.CallOption) (*AdminBlockUserResponse, error) {
	out := new(AdminBlockUserResponse)
	err := c.cc.Invoke(ctx, AuthService_AdminBlockUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AdminUnblockUser(ctx context.Context, in *AdminUnblockUserRequest, opts ...grpc.CallOption) (*AdminUnblockUserResponse, error) {
	out := new(AdminUnblockUserResponse)
	err := c.cc.Invoke(ctx, AuthService_AdminUnblockUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckUserAvalilabilityWithUserID(ctx context.Context, in *CheckUserAvalilabilityWithUserIDRequest, opts ...grpc.CallOption) (*CheckUserAvalilabilityWithUserIDResponse, error) {
	out := new(CheckUserAvalilabilityWithUserIDResponse)
	err := c.cc.Invoke(ctx, AuthService_CheckUserAvalilabilityWithUserID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserData(ctx context.Context, in *UserDataRequest, opts ...grpc.CallOption) (*UserDataResponse, error) {
	out := new(UserDataResponse)
	err := c.cc.Invoke(ctx, AuthService_UserData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckUserAvalilabilityWithTagUserID(ctx context.Context, in *CheckUserAvalilabilityWithTagUserIDRequest, opts ...grpc.CallOption) (*CheckUserAvalilabilityWithTagUserIDResponse, error) {
	out := new(CheckUserAvalilabilityWithTagUserIDResponse)
	err := c.cc.Invoke(ctx, AuthService_CheckUserAvalilabilityWithTagUserID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserNameWithTagUserID(ctx context.Context, in *GetUserNameWithTagUserIDRequest, opts ...grpc.CallOption) (*GetUserNameWithTagUserIDResponse, error) {
	out := new(GetUserNameWithTagUserIDResponse)
	err := c.cc.Invoke(ctx, AuthService_GetUserNameWithTagUserID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ReportUser(ctx context.Context, in *ReportUserRequest, opts ...grpc.CallOption) (*ReportUserResponse, error) {
	out := new(ReportUserResponse)
	err := c.cc.Invoke(ctx, AuthService_ReportUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) FollowREQ(ctx context.Context, in *FollowREQRequest, opts ...grpc.CallOption) (*FollowREQResponse, error) {
	out := new(FollowREQResponse)
	err := c.cc.Invoke(ctx, AuthService_FollowREQ_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ShowFollowREQ(ctx context.Context, in *ShowREQRequest, opts ...grpc.CallOption) (*ShowREQResponse, error) {
	out := new(ShowREQResponse)
	err := c.cc.Invoke(ctx, AuthService_ShowFollowREQ_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AcceptFollowREQ(ctx context.Context, in *AcceptFollowREQRequest, opts ...grpc.CallOption) (*AcceptFollowREQResponse, error) {
	out := new(AcceptFollowREQResponse)
	err := c.cc.Invoke(ctx, AuthService_AcceptFollowREQ_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Following(ctx context.Context, in *FollowingRequest, opts ...grpc.CallOption) (*FollowingResponse, error) {
	out := new(FollowingResponse)
	err := c.cc.Invoke(ctx, AuthService_Following_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Follower(ctx context.Context, in *FollowerRequest, opts ...grpc.CallOption) (*FollowerResponse, error) {
	out := new(FollowerResponse)
	err := c.cc.Invoke(ctx, AuthService_Follower_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	UserSignUp(context.Context, *UserSignUpRequest) (*UserSignUpResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	SendOtp(context.Context, *SendOtpRequest) (*SendOtpResponse, error)
	VerifyOtp(context.Context, *VerifyOtpRequest) (*VerifyOtpResponse, error)
	ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error)
	ForgotPasswordVerifyAndChange(context.Context, *ForgotPasswordVerifyAndChangeRequest) (*ForgotPasswordVerifyAndChangeResponse, error)
	UserDetails(context.Context, *UserDetailsRequest) (*UserDetailsResponse, error)
	UpdateUserDetails(context.Context, *UpdateUserDetailsRequest) (*UpdateUserDetailsResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error)
	ShowAllUsers(context.Context, *ShowAllUsersRequest) (*ShowAllUsersResponse, error)
	AdminBlockUser(context.Context, *AdminBlockUserRequest) (*AdminBlockUserResponse, error)
	AdminUnblockUser(context.Context, *AdminUnblockUserRequest) (*AdminUnblockUserResponse, error)
	CheckUserAvalilabilityWithUserID(context.Context, *CheckUserAvalilabilityWithUserIDRequest) (*CheckUserAvalilabilityWithUserIDResponse, error)
	UserData(context.Context, *UserDataRequest) (*UserDataResponse, error)
	CheckUserAvalilabilityWithTagUserID(context.Context, *CheckUserAvalilabilityWithTagUserIDRequest) (*CheckUserAvalilabilityWithTagUserIDResponse, error)
	GetUserNameWithTagUserID(context.Context, *GetUserNameWithTagUserIDRequest) (*GetUserNameWithTagUserIDResponse, error)
	ReportUser(context.Context, *ReportUserRequest) (*ReportUserResponse, error)
	FollowREQ(context.Context, *FollowREQRequest) (*FollowREQResponse, error)
	ShowFollowREQ(context.Context, *ShowREQRequest) (*ShowREQResponse, error)
	AcceptFollowREQ(context.Context, *AcceptFollowREQRequest) (*AcceptFollowREQResponse, error)
	Following(context.Context, *FollowingRequest) (*FollowingResponse, error)
	Follower(context.Context, *FollowerRequest) (*FollowerResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) UserSignUp(context.Context, *UserSignUpRequest) (*UserSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignUp not implemented")
}
func (UnimplementedAuthServiceServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedAuthServiceServer) SendOtp(context.Context, *SendOtpRequest) (*SendOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOtp not implemented")
}
func (UnimplementedAuthServiceServer) VerifyOtp(context.Context, *VerifyOtpRequest) (*VerifyOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOtp not implemented")
}
func (UnimplementedAuthServiceServer) ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedAuthServiceServer) ForgotPasswordVerifyAndChange(context.Context, *ForgotPasswordVerifyAndChangeRequest) (*ForgotPasswordVerifyAndChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPasswordVerifyAndChange not implemented")
}
func (UnimplementedAuthServiceServer) UserDetails(context.Context, *UserDetailsRequest) (*UserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetails not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUserDetails(context.Context, *UpdateUserDetailsRequest) (*UpdateUserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserDetails not implemented")
}
func (UnimplementedAuthServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthServiceServer) AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAuthServiceServer) ShowAllUsers(context.Context, *ShowAllUsersRequest) (*ShowAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAllUsers not implemented")
}
func (UnimplementedAuthServiceServer) AdminBlockUser(context.Context, *AdminBlockUserRequest) (*AdminBlockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminBlockUser not implemented")
}
func (UnimplementedAuthServiceServer) AdminUnblockUser(context.Context, *AdminUnblockUserRequest) (*AdminUnblockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUnblockUser not implemented")
}
func (UnimplementedAuthServiceServer) CheckUserAvalilabilityWithUserID(context.Context, *CheckUserAvalilabilityWithUserIDRequest) (*CheckUserAvalilabilityWithUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAvalilabilityWithUserID not implemented")
}
func (UnimplementedAuthServiceServer) UserData(context.Context, *UserDataRequest) (*UserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserData not implemented")
}
func (UnimplementedAuthServiceServer) CheckUserAvalilabilityWithTagUserID(context.Context, *CheckUserAvalilabilityWithTagUserIDRequest) (*CheckUserAvalilabilityWithTagUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAvalilabilityWithTagUserID not implemented")
}
func (UnimplementedAuthServiceServer) GetUserNameWithTagUserID(context.Context, *GetUserNameWithTagUserIDRequest) (*GetUserNameWithTagUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserNameWithTagUserID not implemented")
}
func (UnimplementedAuthServiceServer) ReportUser(context.Context, *ReportUserRequest) (*ReportUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportUser not implemented")
}
func (UnimplementedAuthServiceServer) FollowREQ(context.Context, *FollowREQRequest) (*FollowREQResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowREQ not implemented")
}
func (UnimplementedAuthServiceServer) ShowFollowREQ(context.Context, *ShowREQRequest) (*ShowREQResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowFollowREQ not implemented")
}
func (UnimplementedAuthServiceServer) AcceptFollowREQ(context.Context, *AcceptFollowREQRequest) (*AcceptFollowREQResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFollowREQ not implemented")
}
func (UnimplementedAuthServiceServer) Following(context.Context, *FollowingRequest) (*FollowingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Following not implemented")
}
func (UnimplementedAuthServiceServer) Follower(context.Context, *FollowerRequest) (*FollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follower not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_UserSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserSignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserSignUp(ctx, req.(*UserSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SendOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SendOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SendOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SendOtp(ctx, req.(*SendOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyOtp(ctx, req.(*VerifyOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ForgotPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgotPasswordVerifyAndChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordVerifyAndChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgotPasswordVerifyAndChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ForgotPasswordVerifyAndChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgotPasswordVerifyAndChange(ctx, req.(*ForgotPasswordVerifyAndChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserDetails(ctx, req.(*UserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdateUserDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUserDetails(ctx, req.(*UpdateUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AdminLogin(ctx, req.(*AdminLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ShowAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ShowAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ShowAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ShowAllUsers(ctx, req.(*ShowAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AdminBlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminBlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AdminBlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AdminBlockUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AdminBlockUser(ctx, req.(*AdminBlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AdminUnblockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUnblockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AdminUnblockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AdminUnblockUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AdminUnblockUser(ctx, req.(*AdminUnblockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckUserAvalilabilityWithUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAvalilabilityWithUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckUserAvalilabilityWithUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CheckUserAvalilabilityWithUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckUserAvalilabilityWithUserID(ctx, req.(*CheckUserAvalilabilityWithUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserData(ctx, req.(*UserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckUserAvalilabilityWithTagUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAvalilabilityWithTagUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckUserAvalilabilityWithTagUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CheckUserAvalilabilityWithTagUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckUserAvalilabilityWithTagUserID(ctx, req.(*CheckUserAvalilabilityWithTagUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserNameWithTagUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserNameWithTagUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserNameWithTagUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetUserNameWithTagUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserNameWithTagUserID(ctx, req.(*GetUserNameWithTagUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ReportUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ReportUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ReportUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ReportUser(ctx, req.(*ReportUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_FollowREQ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowREQRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).FollowREQ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_FollowREQ_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).FollowREQ(ctx, req.(*FollowREQRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ShowFollowREQ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowREQRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ShowFollowREQ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ShowFollowREQ_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ShowFollowREQ(ctx, req.(*ShowREQRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AcceptFollowREQ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptFollowREQRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AcceptFollowREQ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AcceptFollowREQ_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AcceptFollowREQ(ctx, req.(*AcceptFollowREQRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Following_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Following(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Following_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Following(ctx, req.(*FollowingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Follower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Follower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Follower_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Follower(ctx, req.(*FollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignUp",
			Handler:    _AuthService_UserSignUp_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _AuthService_UserLogin_Handler,
		},
		{
			MethodName: "SendOtp",
			Handler:    _AuthService_SendOtp_Handler,
		},
		{
			MethodName: "VerifyOtp",
			Handler:    _AuthService_VerifyOtp_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _AuthService_ForgotPassword_Handler,
		},
		{
			MethodName: "ForgotPasswordVerifyAndChange",
			Handler:    _AuthService_ForgotPasswordVerifyAndChange_Handler,
		},
		{
			MethodName: "UserDetails",
			Handler:    _AuthService_UserDetails_Handler,
		},
		{
			MethodName: "UpdateUserDetails",
			Handler:    _AuthService_UpdateUserDetails_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AuthService_ChangePassword_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _AuthService_AdminLogin_Handler,
		},
		{
			MethodName: "ShowAllUsers",
			Handler:    _AuthService_ShowAllUsers_Handler,
		},
		{
			MethodName: "AdminBlockUser",
			Handler:    _AuthService_AdminBlockUser_Handler,
		},
		{
			MethodName: "AdminUnblockUser",
			Handler:    _AuthService_AdminUnblockUser_Handler,
		},
		{
			MethodName: "CheckUserAvalilabilityWithUserID",
			Handler:    _AuthService_CheckUserAvalilabilityWithUserID_Handler,
		},
		{
			MethodName: "UserData",
			Handler:    _AuthService_UserData_Handler,
		},
		{
			MethodName: "CheckUserAvalilabilityWithTagUserID",
			Handler:    _AuthService_CheckUserAvalilabilityWithTagUserID_Handler,
		},
		{
			MethodName: "GetUserNameWithTagUserID",
			Handler:    _AuthService_GetUserNameWithTagUserID_Handler,
		},
		{
			MethodName: "ReportUser",
			Handler:    _AuthService_ReportUser_Handler,
		},
		{
			MethodName: "FollowREQ",
			Handler:    _AuthService_FollowREQ_Handler,
		},
		{
			MethodName: "ShowFollowREQ",
			Handler:    _AuthService_ShowFollowREQ_Handler,
		},
		{
			MethodName: "AcceptFollowREQ",
			Handler:    _AuthService_AcceptFollowREQ_Handler,
		},
		{
			MethodName: "Following",
			Handler:    _AuthService_Following_Handler,
		},
		{
			MethodName: "Follower",
			Handler:    _AuthService_Follower_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth/auth.proto",
}
