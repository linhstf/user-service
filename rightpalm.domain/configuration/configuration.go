package configuration

import "time"

type Configuration struct {
	ID          int
	Title       string
	Key         string
	Value       string
	Description string
	Group       ConfigurationGroup
	Order       int
	UpdatedDate time.Time
	UseFunction string
	SetFunction string
}
