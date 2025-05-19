package auth

type LoginRequest struct {
	Number string `json:"number" validate:"required,e164"`
}

type LoginResponse struct {
	SessionId string `json:"session_id"`
}

type VerifyRequest struct {
	SessionId string `json:"session_id" validate:"required,len=12"`
	Code      int    `json:"code" validate:"required"`
}

type VerifyResponse struct {
	Token string `json:"access_token"`
}
