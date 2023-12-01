package eventbus

import "github.com/asaskevich/EventBus"

var (
	ipPort     = "0.0.0.0:4533"
	busPath    = "/hassbus"
	networkBus = EventBus.NewNetworkBus(ipPort, busPath)
	bus        = networkBus.EventBus()
)
