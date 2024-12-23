package models

type Room struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	HostID int    `json:"host_id"`
}
