package ephemeralproxies

type ProxyType int32

const (
	Datacenter ProxyType = iota
	Residential
)

func (t ProxyType) String() string {
	switch t {
	case Datacenter:
		return "datacenter"
	case Residential:
		return "residential"
	}
	panic("invalid proxy type")
}
