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
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	

	ConnMap.ClearConnections()
	connData := ConnMap.CopyMap()
	var response []StatusResponse
	for ch , conns := range connData {
		var item StatusResponse
		item.Channel = ch
		
		for _ , conn := range conns{
			if conn == nil{
				continue
			}
			addr := conn.RemoteAddr().String()
			item.SubscriberAddrs = append(item.SubscriberAddrs, addr)
		}
		item.Amount = len(item.SubscriberAddrs)
		response = append(response, item)
	}
	enconder := json.NewEncoder(w)
	enconder.Encode(response)
}

