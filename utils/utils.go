package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
	ContentTypePNG    = "image/png"
)

type ControllerUtils struct{}

func (u ControllerUtils) JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Add("Content-Type", ContentTypeJSON)
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&data)
}

func (u ControllerUtils) MakeDownloadableFile(w http.ResponseWriter, f *os.File, mimeType string) (e error) {

	fileStats, e := f.Stat()
	if e != nil {
		return
	}

	name := fileStats.Name()

	w.Header().Set("Content-Type", ContentTypePNG)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename="%s"`, name))
	w.Header().Set("File-Name", name)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	fileBytes, e := io.ReadAll(f)
	if e != nil {
		return
	}
	fileBuffer := bytes.NewBuffer(fileBytes)

	_, e = fileBuffer.WriteTo(fileBuffer)

	if e != nil {
		return
	}

	return
}

// Floor2Positions32 floors the given float value to two 0 positions
func Floor2Positions32(v float64) float32 {
	return float32(math.Floor(v*100) / 100)
}

func SeparateByCommas(elements ...string) string {
	builder := strings.Builder{}

	comma := len(elements) - 1
	for i, v := range elements {

		builder.WriteString(v)
		if i != comma {
			builder.WriteString(",")
		}
	}

	return builder.String()
}
