package models

type LogrusMessage struct {
	Version  string  `json:"version,omitempty"`
	Host     string  `json:"host,omitempty"`
	Short    string  `json:"short_message,omitempty"`
	Full     string  `json:"full_message,omitempty"`
	TimeUnix float64 `json:"timestamp,omitempty"`
	Level    int32   `json:"level,omitempty"`
	Facility string  `json:"facility,omitempty"`
	File     string  `json:"file,omitempty"`
	Line     int     `json:"line,omitempty"`
}
