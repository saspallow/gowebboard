package api

// User Controller is the user controller interface
type UserController interface {
	SignIn(*UserSignInRequest) (*UserSignInResponse, error)
	SignOut(*UserSignOutRequest) (*UserSignOutResponse, error)
	ChangePassword(*UserChangePasswordRequest) (*UserChangePasswordResponse, error)
	Get(int) (*User, error)
}

// UserSignInRequest is the sign in request for user controller
type UserSignInRequest struct {
	Username string
	Password string
}

// UserSignInResponse is the sign in response for user controller
type UserSignInResponse struct {
	Token string
	UserID int
}

// UserSignOutRequest is the sign out reponse for user controller
type UserSignOutRequest struct {
	Token string
}
// UserSignOutRepones is the sign out reponse for user controller
type UserSignOutResponse struct {
}

// UserChangePasswordRequest type
type UserChangePasswordRequest struct {
	OldPassword string
	NewPassword string
}

// UserChangePasswordResponse type
type UserChangePasswordResponse struct {
}

type User struct {
	
}