package channels

type Message struct {
	Message string
	ID      uint32
}

var (
	ProcessStop     = make(chan bool, 8)        // Channel to instruct main to exit
	ClientStop      = make(chan bool, 8)        // Channe lt instruct dcsServer.Client to exit
	ClientCallQueue = make(chan struct{}, 1024) // Queue for RPC calls to the DCS server
	ClientState     = make(chan string, 6)      // Queue for dcsServer.Client to report state
	Target          = make(chan struct{}, 1024) // Queue for target processor
	Update          = make(chan struct{}, 1024) // Queue to notify compontents of unit updates
	Notification    = make(chan Message, 1024)
)
