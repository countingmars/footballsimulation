package zone

type Stats struct {
	Grade float32
	LastPoint int
}
func (stats Stats) Clone() Stats {
	return Stats{}
}
