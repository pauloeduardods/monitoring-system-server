package auth

type LoginOutput struct {
	AccessToken  string `json:"accessToken,omitempty"`
	IdToken      string `json:"idToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	Session      string `json:"session,omitempty"`
}

type SignUpOutput struct {
	IsConfirmed bool `json:"isConfirmed"`
}

type ConfirmSignUpOutput struct {
}

type RefreshTokenOutput struct {
	AccessToken string `json:"accessToken"`
	IdToken     string `json:"idToken"`
}

type GetUserOutput struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type CreateAdminOutput struct {
	Username string `json:"username"`
}

type AddMFAOutput struct {
	SecretCode string `json:"secretCode"`
}
