// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Id string `path:"id"`
}

type Response struct {
	Id     string `json:"message"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
