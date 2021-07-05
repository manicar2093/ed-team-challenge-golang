package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Add("Content-Type", ContentTypeJSON)
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&data)
}

func MakeDownloadableFile(w http.ResponseWriter, f *os.File, mimeType string) error {

	fileStats, e := f.Stat()
	if e != nil {
		return e
	}

	name := fileStats.Name()

	w.Header().Set("Content-Type", ContentTypePNG)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename="%s"`, name))
	w.Header().Set("File-Name", name)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	fileBytes, e := io.ReadAll(f)
	if e != nil {
		return e
	}
	fileBuffer := bytes.NewBuffer(fileBytes)

	_, e = fileBuffer.WriteTo(fileBuffer)

	if e != nil {
		return e
	}

	return e
}
