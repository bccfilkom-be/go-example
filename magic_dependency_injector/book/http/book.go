package handler

import (
	"net/http"

	"github.com/bccfilkom-be/go-example/magic_dependency_injector/book/usecase"
	"github.com/bccfilkom-be/go-example/magic_dependency_injector/db/postgresql"
	"github.com/kataras/iris/v12"
)

type IBookHandler interface {
	List(c iris.Context)
	Create(c iris.Context)
}

type handler struct {
	usecase usecase.IBookUsecase
}

func NewBookHandler(usecase usecase.IBookUsecase) IBookHandler {
	return &handler{usecase}
}

func (h *handler) List(c iris.Context) {
	books, err := h.usecase.ListBooks(c.Clone())
	if err != nil {
		c.StopWithProblem(http.StatusInternalServerError, iris.NewProblem().
			DetailErr(err))
		return
	}
	c.JSON(books)
}

func (h *handler) Create(c iris.Context) {
	var b postgresql.Book
	if err := c.ReadJSON(&b); err != nil {
		c.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			DetailErr(err))
		return
	}
	c.StatusCode(iris.StatusCreated)
}
