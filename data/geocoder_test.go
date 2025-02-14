package data_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mariotoffia/geo-golang"
	"github.com/mariotoffia/geo-golang/data"
	"github.com/stretchr/testify/assert"
)

var (
	addressFixture = geo.Address{
		FormattedAddress: "64 Elizabeth Street, Melbourne, Victoria 3000, Australia",
	}
	locationFixture = geo.Location{
		Lat: -37.814107,
		Lng: 144.96328,
	}
	geocoder = data.Geocoder(
		data.AddressToLocation{
			addressFixture: locationFixture,
		},
		data.LocationToAddress{
			locationFixture: addressFixture,
		},
	)
)

func TestGeocode(t *testing.T) {
	location, err := geocoder.Geocode(context.TODO(), addressFixture.FormattedAddress)
	assert.NoError(t, err)
	assert.Equal(t, geo.Location{Lat: -37.814107, Lng: 144.96328}, *location)
}

func TestReverseGeocode(t *testing.T) {
	address, err := geocoder.ReverseGeocode(context.TODO(), locationFixture.Lat, locationFixture.Lng)
	assert.Nil(t, err)
	assert.NotNil(t, address)
	assert.True(t, strings.Contains(address.FormattedAddress, "Melbourne, Victoria 3000, Australia"))
}

func TestReverseGeocodeWithNoResult(t *testing.T) {
	addr, err := geocoder.ReverseGeocode(context.TODO(), 1, 2)
	assert.Nil(t, err)
	assert.Nil(t, addr)
}
