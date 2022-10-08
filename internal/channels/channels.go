package channels

var (
	ProcessStop     = make(chan bool, 8)        // Channel to instruct main to exit
	ClientStop      = make(chan bool, 8)        // Channe lt instruct dcsServer.Client to exit
	ClientCallQueue = make(chan struct{}, 1024) // Queue for RPC calls to the DCS server
	ClientState     = make(chan bool, 8)        // Queue for dcsServer.Client to report state
	Target          = make(chan struct{}, 1024) // Queue for target processor
	Update          = make(chan struct{}, 1024) // Queue to notify compontents of unit updates
	Logs            = make(chan string, 8)
	IADSState       = make(chan bool, 8)
	IADSStop        = make(chan bool, 8)
	PowerState      = make(chan bool, 8)
	PowerStop       = make(chan bool, 8)
	NetworkState    = make(chan bool, 8)
	NetworkStop     = make(chan bool, 8)
	RadarState      = make(chan bool, 8)
	RadarStop       = make(chan bool, 8)
	DCSState        = make(chan bool, 8)
	MissionState    = make(chan int, 8) // 0 = stopped, 1 = running, 2 = mission changed
)
