package middleware

var publicEndpoint = []string{
	"/alfred.vcs_connection.v1.VCSConnectionService/ListAllSupportedVCSProviders",
	"/alfred.auth.v1.UserLoginService/UserLogin",
	"/alfred.auth.v1.UserLoginService/UserLoginCallback",
	"/alfred.user_profile.v1.UserProfileService/GetUserProfileBySub",
	//"/alfred.user_profile.v1.UserProfileService/CreateUserProfile",
}
