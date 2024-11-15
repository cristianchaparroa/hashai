package thegraph

type ENSResponse struct {
	Data ENSData `json:"data"`
}

type ENSData struct {
	Domains []*ENSDomain `json:"domains"`
}

type ENSDomain struct {
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	Owner           *ENSOwner           `json:"owner"`
	ResolvedAddress *ENSResolvedAddress `json:"resolvedAddress"`
	Resolver        ENSResolver         `json:"resolver"`
}

type ENSOwner struct {
	ID string `json:"id"`
}

type ENSResolvedAddress struct {
	ID string `json:"id"`
}

type ENSResolver struct {
	Address string `json:"address"`
}
