package api

import (
	"net/http"

	"github.com/clio-paul-cisek/hackaton/repository"
	"github.com/julienschmidt/httprouter"
)

// API interface describes API endpoint behaviour
type API interface {
	Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

// Controller type is a basic api controller struct which contains repository struct
type Controller struct {
	Repository repository.Connector
}

// New function created new Controller struct
func New(repository repository.Connector) *Controller {
	return &Controller{repository}
}
