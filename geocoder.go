// Package geo is a generic framework to develop geocode/reverse geocode clients
package geo

import (
	"context"
)

// Geocoder can look up (lat, long) by address and address by (lat, long)
type Geocoder interface {
	Geocode(ctx context.Context, address string) (*Location, error)
	ReverseGeocode(ctx context.Context, lat, lng float64) (*Address, error)
}

// Location is the output of Geocode
type Location struct {
	Lat, Lng float64
}

// Address is returned by ReverseGeocode.
// This is a structured representation of an address, including its flat representation
type Address struct {
	FormattedAddress string
	Street           string
	HouseNumber      string
	Suburb           string
	Postcode         string
	State            string
	StateCode        string
	StateDistrict    string
	County           string
	Country          string
	CountryCode      string
	City             string
}

// Log is used to log debug and error messages.
var Log Logger = NullLogger{}

// Logger is used to log debug and error messages.
type Logger interface {
	Debug(context.Context, string, ...any)
	Info(context.Context, string, ...any)
	Error(context.Context, string, ...any)
}

// NullLogger is a logger that does nothing.
type NullLogger struct{}

func (NullLogger) Debug(context.Context, string, ...any) {}
func (NullLogger) Info(context.Context, string, ...any)  {}
func (NullLogger) Error(context.Context, string, ...any) {}
