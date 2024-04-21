// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get image
	// (GET /internal/v1/images/{image_id})
	GetImage(ctx echo.Context, imageId string) error
	// Post plat.
	// (POST /internal/v1/plats)
	PostPlat(ctx echo.Context) error
	// Delete plat by ID.
	// (DELETE /internal/v1/plats/{plat_id})
	DeletePlat(ctx echo.Context, platId PlatId) error
	// Get plat by ID.
	// (GET /internal/v1/plats/{plat_id})
	GetPlat(ctx echo.Context, platId PlatId) error
	// Delete "favorite" to a specific plat by ID.
	// (DELETE /internal/v1/plats/{plat_id}/favorites)
	DeleteFavorite(ctx echo.Context, platId PlatId) error
	// Add "favorite" to a specific plat by ID.
	// (POST /internal/v1/plats/{plat_id}/favorites)
	PostFavorite(ctx echo.Context, platId PlatId) error
	// get timeline by query params
	// (GET /internal/v1/timelines)
	GetTimelineByQuery(ctx echo.Context, params GetTimelineByQueryParams) error
	// get timeline
	// (GET /internal/v1/timelines/{timeline_id})
	GetTimeline(ctx echo.Context, timelineId TimelineId) error
	// Get users.
	// (GET /internal/v1/users)
	GetUsers(ctx echo.Context, params GetUsersParams) error
	// Create new user.
	// (POST /internal/v1/users)
	CreateUser(ctx echo.Context) error
	// Delete User by ID.
	// (DELETE /internal/v1/users/{user_id})
	DeleteUser(ctx echo.Context, userId UserId) error
	// get user
	// (GET /internal/v1/users/{user_id})
	GetUser(ctx echo.Context, userId UserId) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetImage converts echo context to params.
func (w *ServerInterfaceWrapper) GetImage(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "image_id" -------------
	var imageId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "image_id", runtime.ParamLocationPath, ctx.Param("image_id"), &imageId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter image_id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetImage(ctx, imageId)
	return err
}

// PostPlat converts echo context to params.
func (w *ServerInterfaceWrapper) PostPlat(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostPlat(ctx)
	return err
}

// DeletePlat converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePlat(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "plat_id" -------------
	var platId PlatId

	err = runtime.BindStyledParameterWithLocation("simple", false, "plat_id", runtime.ParamLocationPath, ctx.Param("plat_id"), &platId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter plat_id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePlat(ctx, platId)
	return err
}

// GetPlat converts echo context to params.
func (w *ServerInterfaceWrapper) GetPlat(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "plat_id" -------------
	var platId PlatId

	err = runtime.BindStyledParameterWithLocation("simple", false, "plat_id", runtime.ParamLocationPath, ctx.Param("plat_id"), &platId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter plat_id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPlat(ctx, platId)
	return err
}

// DeleteFavorite converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteFavorite(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "plat_id" -------------
	var platId PlatId

	err = runtime.BindStyledParameterWithLocation("simple", false, "plat_id", runtime.ParamLocationPath, ctx.Param("plat_id"), &platId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter plat_id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteFavorite(ctx, platId)
	return err
}

// PostFavorite converts echo context to params.
func (w *ServerInterfaceWrapper) PostFavorite(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "plat_id" -------------
	var platId PlatId

	err = runtime.BindStyledParameterWithLocation("simple", false, "plat_id", runtime.ParamLocationPath, ctx.Param("plat_id"), &platId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter plat_id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostFavorite(ctx, platId)
	return err
}

// GetTimelineByQuery converts echo context to params.
func (w *ServerInterfaceWrapper) GetTimelineByQuery(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTimelineByQueryParams
	// ------------- Required query parameter "user_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "user_id", ctx.QueryParams(), &params.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// ------------- Required query parameter "type" -------------

	err = runtime.BindQueryParameter("form", true, true, "type", ctx.QueryParams(), &params.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter type: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "to_date" -------------

	err = runtime.BindQueryParameter("form", true, false, "to_date", ctx.QueryParams(), &params.ToDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter to_date: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTimelineByQuery(ctx, params)
	return err
}

// GetTimeline converts echo context to params.
func (w *ServerInterfaceWrapper) GetTimeline(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "timeline_id" -------------
	var timelineId TimelineId

	err = runtime.BindStyledParameterWithLocation("simple", false, "timeline_id", runtime.ParamLocationPath, ctx.Param("timeline_id"), &timelineId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeline_id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTimeline(ctx, timelineId)
	return err
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersParams
	// ------------- Optional query parameter "user_name" -------------

	err = runtime.BindQueryParameter("form", true, false, "user_name", ctx.QueryParams(), &params.UserName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_name: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx, params)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// DeleteUser converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, ctx.Param("user_id"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUser(ctx, userId)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "user_id" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "user_id", runtime.ParamLocationPath, ctx.Param("user_id"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUser(ctx, userId)
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

	router.GET(baseURL+"/internal/v1/images/:image_id", wrapper.GetImage)
	router.POST(baseURL+"/internal/v1/plats", wrapper.PostPlat)
	router.DELETE(baseURL+"/internal/v1/plats/:plat_id", wrapper.DeletePlat)
	router.GET(baseURL+"/internal/v1/plats/:plat_id", wrapper.GetPlat)
	router.DELETE(baseURL+"/internal/v1/plats/:plat_id/favorites", wrapper.DeleteFavorite)
	router.POST(baseURL+"/internal/v1/plats/:plat_id/favorites", wrapper.PostFavorite)
	router.GET(baseURL+"/internal/v1/timelines", wrapper.GetTimelineByQuery)
	router.GET(baseURL+"/internal/v1/timelines/:timeline_id", wrapper.GetTimeline)
	router.GET(baseURL+"/internal/v1/users", wrapper.GetUsers)
	router.POST(baseURL+"/internal/v1/users", wrapper.CreateUser)
	router.DELETE(baseURL+"/internal/v1/users/:user_id", wrapper.DeleteUser)
	router.GET(baseURL+"/internal/v1/users/:user_id", wrapper.GetUser)

}
