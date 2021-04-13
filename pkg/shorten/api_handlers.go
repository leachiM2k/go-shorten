package shorten

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jonboulle/clockwork"
	"log"
	"net/http"
	"sync"
)

var awsHandler *ShortenHandler
var m sync.RWMutex

type ApiHandler struct {
	Mutex   sync.RWMutex
	Handler *ShortenHandler
}

func NewApiHandler() *ApiHandler {
	handler, err := GetHandler()
	if err != nil {
		log.Fatalf("cannot create AWS handler: %+v", err)
	}

	return &ApiHandler{
		Handler: handler,
	}
}

func GetHandler() (*ShortenHandler, error) {
	if awsHandler == nil {
		m.Lock()
		defer m.Unlock()

		backend := NewInmemoryBackend()

		handler := NewHandler(clockwork.NewRealClock(), backend)
		return handler, nil
	}

	m.RLock()
	defer m.RUnlock()
	return awsHandler, nil
}

func (m *ApiHandler) MissingCodeHandler(ctx *gin.Context) {
	code := ctx.Param("code")
	if code == "" {
		ctx.String(http.StatusBadRequest, "No code specified")
		return
	}
}

// GetHandler godoc
// @Summary Get short's info
// @Description Get all stored information for a specified short
// @ID read
// @Accept  json
// @Produce  json
// @Param code path string true "short code"
// @Success 200 {object} Entity
// @Failure 500 {string} string "fail"
// @Router /shorten/{code} [get]
func (m *ApiHandler) GetHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	code := ctx.Param("code")

	entity, err := m.Handler.Get(code)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if entity == nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("code not found"))
		return
	}

	ctx.JSON(http.StatusOK, entity)
}

// HandleCodeHandler godoc
// @Summary Handle a short
// @Description Return the right link for short code or "not found" if expired, not started or max count was reached
// @ID handle
// @Accept  json
// @Produce  plain
// @Param code path string true "short code"
// @Success 302 {string} string "Link to follow"
// @Failure 500 {string} string "fail"
// @Router /shorten/handle/{code} [get]
func (m *ApiHandler) HandleCodeHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	code := ctx.Param("code")

	entity, err := m.Handler.Get(code)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	link, err := m.Handler.ConvertEntityToLink(entity)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if link == "" {
		ctx.AbortWithError(http.StatusNotFound, errors.New("code not found"))
		return
	}

	ctx.Redirect(http.StatusFound, link)
}

// AddHandler godoc
// @Summary Add a new short
// @Description Create a new short. Create random code if not specified.
// @ID create
// @Accept  json
// @Produce  json
// @Param account body CreateRequest true "Create Request"
// @Success 200 {object} Entity
// @Router /shorten/ [post]
func (m *ApiHandler) AddHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	var createRequest CreateRequest
	err := ctx.BindJSON(&createRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if createRequest.Owner == nil || createRequest.Link == nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("mandatory parameters missing (owner, link)"))
		return
	}

	entity, err := m.Handler.Add(createRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, entity)
}

// UpdateHandler godoc
// @Summary Update a short
// @Description Updates several fields of a short, while maintaining count, owner and creation date
// @ID update
// @Accept  json
// @Produce  json
// @Param code path string true "short code"
// @Param account body UpdateRequest true "Update Request"
// @Success 200 {object} Entity
// @Router /shorten/{code} [put]
func (m *ApiHandler) UpdateHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	var updateRequest UpdateRequest
	err := ctx.BindJSON(&updateRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	code := ctx.Param("code")

	entity, err := m.Handler.Update(code, updateRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, entity)
}

// DeleteHandler godoc
// @Summary Delete a short
// @Description Delete a short
// @ID delete
// @Produce  json
// @Param code path string true "short code"
// @Success 204
// @Router /shorten/{code} [delete]
func (m *ApiHandler) DeleteHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	code := ctx.Param("code")

	err := m.Handler.Delete(code)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
