package http

import (
	"encoding/json"
	"github.com/AnyKey/service/user"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	usecase user.Usecase
}

func Launch(router *mux.Router, uc user.Usecase) error {
	userHandler := &handler{usecase: uc}
	router.HandleFunc("/subscriptions_list", userHandler.GetList).Methods(http.MethodGet, http.MethodOptions)
	return nil
}
func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res, err := h.usecase.GetList()
	if err != nil {
		log.Errorln("[GetToken] Error: ", err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}
	bytes, _ := json.Marshal(res)
	w.Write(bytes)
	return
}
