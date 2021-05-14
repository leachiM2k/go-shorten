package shorten

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jonboulle/clockwork"
	"github.com/leachim2k/go-shorten/pkg/auth/middleware"
	"github.com/leachim2k/go-shorten/pkg/dataservice"
	"github.com/leachim2k/go-shorten/pkg/dataservice/interfaces"
	"github.com/mrcrgl/pflog/log"
	"net/http"
	"sync"
)

var handler *Handler
var m sync.RWMutex

type ApiHandler struct {
	Mutex   sync.RWMutex
	Handler *Handler
}

func NewApiHandler() *ApiHandler {
	handler, err := GetHandler()
	if err != nil {
		log.Fatalf("cannot create shortener handler: %+v", err)
	}

	return &ApiHandler{
		Handler: handler,
	}
}

func GetHandler() (*Handler, error) {
	if handler == nil {
		m.Lock()
		defer m.Unlock()

		backend := dataservice.GetDataServiceByConfig()

		handler := NewHandler(clockwork.NewRealClock(), backend)
		return handler, nil
	}

	m.RLock()
	defer m.RUnlock()
	return handler, nil
}

func (m *ApiHandler) MissingCodeHandler(ctx *gin.Context) {
	code := ctx.Param("code")
	if code == "" {
		ctx.String(http.StatusBadRequest, "No code specified")
		return
	}
}

// GetAllHandler godoc
// @Summary Get all user shorts
// @Description Get all shorts for an user
// @ID readAll
// @Accept  json
// @Produce  json
// @Success 200 {array} interfaces.Entity
// @Failure 500 {string} string "fail"
// @Router /shorten [get]
func (m *ApiHandler) GetAllHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	jwtClaim := ctx.MustGet("JWT_CLAIMS").(*middleware.AuthCustomClaims)

	entities, err := m.Handler.GetAll(jwtClaim.Subject)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if entities == nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("code not found"))
		return
	}

	ctx.JSON(http.StatusOK, entities)
}

// GetHandler godoc
// @Summary Get short's info
// @Description Get all stored information for a specified short
// @ID read
// @Accept  json
// @Produce  json
// @Param code path string true "short code"
// @Success 200 {object} interfaces.Entity
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
	code := ctx.Param("code")
	link, err := m.HandleCode(code, ctx.ClientIP(), ctx.Request.UserAgent(), ctx.Request.Referer())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if link == nil || *link == "" {
		ctx.AbortWithError(http.StatusNotFound, errors.New("code not found"))
		return
	}

	switch ctx.NegotiateFormat(gin.MIMEHTML, gin.MIMEJSON) {
	case gin.MIMEJSON:
		ctx.JSON(200, *link)
	default:
		ctx.Redirect(http.StatusFound, *link)
	}
}

func (m *ApiHandler) HandleCode(code string, clientIp string, userAgent string, referer string) (*string, error) {
	if m.Handler == nil {
		return nil, errors.New("cannot create handler")
	}

	entity, err := m.Handler.Get(code)
	if entity == nil || err != nil {
		return nil, err
	}

	go func() {
		_, errStat := m.Handler.AddStat(entity.ID, clientIp, userAgent, referer)
		if errStat != nil {
			log.Warningf("Could not write stats: %s", err)
		}
	}()

	link, err := m.Handler.ConvertEntityToLink(entity)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

// AddHandler godoc
// @Summary Add a new short
// @Description Create a new short. Create random code if not specified.
// @ID create
// @Accept  json
// @Produce  json
// @Param account body interfaces.CreateRequest true "Create Request"
// @Success 200 {object} interfaces.Entity
// @Router /shorten/ [post]
func (m *ApiHandler) AddHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	var createRequest interfaces.CreateRequest
	err := ctx.BindJSON(&createRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	jwtClaim := ctx.MustGet("JWT_CLAIMS").(*middleware.AuthCustomClaims)
	createRequest.Owner = &jwtClaim.Subject

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
// @Param account body interfaces.UpdateRequest true "Update Request"
// @Success 200 {object} interfaces.Entity
// @Router /shorten/{code} [put]
func (m *ApiHandler) UpdateHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	var updateRequest interfaces.UpdateRequest
	err := ctx.BindJSON(&updateRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	jwtClaim := ctx.MustGet("JWT_CLAIMS").(*middleware.AuthCustomClaims)
	updateRequest.Owner = &jwtClaim.Subject

	if updateRequest.Owner == nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("mandatory parameters missing (owner)"))
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

	jwtClaim := ctx.MustGet("JWT_CLAIMS").(*middleware.AuthCustomClaims)

	code := ctx.Param("code")

	err := m.Handler.Delete(jwtClaim.Subject, code)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetStatsHandler godoc
// @Summary Get all stats for a code
// @Description Get all stats for a code
// @ID readStats
// @Accept  json
// @Produce  json
// @Param code path string true "short code"
// @Success 200 {array} interfaces.StatEntity
// @Failure 500 {string} string "fail"
// @Router /shorten/{code}/stats [get]
func (m *ApiHandler) GetStatsHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	code := ctx.Param("code")

	entities, err := m.Handler.GetStats(code)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if entities == nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("code not found"))
		return
	}

	ctx.JSON(http.StatusOK, entities)
}
