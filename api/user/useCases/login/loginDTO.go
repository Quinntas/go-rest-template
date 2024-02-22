package login

type DTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Token      string `json:"token"`
	ExpiresIn  int    `json:"expiresIn"`
	ExpireDate string `json:"expireDate"`
}

type PrivateTokenClaim struct {
	Id    int
	Email string
	UUID  string
}

type PublicTokenClaim struct {
	UUID string
}
