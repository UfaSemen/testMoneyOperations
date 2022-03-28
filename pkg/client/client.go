package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/UfaSemen/testMoneyOperations/pkg/data"
)

func parseWithdrawOrDeposit(operation []string) (int, int, error) {
	if len(operation) != 3 {
		return 0, 0, errors.New("not enough values")
	}
	userId, err := strconv.Atoi(operation[1])
	if err != nil {
		return 0, 0, err
	}
	amount, err := strconv.Atoi(operation[2])
	if err != nil {
		return 0, 0, err
	}
	return userId, amount, nil
}

func parseTransfer(operation []string) (int, int, int, error) {
	if len(operation) != 4 {
		return 0, 0, 0, errors.New("not enogh values")
	}
	senderId, err := strconv.Atoi(operation[1])
	if err != nil {
		return 0, 0, 0, err
	}
	receiverId, err := strconv.Atoi(operation[2])
	if err != nil {
		return 0, 0, 0, err
	}
	amount, err := strconv.Atoi(operation[3])
	if err != nil {
		return 0, 0, 0, err
	}
	return senderId, receiverId, amount, nil
}

func makeRequest(url string, jsonBody []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("response Body:", string(body))
	return nil
}

func ExecuteClient(operation, serverAddress string) error {
	if operation == "" {
		return errors.New("no operation passed")
	}
	opSl := strings.Fields(operation)

	switch opSl[0] {
	case "withdraw":
		userId, amount, err := parseWithdrawOrDeposit(opSl)
		if err != nil {
			return err
		}
		reqStruct := data.WithdrawRequest{
			UserId: userId,
			Amount: amount,
		}
		reqStr, err := json.Marshal(reqStruct)
		if err != nil {
			return err
		}
		err = makeRequest(serverAddress+"/withdraw", reqStr)
		if err != nil {
			return err
		}
	case "deposit":
		userId, amount, err := parseWithdrawOrDeposit(opSl)
		if err != nil {
			return err
		}
		reqStruct := data.DepositRequest{
			UserId: userId,
			Amount: amount,
		}
		reqStr, err := json.Marshal(reqStruct)
		if err != nil {
			return err
		}
		err = makeRequest(serverAddress+"/deposit", reqStr)
		if err != nil {
			return err
		}
	case "transfer":
		senderId, receiverId, amount, err := parseTransfer(opSl)
		if err != nil {
			return err
		}
		reqStruct := data.TransferRequest{
			SenderId:   senderId,
			RecieverId: receiverId,
			Amount:     amount,
		}
		reqStr, err := json.Marshal(reqStruct)
		if err != nil {
			return err
		}
		err = makeRequest(serverAddress+"/transfer", reqStr)
		if err != nil {
			return err
		}
	default:
		return errors.New("no operation passed")
	}
	return nil
}
