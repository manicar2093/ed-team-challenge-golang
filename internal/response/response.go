package response

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Add("Content-Type", ContentTypeJSON.Mime)
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&data)
}

// MakeDownloadableFile adds headers to dowload by client the data provided. The name of the file does not need extension due mimetype creates it
func MakeDownloadableFile(w http.ResponseWriter, f io.WriterTo, name string, mimeType ContentType) error {

	fileName := fmt.Sprintf("%s.%s", name, mimeType.Extension)

	w.Header().Set("Content-Type", mimeType.Mime)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename="%s"`, fileName))
	w.Header().Set("File-Name", fileName)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	_, e := f.WriteTo(w)

	if e != nil {
		return e
	}

	return e
}
