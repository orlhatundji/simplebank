package token

import "time"

// Maker is the interface that wraps the basic methods for token creation.
type Maker interface {
	// CreateToken creates a new token for a specific username and duration.
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not.
	VerifyToken(token string) (*Payload, error)
}
