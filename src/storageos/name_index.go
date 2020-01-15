package storageos

type nameIndex struct {
	Prefix string `json:"prefix"`
	Key    string `json:"key"`
	ID     string `json:"objectID"`
}

func (ni *nameIndex) getID() string {
	return ni.ID
}
