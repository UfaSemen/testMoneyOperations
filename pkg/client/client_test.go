package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_parseWithdrawOrDeposit(t *testing.T) {
	type args struct {
		operation []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "successfull test",
			args: args{
				operation: []string{"deposit", "1", "2"},
			},
			want:    1,
			want1:   2,
			wantErr: false,
		},
		{
			name: "wrong slice length test",
			args: args{
				operation: []string{"deposit", "1"},
			},
			wantErr: true,
		},
		{
			name: "wrong first element test",
			args: args{
				operation: []string{"deposit", "w", "2"},
			},
			wantErr: true,
		},
		{
			name: "wrong second element test",
			args: args{
				operation: []string{"deposit", "1", "w"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseWithdrawOrDeposit(tt.args.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseWithdrawOrDeposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseWithdrawOrDeposit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseWithdrawOrDeposit() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseTransfer(t *testing.T) {
	type args struct {
		operation []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   int
		wantErr bool
	}{
		{
			name: "successfull test",
			args: args{
				operation: []string{"transfer", "1", "2", "3"},
			},
			want:    1,
			want1:   2,
			want2:   3,
			wantErr: false,
		},
		{
			name: "wrong slice length test",
			args: args{
				operation: []string{"transfer", "1", "2"},
			},
			wantErr: true,
		},
		{
			name: "wrong first element test",
			args: args{
				operation: []string{"transfer", "w", "2", "3"},
			},
			wantErr: true,
		},
		{
			name: "wrong second element test",
			args: args{
				operation: []string{"transfer", "1", "w", "3"},
			},
			wantErr: true,
		},
		{
			name: "wrong third element test",
			args: args{
				operation: []string{"transfer", "1", "2", "w"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := parseTransfer(tt.args.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseTransfer() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseTransfer() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseTransfer() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_makeRequest(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "response")
	}))
	defer svr.Close()
	type args struct {
		url      string
		jsonBody []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successfull test",
			args: args{
				url:      svr.URL,
				jsonBody: []byte{},
			},
			wantErr: false,
		},
		{
			name: "wrong url test",
			args: args{
				url:      "wrong",
				jsonBody: []byte{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := makeRequest(tt.args.url, tt.args.jsonBody); (err != nil) != tt.wantErr {
				t.Errorf("makeRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExecuteClient(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "response")
	}))
	defer svr.Close()
	type args struct {
		operation     string
		serverAddress string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successfull deposit test",
			args: args{
				operation:     "deposit 2 3",
				serverAddress: svr.URL,
			},
			wantErr: false,
		},
		{
			name: "successfull withdraw test",
			args: args{
				operation:     "withdraw 2 3",
				serverAddress: svr.URL,
			},
			wantErr: false,
		},
		{
			name: "successfull transfer test",
			args: args{
				operation:     "transfer 2 3 4",
				serverAddress: svr.URL,
			},
			wantErr: false,
		},
		{
			name: "wrong url deposit test",
			args: args{
				operation:     "deposit 2 3",
				serverAddress: "wrong",
			},
			wantErr: true,
		},
		{
			name: "wrong url withdraw test",
			args: args{
				operation:     "withdraw 2 3",
				serverAddress: "wrong",
			},
			wantErr: true,
		},
		{
			name: "wrong url transfer test",
			args: args{
				operation:     "transfer 2 3 4",
				serverAddress: "wrong",
			},
			wantErr: true,
		},
		{
			name: "empty operation test",
			args: args{
				operation:     "",
				serverAddress: svr.URL,
			},
			wantErr: true,
		},
		{
			name: "wrong operation test",
			args: args{
				operation:     "wrong 1 2",
				serverAddress: svr.URL,
			},
			wantErr: true,
		},
		{
			name: "wrong withdraw test",
			args: args{
				operation:     "withdraw wrong",
				serverAddress: svr.URL,
			},
			wantErr: true,
		},
		{
			name: "wrong deposit test",
			args: args{
				operation:     "deposit wrong",
				serverAddress: svr.URL,
			},
			wantErr: true,
		},
		{
			name: "wrong transfer test",
			args: args{
				operation:     "transfer wrong",
				serverAddress: svr.URL,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExecuteClient(tt.args.operation, tt.args.serverAddress); (err != nil) != tt.wantErr {
				t.Errorf("ExecuteClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
