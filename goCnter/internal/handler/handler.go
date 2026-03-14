package handler

import (
	"fmt"
	"goCnter/internal/cache"
	"goCnter/internal/model"
	"net/http"
	"strconv"
)

type Handle struct {
	Cacher cache.Casher
	TaskCh chan *model.Task
}

func NewHandle(cacher cache.Casher) *Handle {
	return &Handle{
		Cacher: cacher,
		TaskCh: make(chan *model.Task, 5),
	}
}

func (h *Handle) GetCnt(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	cnt, err := h.Cacher.GetCnt(id)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(cnt)))
}

func (h *Handle) IncrementCnt(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err := h.Cacher.IncrementCnt(id)
	if err != nil {
		http.Error(w, "server error", http.StatusBadRequest)
	}

	//
	go h.sendToTaskChan(id)

	// Kafka
	// go func() {
	// 	v, err := h.Cacher.GetCnt(id)
	// 	if err != nil {
	// 		fmt.Println(fmt.Errorf("error: %v", err))
	// 		return
	// 	}

	// 	msg := kafka.Mess{
	// 		ID:     id,
	// 		Type:   "increment",
	// 		NewVal: strconv.Itoa(v),
	// 	}

	// 	err = kafka.ProduceMess(msg)
	// 	if err != nil {
	// 		fmt.Println(fmt.Errorf("error: %v", err))
	// 		return
	// 	}
	// }()

	//
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("cnt successfully incremented"))
}

func (h *Handle) sendToTaskChan(id string) {
	val, err := h.Cacher.GetCnt(id)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		return
	}

	h.TaskCh <- &model.Task{ID: id, Count: val}
}

func (h *Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 	GET /counter/{id}
	// POST /increment/{id}

	switch {
	case r.URL.Path == "/counter" && r.Method == "GET":
		h.GetCnt(w, r)
	case r.URL.Path == "/increment" && r.Method == "POST":
		h.IncrementCnt(w, r)
	default:
		http.NotFound(w, r)
	}
}
