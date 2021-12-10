package handle

import (
	"golang-memory/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModelHandler interface {
	GetAll(*gin.Context)
}

type modelHandler struct {
	repo repository.ModelRepository
}

func NewModelHandler() ModelHandler {
	return &modelHandler{
		repo: repository.NewModelRepository(),
	}
}

func (h *modelHandler) GetAll(ctx *gin.Context){
	model, err := h.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, model)
	// response := new(utils.ApiResponse)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// defer utils.MapPanic(r, response)
	// if r.Method == http.MethodGet {
	// 	w.WriteHeader(http.StatusOK)
	// 	utils.CreateResponse(w, &utils.InMemStore)
	// 	return
	// }
	// err := http.StatusText(http.StatusMethodNotAllowed)
	// response.Error = err
	// utils.CreateResponse(w, &utils.InMemStore)
	// w.WriteHeader(http.StatusMethodNotAllowed)
	// return
}