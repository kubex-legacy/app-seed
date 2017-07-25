package main

import (
	"bytes"
	"html/template"
	"github.com/kubex/potens-go/webui"
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
)

func PageDefinition(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error) {
	response := webui.CreateResponse()
	webui.SetPageTitle(response, "Page Title")
	return response, nil
}

func DefaultPage(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error) {
	response := webui.CreateResponse()
	buf := new(bytes.Buffer)

	defaultTemplate := template.Must(template.ParseFiles("templates/default.html"))
	defaultTemplate.Execute(buf, "")
	response.Body = buf.String()

	return response, nil
}

func Panel(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error) {
	response := webui.CreateResponse()
	response.Body = "Default Response"
	return response, nil
}
