package routes

func Post(route string) string {
	return "POST " + route
}

func Get(route string) string {
	return "GET " + route
}

func Patch(route string) string {
	return "PATCH " + route
}

func Delete(route string) string {
	return "DELETE " + route
}
