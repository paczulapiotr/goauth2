package router

import "time"

// StatusResp response data type
type StatusResp struct {
	Status string `json:"status"`
}

// AuthReq auth request data type
type AuthReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// RegisterReq register request data type
type RegisterReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// AuthResp auth code response data type
type AuthResp struct {
	Code string `json:"code"`
}

// UseCodeReq use code request data type
type UseCodeReq struct {
	Code string `json:"code"`
}

// UseCodeResp use code response data type
type UseCodeResp struct {
	AccessToken       string    `json:"accessToken"`
	ValidUntil        time.Time `json:"validUntil"`
	RefreshToken      string    `json:"refreshToken"`
	RefreshValidUntil time.Time `json:"refreshValidUntil"`
}

// RefreshTokenReq revoke refresh token request data type
type RefreshTokenReq struct {
	RefreshToken string `json:"refreshToken"`
}

// RefreshAccessTokenResp refresh access token response data type
type RefreshAccessTokenResp struct {
	AccessToken string `json:"accessToken"`
}

// CheckAccessTokenReq checks access token request data type
type CheckAccessTokenReq struct {
	AccessToken string `json:"accessToken"`
}
