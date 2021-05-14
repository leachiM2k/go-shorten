package urlutil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru"
	"github.com/jonboulle/clockwork"
	"github.com/mrcrgl/pflog/log"
	"github.com/pkg/errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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
		log.Fatalf("cannot create urlutil handler: %+v", err)
	}

	return &ApiHandler{
		Handler: handler,
	}
}

func GetHandler() (*Handler, error) {
	if handler == nil {
		m.Lock()
		defer m.Unlock()

		cache, _ := lru.NewARC(1024)

		handler := NewHandler(clockwork.NewRealClock(), cache)
		return handler, nil
	}

	m.RLock()
	defer m.RUnlock()
	return handler, nil
}

// GetUrlMetaHandler godoc
// @Summary Get all stats for a code
// @Description Get all stats for a code
// @ID getUrlMetaData
// @Accept  json
// @Produce  json
// @Param url query string true "URL to get meta information from"
// @Success 200 {object} interfaces.HTMLMeta
// @Failure 500 {string} string "fail"
// @Router /url/meta [get]
func (m *ApiHandler) GetUrlMetaHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	url := ctx.Query("url")
	if ok, _ := regexp.MatchString("^http(s?)://", url); !ok {
		ctx.String(http.StatusBadRequest, "No url specified or does not start with http:// or https://")
		return
	}

	entity, err := m.Handler.GetUrlMeta(url)
	if err != nil {
		ctx.AbortWithError(600, err)
		return
	}

	ctx.JSON(http.StatusOK, entity)
}

// GetQrCodeHandler godoc
// @Summary Generate QR Code
// @Description Generate QR Code for an URL
// @ID getQrCodeForUrl
// @Accept  json
// @Produce  json
// @Param url query string true "URL for QR Code"
// @Param format query string false "image format (svg or png)"
// @Param width query int false "for PNG: width of image"
// @Param height query int false "for PNG: height of image"
// @Success 200
// @Failure 500 {string} string "fail"
// @Router /url/qrcode [get]
func (m *ApiHandler) GetQrCodeHandler(ctx *gin.Context) {
	if m.Handler == nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot create handler"))
		return
	}

	url := ctx.Query("url")
	if ok, _ := regexp.MatchString("^http(s?)://", url); !ok {
		ctx.String(http.StatusBadRequest, "No url specified or does not start with http:// or https://")
		return
	}

	var err error

	switch strings.ToUpper(ctx.DefaultQuery("format", "PNG")) {
	case "SVG":
		ctx.Writer.Header().Set("Content-type", "image/svg+xml")
		err = m.Handler.GetQrCodeSvg(ctx.Writer, url)
	case "PNG":
		ctx.Writer.Header().Set("Content-type", "image/png")

		width, _ := strconv.ParseInt(ctx.DefaultQuery("width", "200"), 10, 32)
		height, _ := strconv.ParseInt(ctx.DefaultQuery("height", "200"), 10, 32)

		err = m.Handler.GetQrCodePng(ctx.Writer, url, int(width), int(height))
	default:
		ctx.AbortWithError(http.StatusNotImplemented, fmt.Errorf("unknown format"))
	}

	if err != nil {
		ctx.AbortWithError(600, err)
		return
	}
}
