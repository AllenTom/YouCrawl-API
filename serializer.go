package api

// for item in request body => real task item
type ItemSerializer interface {
	Unmarshal(rawData map[string]interface{}) (interface{},error)
}

