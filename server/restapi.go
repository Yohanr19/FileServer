package main

import (
	"net/http"
	"encoding/json"
)

type StatusResponse struct {
	Channel string `json:"channel"`
	Amount int `json:"amount"`
	SubscriberAddrs []string `json:"subscriber_addrs"`
}

func httpHandleStatus(w http.ResponseWriter, r *http.Request){
	connData := ConnMap.CopyMap()
	var response []StatusResponse
	for ch , conns := range connData {
		var item StatusResponse
		item.Channel = ch
		item.Amount = len(conns)
		for _ , conn := range conns{
			addr := conn.RemoteAddr().String()
			item.SubscriberAddrs = append(item.SubscriberAddrs, addr)
		}
		response = append(response, item)
	}
	enconder := json.NewEncoder(w)
	enconder.Encode(response)
}

