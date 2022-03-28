package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockBalansController struct{}

func (mbc mockBalansController) withdraw(userId, amount int) error {
	if userId == 10 {
		return errors.New("error")
	}
	return nil
}
func (mbc mockBalansController) deposit(userId, amount int) error {
	if userId == 10 {
		return errors.New("error")
	}
	return nil
}
func (mbc mockBalansController) transfer(senderId, receiverId, amount int) error {
	if senderId == 10 {
		return errors.New("error")
	}
	return nil
}
func Test_handlerContext_withdrawHandler(t *testing.T) {
	correctReader := strings.NewReader("{\"userId\":1, \"amount\":1}")
	incorrectReader := strings.NewReader("incorrect")
	incorrectForBCReader := strings.NewReader("{\"userId\":10, \"amount\":1}")
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name          string
		hctx          handlerContext
		args          args
		resStatusCode int
	}{
		{
			name: "corrext test",
			hctx: handlerContext{bc: mockBalansController{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/withdraw", correctReader),
			},
			resStatusCode: http.StatusOK,
		},
		{
			name: "incorrect request body",
			hctx: handlerContext{bc: mockBalansController{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/withdraw", incorrectReader),
			},
			resStatusCode: http.StatusBadRequest,
		},
		{
			name: "error from Balans Controller",
			hctx: handlerContext{bc: mockBalansController{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/withdraw", incorrectForBCReader),
			},
			resStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hctx.withdrawHandler(tt.args.w, tt.args.r)
			res := tt.args.w.Result()
			if res.StatusCode != tt.resStatusCode {
				t.Errorf("withdrawHandler() statusCode = %v, want %v", res.StatusCode, tt.resStatusCode)
			}
		})
	}
}

func Test_handlerContext_depositHandler(t *testing.T) {
	correctReader := strings.NewReader("{\"userId\":1, \"amount\":1}")
	incorrectReader := strings.NewReader("incorrect")
	incorrectForBCReader := strings.NewReader("{\"userId\":10, \"amount\":1}")
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name          string
		hctx          handlerContext
		args          args
		resStatusCode int
	}{
		{
			name: "corrext test",
			hctx: handlerContext{bc: mockBalansController{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/deposit", correctReader),
			},
			resStatusCode: http.StatusOK,
		},
		{
			name: "incorrect request body",
			hctx: handlerContext{bc: mockBalansController{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/deposit", incorrectReader),
			},
			resStatusCode: http.StatusBadRequest,
		},
		{
			name: "error from Balans Controller",
			hctx: handlerContext{bc: mockBalansController{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/deposit", incorrectForBCReader),
			},
			resStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hctx.depositHandler(tt.args.w, tt.args.r)
			res := tt.args.w.Result()
			if res.StatusCode != tt.resStatusCode {
				t.Errorf("depositHandler() statusCode = %v, want %v", res.StatusCode, tt.resStatusCode)
			}
		})
	}
}

func Test_handlerContext_transferHandler(t *testing.T) {
	correctReader := strings.NewReader("{\"senderId\":1, \"recieverId\":1, \"amount\":1}")
	incorrectReader := strings.NewReader("incorrect")
	incorrectForBCReader := strings.NewReader("{\"senderId\":10, \"recieverId\":1, \"amount\":1}")
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name          string
		hctx          handlerContext
		args          args
		resStatusCode int
	}{
		{
			name: "corrext test",
			hctx: handlerContext{bc: mockBalansController{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/transfer", correctReader),
			},
			resStatusCode: http.StatusOK,
		},
		{
			name: "incorrect request body",
			hctx: handlerContext{bc: mockBalansController{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/transfer", incorrectReader),
			},
			resStatusCode: http.StatusBadRequest,
		},
		{
			name: "error from Balans Controller",
			hctx: handlerContext{bc: mockBalansController{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/transfer", incorrectForBCReader),
			},
			resStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hctx.transferHandler(tt.args.w, tt.args.r)
			res := tt.args.w.Result()
			if res.StatusCode != tt.resStatusCode {
				t.Errorf("transferHandler() statusCode = %v, want %v", res.StatusCode, tt.resStatusCode)
			}
		})
	}
}
