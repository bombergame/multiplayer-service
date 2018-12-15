package rest

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func NewNumRooms() prometheus.Gauge {
	return promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "multiplayer_service_rooms_number",
		},
	)
}
