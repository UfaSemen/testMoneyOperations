package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/UfaSemen/testMoneyOperations/pkg/data"
)

type balansController interface {
	withdraw(userId, amount int) error
	deposit(userId, amount int) error
	transfer(senderId, receiverId, amount int) error
}
type handlerContext struct {
	bc balansController
}

func (hctx handlerContext) withdrawHandler(w http.ResponseWriter, r *http.Request) {
	var req data.WithdrawRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = hctx.bc.withdraw(req.UserId, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("successfull operation"))
}

func (hctx handlerContext) depositHandler(w http.ResponseWriter, r *http.Request) {
	var req data.DepositRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = hctx.bc.deposit(req.UserId, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("successfull operation"))
}

func (hctx handlerContext) transferHandler(w http.ResponseWriter, r *http.Request) {
	var req data.TransferRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = hctx.bc.transfer(req.SenderId, req.RecieverId, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("successfull operation"))
}

func StartServer(port int, bc balansController) {

	hctx := handlerContext{bc: bc}
	http.HandleFunc("/withdraw", hctx.withdrawHandler)
	http.HandleFunc("/deposit", hctx.depositHandler)
	http.HandleFunc("/transfer", hctx.transferHandler)

	err := http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
