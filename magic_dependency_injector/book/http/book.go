package http

import (
	"net/http"

	"github.com/bccfilkom-be/go-example/magic_dependency_injector/book/usecase"
	"github.com/bccfilkom-be/go-example/magic_dependency_injector/db/postgresql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type handler struct {
	usecase usecase.IBookUsecase
}

func RegisterBookHTTP(r router.Party, usecase usecase.IBookUsecase) {
	h := handler{usecase}
	v1 := r.Party("/v1")
	books := v1.Party("/books")
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
	// TODO: put in use new generated resource id
	_, err := h.usecase.CreateBook(c.Clone(), b.Title)
	if err != nil {
		c.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			DetailErr(err))
		return
	}
	c.StatusCode(iris.StatusCreated)
}
