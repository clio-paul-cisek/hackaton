package api

import (
	"encoding/json"
	"net/http"

	"github.com/clio-paul-cisek/hackaton/api"
	"github.com/clio-paul-cisek/hackaton/lawpay/recurring/create"
	"github.com/clio-paul-cisek/hackaton/repository"
	"github.com/julienschmidt/httprouter"
)

// LawPay is a struct responsible for handling requests
type LawPay struct {
	*api.Controller
}

// New function return pointer to new LawPay struct
func New(repository repository.Connector) *LawPay {
	return &LawPay{
		api.New(repository),
	}
}

// Get function allows us to fetch response with given id
func (lp LawPay) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	data, err := lp.Repository.Fetch(ps.ByName("uuid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

// Create function allows us to create new stubbed responses
func (lp LawPay) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rcRequest := &create.RecurringCharge{}
	requestDecoder := json.NewDecoder(r.Body)
	err := requestDecoder.Decode(rcRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	data, err := lp.Repository.Fetch(rcRequest.AccountID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// CreateResponse function allows us to create stubbed response
func (lp LawPay) CreateResponse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rcRequest := &create.RecurringChargeGenerateRequest{}
	requestDecoder := json.NewDecoder(r.Body)
	err := requestDecoder.Decode(rcRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	mapper := create.NewMapper(rcRequest)
	data, err := mapper.MapRequest()

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = lp.Repository.Put(rcRequest.AccountID, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
