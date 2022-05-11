package token

type Token interface {
	Make(payload *Payload) (string, error)
	Verify(token string) (*Payload, error)
}
