package api

import (
	. "andrew.com/notifications/cmd/app/accessors"
	. "andrew.com/notifications/cmd/app/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"html/template"
	"log"
)

var (
	ContentType     = []byte("Content-Type")
	ApplicationJSON = []byte("application/json")
)

type Api struct {
	router        *router.Router
	emailAccessor EmailAccessor
}

func CreateRouter(emailAcc *EmailAccessor) *Api {
	r := router.New()
	route := Api{
		router:        r,
		emailAccessor: *emailAcc,
	}
	r.GET("/notifications", route.getAllHandler)
	api := r.Group("/notifications/{id}")
	api.GET("/", route.getHandler)
	api.PUT("/", route.upsertHandler)
	return &route
}

func (api *Api) GetHandler() func(ctx *fasthttp.RequestCtx) {
	return api.router.Handler
}

func createError(err error, ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
	errMsg := fmt.Sprintf("%s", err.Error())
	ctx.Response.AppendBodyString(errMsg)
}

func (api *Api) upsertHandler(ctx *fasthttp.RequestCtx) {
	nType := ctx.UserValue("type")
	switch nType {
	case "email":
		record := DbENRecord{}
		if err := json.Unmarshal(ctx.PostBody(), &record); err != nil {
			createError(err, ctx)
		}
		_, err := template.ParseGlob(record.Template)
		if err != nil {
			log.Printf("Wrong template format \n%s \n", err)
			createError(errors.New("incorrect template"), ctx)
		}
		notification := EmailNotification{}
		record.SetConfigToNotification(&notification)
		_, err = api.emailAccessor.InsertNotification(notification)
		if err != nil {
			createError(err, ctx)
		}
	default:
		createError(errors.New("unsupported notification type"), ctx)
	}

	ctx.Response.Header.SetCanonical(ContentType, ApplicationJSON)
	ctx.Response.SetStatusCode(204)
}

func (api *Api) getHandler(ctx *fasthttp.RequestCtx) {
	value := ctx.UserValue("id")
	id, ok := value.(int)
	if ok == false {
		createError(errors.New("empty or wrong format id"), ctx)
	}

	notification, err := api.emailAccessor.GetNotification(id)
	if err != nil {
		createError(err, ctx)
	} else {
		record := notification.GetDbNotification()
		ctx.Response.Header.SetCanonical(ContentType, ApplicationJSON)
		ctx.Response.SetStatusCode(200)
		if err := json.NewEncoder(ctx).Encode(record); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
	}
}

func (api *Api) getAllHandler(ctx *fasthttp.RequestCtx) {
	records, err := api.emailAccessor.GetNotifications()
	if err != nil {
		createError(err, ctx)
	} else {
		ctx.Response.Header.SetCanonical(ContentType, ApplicationJSON)
		ctx.Response.SetStatusCode(200)
		if err := json.NewEncoder(ctx).Encode(records); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
	}
}
