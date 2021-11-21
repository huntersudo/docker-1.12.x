package build

import "github.com/docker/docker/api/server/router"

// buildRouter is a router to talk with the build controller
type buildRouter struct {
	backend Backend
	routes  []router.Route
}

// NewRouter initializes a new build router
func NewRouter(b Backend) router.Router {
	r := &buildRouter{
		backend: b,
	}
	r.initRoutes()
	return r
}

// Routes returns the available routers to the build controller
func (r *buildRouter) Routes() []router.Route {
	return r.routes
}

func (r *buildRouter) initRoutes() {
	//TODO-SML :docker build 在 /docker/builder中
	//TODO-SML :docker image 在 /docker/daemon/image中
	r.routes = []router.Route{
		router.Cancellable(router.NewPostRoute("/build", r.postBuild)),
	}
}
