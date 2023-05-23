package handler

import "github.com/csguojin/reserve/service"

type HandlerStruct struct {
	svc service.Service
}

func NewHandler(s service.Service) *HandlerStruct {
	return &HandlerStruct{
		svc: s,
	}
}
