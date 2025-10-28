package routes

import "github.com/dominicsieli/go-server/config"

func Post(route string) string {
	return "POST " + config.Envs.APIVersion + route
}

func Get(route string) string {
	return "GET " + config.Envs.APIVersion + route
}

func Patch(route string) string {
	return "PATCH " + config.Envs.APIVersion + route
}

func Delete(route string) string {
	return "DELETE " + config.Envs.APIVersion + route
}
