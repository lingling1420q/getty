package micro_examples

import (
	jerrors "github.com/juju/errors"
	// "errors"
)

type TestService struct {
	i int
}

func (r *TestService) Service() string {
	return "TestService"
}

func (r *TestService) Version() string {
	return "v1.0"
}

func (r *TestService) Test(req *TestReq, rsp *TestRsp) error {
	rsp.A = req.A + ", " + req.B + ", " + req.C
	return nil
}

func (r *TestService) Add(req *AddReq, rsp *AddRsp) error {
	rsp.Sum = req.A + req.B
	return nil
}

func (r *TestService) Err(req *ErrReq, rsp *ErrRsp) error {
	return jerrors.New("this is a error test")
}
