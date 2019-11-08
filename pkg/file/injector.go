package file

func ProvideFileDataSource() DataSource {
	return NewDataSource()
}

func ProvideRepository() Repository {
	return NewRepository(ProvideFileDataSource())
}

func ProvideUseCase() UseCase {
	return NewUseCase(ProvideRepository())
}

func ProvideHandler() Handler {
	return NewHandler(ProvideUseCase())
}

func ProvideRoute() Route {
	return NewRoute(ProvideHandler())
}
