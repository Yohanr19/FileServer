package controlers

import (
	"encoding/json"
	"github.com/yohanr19/fileserver/server/pkg/models"
	"net/http"
	"time"
)

func NewReportControler() (*ReportControler, error) {
	var rc ReportControler
	err := rc.store.Init()
	if err != nil {
		return nil, err
	}
	return &rc, nil
}

type ReportControler struct {
	store models.ReportStore
}
type ReportData struct {
	ID               uint   `json:"id"`
	Date             string `json:"date"`
	Filename         string `json:"filename"`
	Status           string `json:"status"`
	Filesize         int    `json:"filesize"`
	Channel          string `json:"channel"`
	SenderAdd        string `json:"sender_add"`
	SubscriberAmount int    `json:"subscriber_amount"`
}

func (rc *ReportControler) GetReports(w http.ResponseWriter, r *http.Request) {
	var responseData []ReportData

	data, err := rc.store.GetAll()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	for _, report := range data {
		var rep ReportData
		rep.ID = report.ID
		rep.Date = report.CreatedAt.Format(time.Layout)
		rep.Filename = report.Filename
		rep.Status = report.Status
		rep.Filesize = report.Filesize
		rep.Channel = report.Channel
		rep.SenderAdd = report.SenderAdd
		rep.SubscriberAmount = report.SubscriberAmount
		responseData = append(responseData, rep)
	}
	decoder := json.NewEncoder(w)
	err = decoder.Encode(&responseData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (rc *ReportControler) AddReport(r ReportData) error {
	var report models.Report
	report.Filename = r.Filename
	report.Status = r.Status
	report.Filesize = r.Filesize
	report.Channel = r.Channel
	report.SenderAdd = r.SenderAdd
	report.SubscriberAmount = r.SubscriberAmount
	err := rc.store.Create(report)
	if err != nil {
		return err
	}
	return nil
}
