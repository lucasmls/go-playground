package infra

// DecodedJWT ...
type DecodedJWT struct {
	UserID string `json:"userId"`
	Exp    int64  `json:"exp"`
}
