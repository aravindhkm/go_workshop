package interfacehandle

type Crd interface {
	PodStatus
}

type PodStatus interface {
	Connect() (string, error)
}

type GameServer struct {
	Name string
	Port int
}

func (gs *GameServer) Connect() (string, error) {
	return "GameServer Connected Successfullly", nil
}

type Fleet struct {
	Name string
	Port int
}

func (gs *Fleet) Connect() (string, error) {
	return "Fleet Connected Successfullly", nil
}

// func GetGSS(gs GameServer, flt Fleet) Crd {

// }

func ItrOwnWork() {
	// gs := GameServer{"Gs 1", 4000}
	// flt := Fleet{"flt 1", 4000}

}
