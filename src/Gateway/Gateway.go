package Gateway

type Gateway struct {
}

func (*Gateway) ShowConfig() string {
	return "GateWay"
}

func GetNewGate() *Gateway {
	return new(Gateway)
}
