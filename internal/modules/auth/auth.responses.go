package auth

type ResponseType map[string]any

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string `json:"expires_at"`
}

type TokensResponseType struct {
	Tokens Tokens `json:"tokens"`
}

func TokensResponse(tokens Tokens) ResponseType {
	return ResponseType{
		"tokens": tokens,
	}
}
