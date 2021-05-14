package urlutil

import (
	"github.com/aaronarduino/goqrsvg"
	"github.com/ajstarks/svgo"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/hashicorp/golang-lru"
	"github.com/jonboulle/clockwork"
	"github.com/leachim2k/go-shorten/pkg/dataservice/interfaces"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"image/png"
	"io"
	"net/http"
)

type Handler struct {
	Clock clockwork.Clock
	Cache *lru.ARCCache
}

func NewHandler(clock clockwork.Clock, cache *lru.ARCCache) *Handler {
	return &Handler{
		Clock: clock,
		Cache: cache,
	}
}

func (m *Handler) GetUrlMeta(url string) (*interfaces.HTMLMeta, error) {
	meta, ok := m.Cache.Get(url)

	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return nil, errors.New("can not fetch given url")
		}
		defer resp.Body.Close()

		meta = extract(resp.Body)
		m.Cache.Add(url, meta)
	}

	return meta.(*interfaces.HTMLMeta), nil
}

type ImageFormat string

const (
	PNG ImageFormat = "PNG"
	SVG             = "SVG"
)

func (m *Handler) GetQrCodePng(writer io.Writer, url string, width int, height int) error {
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return err
	}

	qrCode, err = barcode.Scale(qrCode, width, height)
	if err != nil {
		return err
	}

	err = png.Encode(writer, qrCode)
	if err != nil {
		return err
	}

	return nil
}

func (m *Handler) GetQrCodeSvg(writer io.Writer, url string) error {
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return err
	}

	s := svg.New(writer)
	qs := goqrsvg.NewQrSVG(qrCode, 5)
	qs.StartQrSVG(s)
	err = qs.WriteQrSVG(s)
	if err != nil {
		return err
	}
	s.End()

	return nil
}

func extract(resp io.Reader) *interfaces.HTMLMeta {
	z := html.NewTokenizer(resp)

	titleFound := false

	hm := new(interfaces.HTMLMeta)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return hm
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == `body` {
				return hm
			}
			if t.Data == "title" {
				titleFound = true
			}
			if t.Data == "meta" {
				desc, ok := extractMetaProperty(t, "description")
				if ok {
					hm.Description = desc
				}

				ogTitle, ok := extractMetaProperty(t, "og:title")
				if ok {
					hm.Title = ogTitle
				}

				ogDesc, ok := extractMetaProperty(t, "og:description")
				if ok {
					hm.Description = ogDesc
				}

				ogImage, ok := extractMetaProperty(t, "og:image")
				if ok {
					hm.Image = ogImage
				}

				ogSiteName, ok := extractMetaProperty(t, "og:site_name")
				if ok {
					hm.SiteName = ogSiteName
				}
			}
		case html.TextToken:
			if titleFound {
				t := z.Token()
				hm.Title = t.Data
				titleFound = false
			}
		}
	}
	return hm
}

func extractMetaProperty(t html.Token, prop string) (content string, ok bool) {
	for _, attr := range t.Attr {
		if attr.Key == "property" && attr.Val == prop {
			ok = true
		}

		if attr.Key == "content" {
			content = attr.Val
		}
	}

	return
}
