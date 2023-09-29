package handler

import (
	"encoding/json"
	"encurtador_url/src/internal/entity"
	"encurtador_url/src/internal/repository"
	"net/http"
	"strings"
)

type RequestBodyInput struct {
	Url string `json:"url"`
}

type Handler struct {
	Repository repository.CodeRepositoryInterface
}

func NewHandler(repo repository.CodeRepositoryInterface) *Handler {
	return &Handler{
		Repository: repo,
	}
}

func (h *Handler) CreateCode(w http.ResponseWriter, r *http.Request) {
	var input RequestBodyInput
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	json.NewDecoder(r.Body).Decode(&input)

	code, err := entity.NewCode(input.Url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.Repository.Save(*code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *Handler) GetByCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	param := strings.Split(r.URL.Path, "/")

	if len(param) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("codigo nao fornecido"))
		return
	}

	c := param[2]

	res, err := h.Repository.Get(c)
	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("codigo nao encontrado"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(res)

	w.WriteHeader(http.StatusOK)
}
