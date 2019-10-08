package config

var (
	PORT = 8000
	SECRET_KEY="reqrqr3154"
)

func GetKey() string{
	return SECRET_KEY
}