// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package server

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
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	. "github.com/openclarity/vmclarity/pkg/uibackend/api/models"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of findings impact for the dashboard.
	// (GET /dashboard/findingsImpact)
	GetDashboardFindingsImpact(ctx echo.Context) error
	// Get a list of finding trends for all finding types.
	// (GET /dashboard/findingsTrends)
	GetDashboardFindingsTrends(ctx echo.Context, params GetDashboardFindingsTrendsParams) error
	// Get a list of riskiest assets for the dashboard.
	// (GET /dashboard/riskiestAssets)
	GetDashboardRiskiestAssets(ctx echo.Context) error
	// Get a list of riskiest regions for the dashboard.
	// (GET /dashboard/riskiestRegions)
	GetDashboardRiskiestRegions(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetDashboardFindingsImpact converts echo context to params.
func (w *ServerInterfaceWrapper) GetDashboardFindingsImpact(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetDashboardFindingsImpact(ctx)
	return err
}

// GetDashboardFindingsTrends converts echo context to params.
func (w *ServerInterfaceWrapper) GetDashboardFindingsTrends(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDashboardFindingsTrendsParams
	// ------------- Required query parameter "startTime" -------------

	err = runtime.BindQueryParameter("form", true, true, "startTime", ctx.QueryParams(), &params.StartTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter startTime: %s", err))
	}

	// ------------- Required query parameter "endTime" -------------

	err = runtime.BindQueryParameter("form", true, true, "endTime", ctx.QueryParams(), &params.EndTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter endTime: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetDashboardFindingsTrends(ctx, params)
	return err
}

// GetDashboardRiskiestAssets converts echo context to params.
func (w *ServerInterfaceWrapper) GetDashboardRiskiestAssets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetDashboardRiskiestAssets(ctx)
	return err
}

// GetDashboardRiskiestRegions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDashboardRiskiestRegions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetDashboardRiskiestRegions(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/dashboard/findingsImpact", wrapper.GetDashboardFindingsImpact)
	router.GET(baseURL+"/dashboard/findingsTrends", wrapper.GetDashboardFindingsTrends)
	router.GET(baseURL+"/dashboard/riskiestAssets", wrapper.GetDashboardRiskiestAssets)
	router.GET(baseURL+"/dashboard/riskiestRegions", wrapper.GetDashboardRiskiestRegions)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RaT3PiOBb/KirtHr1JZrb2wo0mJO1qSCgg3Ts1NQdhP0ATW3JLctJsF999S5KNjS2B",
	"3QM9N7Dek37vj98/+TuOeJpxBkxJPPiOMyJICgqE+QcsXtIU9E/K8AB/zUHscIAZ0Q8PywEW8DWnAmI8",
	"UCKHAMtoCynRfGsuUqLwAMdEwb+UJVe7TPNLJSjb4P0+wPCNpFkCDzRRILznWSJc37+9lVREqFOwK4K/",
	"DHyvd5AZZxKMwl7YK+PvbCwEN1JEnClgSv8kWZbQiCjK2e2fkjP9rDrtnwLWeID/cVuZ49auytthRufF",
	"IfbIGGQkaKa3woPyTATmUKMBy6j3rfMOvjc4hwzx1Z8QKaS2RCEqkQCVCwYxogyRJEERkSARX6M1oUku",
	"QN7gAGeCZyAUtSKnICXZmN0FkPiZJbtSmW3bFE/sqXgf4KGUoEK25sb5jjZOuNWWw8qlKR0L9sEZhepD",
	"l5rQj2lZ7AMsT/Hgdzz8skDj0a8oZFIRFmlnGP4vF1B/8Dia1f/e8+gVRP3J+JsCwUhSfzbiTBHKQNR/",
	"ozDVOv0jaAs4/pYlnKq2vqI3CO+dOjmyeh9lSp6LCO4/uDVNVeJmy0ViEFEFqXSfmCcJWWn2I08hQpCd",
	"2yiF2A+UxZRtwjQjkUMHZL2GSEFsTChHPLfvnscxKVOwAYFN/Dlo9ZTnlMp3QiywLQWwuP2yzSETIPVu",
	"SG0BKa5IglierkCYF8wyS0QUIkhmENE1jVARdxqWLuVqy6GKuNcx7J6UQbaFmFCpNNqTEmQgDG4kE67Q",
	"mgtDfhCpoNMvWDua1BbP2eKhRqpFOUA+uF0XbmMsE8adLnLCIx+OoZaBYjYcfRo+jnGAP79Mnsbz4Ydw",
	"Ei5/wwGeDidfhnO9shiP5uOlfhQuRs9PD+Hjy3y4DJ+fcIDnz8/LT6FeHP93NnkOl84oUBxeufixnaxt",
	"jJ9o0wCJtqXekdmrqffC/6Xbq1KSvBMBnkUqI87WdJMLE689ewjO1av3BAmRAN/iW54wEGRFE1ribRKd",
	"MJD0BYu6zMfqW/IM/QcV65VjSy4UxGi1Q9RsCTEiJtBYTeOgm+s5Q9lZFzyyggtusXxxuFO7b3+4Lr9w",
	"Am8QXl6CxgG9RclI9Eo24JWgWL848Jndtzfe+rvmwlusXxzv3O7bG2/t7XfBtcsXR7sw2/YG64hGLtB1",
	"st3FsX+u795ThFOx8lziPyQRQ2eSu+4T6rnF9Ai9c7DsovppFQEbTYhdePIVssV6l7JiWiM1r77attUx",
	"I2pb1kFrmoBtoCJbvssyFGNH4hZ54oPpsosz+F6u7K2llA46OQmx1G1L4mb0dVivaiHb+oIUYupvBGVE",
	"GIN4VpjJsy68niHhDQRVu745ZFHyaZWAVCOiYMPFzt0qgVT3Z5owTePs35w6P5nRLugfDtv10VI39Iua",
	"DcoyuknzkW62B7r2FlOIaZ6eIJjw98Oqq6AuUm1bd97mOMtF4lx4AyHdVnYpw5njL2fBrJKrQ6XhhjiH",
	"TeVjzoSnu41DijNJAQnD5GvwKhE6JIiC2EQDvWmP+Dmn8pWCVFZv/XsAUfCXSbpK3yVnzwqJytedAXOB",
	"it8PruwFromta3l/AmVzi2viPVsTe2GWnNdEd6YC9oMrGK+JrWPB68fY3OBHatw+kE8FAhvLHJFAVAvu",
	"0rcgQO9UbYvCrwh4pvh7Jzry5SxGnOnl9AbN6xyMVwzvNEkQ4wqtAAnIjKI6V82NaPzD2ii06Ynmjrme",
	"ievMmrcV10l9kH92+G4I94F/kukEbd9Dh+nsgrfGK9a7VP/zGukpENdK16KSsQPMkxCbg8npePo8/w0H",
	"+NN4/jSe4AAPZ7NJOCoHjw/hfGrmk67yyPbKjvzJ4hFP8pS5J3fA4gllnsGhbpxmzvZKW/KovSo6K9Ni",
	"bqEIeq72ak3ZBkQmqGsq+sQVDJDaUomoNO9fzujX3NmnmUvCU6IZAp9wLrO4xg2Xcxx5MND5kYcb3+fj",
	"KN0CevW5xDGEneuGS8ofQzLSnGfvnbp3g0eb11vBo6mPv071KMJtDIve0TUrQaP+epgWfKZTiZS9qP6L",
	"TYz3kBbqFZGwiPjRZYLNNbVruFKxXjo7O/Otn0V4rZfwrem/nQ3TAXSfnO0ZPl4jgwuqaESSRvQY+a8o",
	"t3Sz7U6d8PfuxKmZAnSnZ7BJ6IauEujKc9ZKrlnGaB4uw9FQp9yP4eNHHODp+D58meIAT56/4AA/jR8n",
	"4WP4YeJKvvpMWpiluHPHn6ejhOhj0EuIhrNQ19SHNxb/cnN3c6eR8QwYySge4H/f3N38gu0401j7NiZy",
	"u+JExLfr1j3ZxvqY9g7TmIUxHuBHUPclT+NqrfEVzK93dxf7+KVxkuP7l0UeRWDDewxrkifeLHgAeXv0",
	"nY75ZCZPUyJ2VkxEUHI875bFtP5wm33Q3o1hd2izmqR31mbBEhx9hfW7W5aK5Lb6nmkfnCUuv9na//ET",
	"jFZO9v8eo525pGjYTbQmRWft1hguXVGhjZN+tkKbrf35t0C02+3O6ix5foI+y6P+NoWWQwWnRvf7/wcA",
	"AP//0pzDgZ8pAAA=",
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
