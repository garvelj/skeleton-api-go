package api

func (a *Api) RoutesRegister() {
	public := a.Router.Group("/v1")

	public.GET("/skeleton", a.Ping)
}
