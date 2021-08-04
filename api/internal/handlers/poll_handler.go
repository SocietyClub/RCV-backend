package handlers

import (
	"encoding/json"
	"net/http"
	"reflect"

	sv "github.com/core-go/service"
	"github.com/gorilla/mux"

	. "go-service/internal/models"
	. "go-service/internal/services"
)

type PollHandler struct {
	service PollService
}

func NewPollHandler(service PollService) *PollHandler {
	return &PollHandler{service: service}
}

func (h *PollHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func (h *PollHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	result, err := h.service.Load(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func (h *PollHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var poll Poll
	er1 := json.NewDecoder(r.Body).Decode(&poll)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	result, er2 := h.service.Insert(r.Context(), &poll)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func (h *PollHandler) Update(w http.ResponseWriter, r *http.Request) {
	var poll Poll
	er1 := json.NewDecoder(r.Body).Decode(&poll)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(poll.Id) == 0 {
		poll.Id = id
	} else if id != poll.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}

	result, er2 := h.service.Update(r.Context(), &poll)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func (h *PollHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	ids := []string{"id"}

	var poll Poll
	pollType := reflect.TypeOf(poll)
	_, jsonMap := sv.BuildMapField(pollType)
	body, _ := sv.BuildMapAndStruct(r, &poll)
	if len(poll.Id) == 0 {
		poll.Id = id
	} else if id != poll.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	json, er1 := sv.BodyToJson(r, poll, body, ids, jsonMap, nil)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}

	result, er2 := h.service.Patch(r.Context(), json)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func (h *PollHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	result, err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func respond(w http.ResponseWriter, result interface{}) {
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}