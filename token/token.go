package token

type Token interface {
	Make()
	Valid()
}
