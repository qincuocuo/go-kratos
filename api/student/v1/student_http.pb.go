// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.21.4
// source: api/student/v1/student.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationStudentGetStudent = "/api.student.v1.Student/GetStudent"

type StudentHTTPServer interface {
	GetStudent(context.Context, *GetStudentRequest) (*GetStudentReply, error)
}

func RegisterStudentHTTPServer(s *http.Server, srv StudentHTTPServer) {
	r := s.Route("/")
	r.GET("/student/{id}", _Student_GetStudent0_HTTP_Handler(srv))
}

func _Student_GetStudent0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetStudentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentGetStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetStudent(ctx, req.(*GetStudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetStudentReply)
		return ctx.Result(200, reply)
	}
}

type StudentHTTPClient interface {
	GetStudent(ctx context.Context, req *GetStudentRequest, opts ...http.CallOption) (rsp *GetStudentReply, err error)
}

type StudentHTTPClientImpl struct {
	cc *http.Client
}

func NewStudentHTTPClient(client *http.Client) StudentHTTPClient {
	return &StudentHTTPClientImpl{client}
}

func (c *StudentHTTPClientImpl) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...http.CallOption) (*GetStudentReply, error) {
	var out GetStudentReply
	pattern := "/student/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStudentGetStudent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
