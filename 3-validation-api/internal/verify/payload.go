package verify

type VerifyRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyResponse struct {
	IsValid bool `json:"is_valid"`
}
