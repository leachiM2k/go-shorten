package middleware

import (
	"bytes"
	"github.com/mrcrgl/pflog/log"
	"io"
	"io/ioutil"
	"net/http"
)

func FullAuditLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || r.Method == "HEAD" || r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		save := r.Body
		save, r.Body, _ = drainBody(r.Body)

		defer r.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(r.Body)

		r.Body = save

		log.Infof("[%s] [%s] %s", r.Method, r.RequestURI, string(bodyBytes))
		/*
			logItem := models.LogItem{
				Route:     r.RequestURI,
				Method:    r.Method,
				Payload:   bodyBytes,
				Timestamp: time.Time{},
			}
			_ = models.PersistLog(logItem)
		*/
		next.ServeHTTP(w, r)
	})
}

func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == nil || b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
