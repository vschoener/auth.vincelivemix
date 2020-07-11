package controllers

type Controllers struct {
	AuthController AuthController
}

func ProvideControllers(authController AuthController) Controllers {
	return Controllers{
		AuthController: authController,
	}
}
