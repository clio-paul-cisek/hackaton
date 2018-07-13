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
	uid := ps.ByName("uuid")
	data, err := lp.Repository.Fetch(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if len(data) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	idKeyMap := map[string]string{}

	err = json.Unmarshal(data, &idKeyMap)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var responseData []byte

	if key, ok := idKeyMap[uid]; ok {
		responseData, err = lp.Repository.Fetch(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(responseData)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
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

	kg := create.NewKeyGenerator(
		rcRequest.AccountID,
		rcRequest.Description,
		rcRequest.Schedule.Start,
		rcRequest.Schedule.IntervalUnit,
		rcRequest.Schedule.IntervalDelay,
	)

	key := kg.GenerateKey()

	data, err := lp.Repository.Fetch(key)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// RecordCreate function handles raw response recording
func (lp LawPay) RecordCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rcrResponse := &create.SuccessResponse{}
	requestDecoder := json.NewDecoder(r.Body)
	err := requestDecoder.Decode(rcrResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	data, err := json.Marshal(rcrResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	kg := create.NewKeyGenerator(
		rcrResponse.AccountID,
		rcrResponse.Description,
		rcrResponse.Schedule.Start,
		rcrResponse.Schedule.IntervalUnit,
		rcrResponse.Schedule.IntervalDelay,
	)

	key := kg.GenerateKey()

	err = lp.Repository.Put(key, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	idKeyMap := map[string]string{
		rcrResponse.ID: key,
	}

	keyData, err := json.Marshal(idKeyMap)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = lp.Repository.Put(rcrResponse.ID, keyData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
