package userVm

type LoginUserViewModel struct {
	UserName string `validate:"required"`
	Password string `validate:"required"`
}
