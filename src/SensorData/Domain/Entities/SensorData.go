package entities



type SensorData struct {
	ID             int64    `json:"id"`
	DireccionMac string `json:"direccion_mac"`
	Ubicacion    string `json:"ubicacion"`
}
