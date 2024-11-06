package http

import (
	"net/http"

	"github.com/bccfilkom-be/go-example/magic_dependency_injector/book/usecase"
	"github.com/bccfilkom-be/go-example/magic_dependency_injector/db/postgresql"
	"github.com/kataras/iris/v12"
)

type handler struct {
	usecase usecase.IBookUsecase
}

func RegisterBookHTTP(r *iris.Application, usecase usecase.IBookUsecase) {
	h := handler{usecase}
	books := r.Party("/books")
	books.Get("/", h.list)
	books.Post("/", h.create)
}

func (h *handler) list(c iris.Context) {
	books, err := h.usecase.ListBooks(c.Clone())
	if err != nil {
		c.StopWithProblem(http.StatusInternalServerError, iris.NewProblem().
			DetailErr(err))
		return
	}
	c.JSON(books)
}

func (h *handler) create(c iris.Context) {
	var b postgresql.Book

	if err := c.ReadJSON(&b); err != nil {
		c.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			DetailErr(err))
		return
	}
	c.StatusCode(iris.StatusCreated)
}
