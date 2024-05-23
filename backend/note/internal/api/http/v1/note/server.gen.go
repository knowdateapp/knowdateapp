// Package note provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package note

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// Note defines model for Note.
type Note struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Workspace string `json:"workspace"`
}

// Workspace defines model for Workspace.
type Workspace = string

// CreateNoteResponse defines model for CreateNoteResponse.
type CreateNoteResponse struct {
	Id string `json:"id"`
}

// DefaultErrorResponse defines model for DefaultErrorResponse.
type DefaultErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ListNotesResponse defines model for ListNotesResponse.
type ListNotesResponse struct {
	Notes []Note `json:"notes"`
}

// CreateNoteRequest defines model for CreateNoteRequest.
type CreateNoteRequest struct {
	Title string `json:"title"`
}

// CreateNoteJSONBody defines parameters for CreateNote.
type CreateNoteJSONBody struct {
	Title string `json:"title"`
}

// CreateNoteJSONRequestBody defines body for CreateNote for application/json ContentType.
type CreateNoteJSONRequestBody CreateNoteJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /{workspace}/notes)
	ListNotes(w http.ResponseWriter, r *http.Request, workspace Workspace)

	// (POST /{workspace}/notes)
	CreateNote(w http.ResponseWriter, r *http.Request, workspace Workspace)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /{workspace}/notes)
func (_ Unimplemented) ListNotes(w http.ResponseWriter, r *http.Request, workspace Workspace) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /{workspace}/notes)
func (_ Unimplemented) CreateNote(w http.ResponseWriter, r *http.Request, workspace Workspace) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// ListNotes operation middleware
func (siw *ServerInterfaceWrapper) ListNotes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "workspace" -------------
	var workspace Workspace

	err = runtime.BindStyledParameterWithOptions("simple", "workspace", chi.URLParam(r, "workspace"), &workspace, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workspace", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListNotes(w, r, workspace)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateNote operation middleware
func (siw *ServerInterfaceWrapper) CreateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "workspace" -------------
	var workspace Workspace

	err = runtime.BindStyledParameterWithOptions("simple", "workspace", chi.URLParam(r, "workspace"), &workspace, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workspace", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateNote(w, r, workspace)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/{workspace}/notes", wrapper.ListNotes)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/{workspace}/notes", wrapper.CreateNote)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RVS4vbPhD/KmL+fygEEWe7Wyi+9XVYKKW0hR6WPWjlcaKtLamacUoI/u5FshO7tUMC",
	"u735MZr5PWZGe9Cu9s6iZYJ8D14FVSNjSG/fXfhBXmmMLwWSDsazcRZy+LYxJI7RosDSWCTBGxS6CQEt",
	"i1+H00uQYOIhr3gDEqyqEXI4/gcJAX82JmABOYcGJZDeYK1iWd75GEwcjF1D27ZdMBK/dYXBhPNdQMX4",
	"yTF+6X7Fj9pZRpselfeV0SpCzx4p4t+PKvjgPAbuc7HhCucKj0He9WH38hDmHh5Rc8TXRZJ3lqbous9P",
	"gGeK89hMMQtMTixEYR2j2CgSD4hWUKM1EpVNVe2ETrCLJbQS3mOpmoo/hODCM5DQrpiTWEKNRGp9gfwp",
	"wxB/Cd2egzh4I2pXYCVKF4ZPMS2Jm8VCKFuIV4tFYv/REEf76BmoR8E7Ixnr9PB/wBJy+C8bJjHrTlMW",
	"q0YEPTkVgtpNtOhSXiJBJCJcKRrCkKynZcrWl4toUsXLuk6enBU5mu1LuvWQaXxuno+xpTuxjN58vhU+",
	"uK2JFiorjGUMpdKYLE7dbOw6GVsrq9bxpdfgyCTRj5lAwhYDdemvlqvlKrJyHq3yBnK4Xq6W1yDTRksK",
	"Zfsj9DY7mrxGPoEWbeGdsbEduQm225zRlxc0oIoWpLa6LXr3UhumusOevpvvoSEkG/Z4e//Xcnq5Wp1q",
	"wmNcNh2A1Fppns6fnl0e0UxWaxo1cCvBOzorWLeXSKixXFO1hq37ZLkOt83uNNfRhZRNb6N2IvrVedlm",
	"ro1/onpaABi2B3H+1P5r+pNGqHJaVaLALVbO13H3SWhCBTlsmH2eZSlg44jz16ubVaa8ybZX0N63vwMA",
	"AP//89gfA2YIAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
