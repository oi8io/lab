package transports

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"missrv/endpoints"
	"net/http"
	"strconv"
)

//DecodeUserRequest 初始化参数 编码
func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if uid, ok := vars["uid"]; ok && uid != "" {
		atoi, _ := strconv.Atoi(uid)
		return endpoints.UserRequest{Uid: int64(atoi)}, nil
	}
	return nil, errors.New("param error")
}

//EncodeUserResponse
func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Context-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
