package main

import "strings"

type Record struct {
	Domain        string `json:"NTDom"`
	User          string `json:"User"`
	NTHash        string `json:"Hash"`
	Data          string `json:"Data"`
	IsCompromised bool   `json:"IsCompromised"`
}

func NewRecord(l string) (r Record) {

	line := strings.Split(l, ":")

	r.Domain = line[0][:strings.Index(l, "\\")]
	r.User = line[0][strings.Index(l, "\\")+1:]
	r.NTHash = line[3]
	r.Data = line[4]

	return
}
