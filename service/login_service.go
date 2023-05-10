package service

type LoginServiceInterface interface {
	Login(username string, password string) bool
}

type LoginServiceImplementation struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginServiceInterface {
	return &LoginServiceImplementation{
		authorizedUsername: "felicia",
		authorizedPassword: "bdbcdhbcx",
	}
}

func (service *LoginServiceImplementation) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
