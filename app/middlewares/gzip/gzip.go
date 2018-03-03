package gzip

// ref: https://github.com/gin-gonic/contrib/blob/master/gzip/gzip.go
// ref: https://github.com/gin-contrib/gzip/blob/master/gzip.go
import (
	"compress/gzip"
	"net/http"
	"path/filepath"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"
)

const (
	// BestCompression ...
	BestCompression = gzip.BestCompression
	// BestSpeed ...
	BestSpeed = gzip.BestSpeed
	// DefaultCompression ...
	DefaultCompression = gzip.DefaultCompression
	// NoCompression ...
	NoCompression = gzip.NoCompression
)

// Gzip ...
func Gzip(level int) gin.HandlerFunc {
	/*
		var gzPool sync.Pool
		gzPool.New = func() interface{} {
			gz, err := gzip.NewWriterLevel(ioutil.Discard, level)
			if err != nil {
				panic(err)
			}
			return gz
		}
	*/

	return func(c *gin.Context) {
		if !shouldCompress(c.Request) {
			return
		}

		gz, err := gzip.NewWriterLevel(c.Writer, level)
		if err != nil {
			return
		}
		/*
			gz := gzPool.Get().(*gzip.Writer)
			defer gzPool.Put(gz)
			gz.Reset(c.Writer)
		*/

		c.Header("Content-Encoding", "gzip")
		c.Header("Vary", "Accept-Encoding")
		c.Writer = &gzipWriter{c.Writer, gz}

		defer func() {
			c.Header("Content-Length", "0")
			gz.Close()
		}()
		c.Next()
	}
}

type gzipWriter struct {
	gin.ResponseWriter
	writer *gzip.Writer
}

func (g *gzipWriter) WriteString(s string) (int, error) {
	return g.writer.Write([]byte(s))
}

func (g *gzipWriter) Write(data []byte) (int, error) {
	return g.writer.Write(data)
}

func shouldCompress(req *http.Request) bool {
	if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		return false
	}
	extension := filepath.Ext(req.URL.Path)
	if len(extension) < 4 { // fast path
		return true
	}

	switch extension {
	case ".png", ".gif", ".jpeg", ".jpg":
		return false
	default:
		return true
	}
}