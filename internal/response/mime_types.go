package response

var (
	ContentTypeJSON = ContentType{Extension: "json", Mime: "application/json"}
	ContentTypePNG  = ContentType{Extension: "png", Mime: "image/png"}
)

type ContentType struct {
	// Extension indicates the extension use by the type
	Extension string
	// Type is the representation of http content type
	Mime string
}
