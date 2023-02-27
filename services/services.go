package services

type Services struct {
	users *IUsers
}

func New(userSrv *IUsers) *Services {
	return &Services{userSrv}
}
