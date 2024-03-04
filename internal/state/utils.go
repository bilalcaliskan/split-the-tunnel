package state

import "encoding/json"

// FromStringSlice deserializes the input string into a slice of DomainInfo structs.
func FromStringSlice(input string) ([]*RouteEntry, error) {
	var domainInfos []*RouteEntry
	err := json.Unmarshal([]byte(input), &domainInfos)
	if err != nil {
		return nil, err
	}
	return domainInfos, nil
}

// ToStringSlice serializes a slice of DomainInfo into a JSON string.
func ToStringSlice(domainInfos []*RouteEntry) (string, error) {
	bytes, err := json.Marshal(domainInfos)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
