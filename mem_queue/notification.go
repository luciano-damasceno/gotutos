package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var claimIDs []int

type NotificationsMeli struct {
	Resource int    `json:"resource"`
	Label    string `json:"label"`
}

func Notification(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var notification NotificationsMeli
	json.Unmarshal(reqBody, &notification)
	for {
		flag := false
		i := findIndexByValue(claimIDs, notification.Resource)
		if i != -1 {
			flag = true
		}
		if !flag {
			claimIDs = append(claimIDs, notification.Resource)
			j := findIndexByValue(claimIDs, notification.Resource)
			if notification.Resource == 1000 {
				time.Sleep(5 * time.Second)
			}
			log.Println(notification.Label)
			claimIDs = append(claimIDs[:j], claimIDs[j+1:]...)
			break
		}
	}
}

func findIndexByValue(list []int, value int) int {
	for i, element := range list {
		if element == value {
			return i
		}
	}
	return -1
}
