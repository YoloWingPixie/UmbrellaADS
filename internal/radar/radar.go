package radar

type RadarUnit struct {
}

type RadarCapabilities struct {
}

type SearchRadar struct {
	Radar RadarUnit
}

type EWRadar struct {
	Radar RadarUnit
}

type TrackRadar struct {
	Radar RadarUnit
}

func NewRadar() *RadarUnit {

	return &RadarUnit{}
}

func NewSearchRadar() *SearchRadar {

	return &SearchRadar{}
}

func NewEWRadar() *EWRadar {

	return &EWRadar{}
}

func NewTrackRadar() *TrackRadar {

	return &TrackRadar{}
}

func GetRadarCapabilities() {

}

func Run() {

}
