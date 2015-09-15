package main

import (
	"encoding/json"
	"time"
)

// Status represents the status of the running piladb
// instance.
type Status struct {
	Code       string    `json:"status"`
	Version    string    `json:"version"`
	StartedAt  time.Time `json:"started_at"`
	RunningFor float64   `json:"running_for"`
}

// NewStatus returns a new piladb status.
func NewStatus(now time.Time) *Status {
	status := &Status{}
	status.Code = "OK"
	status.Version = "0"
	status.StartedAt = now
	return status
}

// RunningFor returns the time piladb was running in
// seconds given the current time.
func (s *Status) SetRunningFor(now time.Time) float64 {
	s.RunningFor = now.Sub(s.StartedAt).Seconds()
	return s.RunningFor
}

// ToJson returns the Status into a JSON file in []byte
// format.
func (s *Status) ToJson(now time.Time) []byte {
	_ = s.SetRunningFor(now)

	// Do not check error as the Status type does
	// not contain types that could cause such case.
	// See http://golang.org/src/encoding/json/encode.go?s=5438:5481#L125
	b, _ := json.Marshal(s)
	return b
}
