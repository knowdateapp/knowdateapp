package common

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/deepmap/oapi-codegen/v2/pkg/util"
)

func DecodeJsonRequest[T any](r *http.Request, result *T) error {
	if err := isMediaTypeJson(r.Header.Get("Content-Type")); err != nil {
		return err
	}

	return json.NewDecoder(r.Body).Decode(result)
}

func DecodeJsonMultipart[T any](part *multipart.Part, result *T) error {
	mediaType := part.Header.Get("Content-Type")
	if len(mediaType) > 0 {
		if err := isMediaTypeJson(mediaType); err != nil {
			return err
		}
	}

	return json.NewDecoder(part).Decode(result)
}

func isMediaTypeJson(mediaType string) error {
	if !util.IsMediaTypeJson(mediaType) {
		return fmt.Errorf("invalid content type: expected json, presented %s", mediaType)
	}
	return nil
}
