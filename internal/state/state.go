package state

import (
	"encoding/json"
	"os"

	"github.com/bilalcaliskan/split-the-tunnel/internal/utils"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/pkg/errors"
)

// State is the struct that holds the state of the application
type State struct {
	Entries []*RouteEntry `json:"entries"`
}

// NewState creates a new State with an empty list of RouteEntry
func NewState() *State {
	return &State{
		Entries: []*RouteEntry{},
	}
}

// RouteEntry is the struct that holds the State of a single route entry
type RouteEntry struct {
	Domain      string   `json:"domain"`
	Gateway     string   `json:"gateway"`
	ResolvedIPs []string `json:"resolvedIPs"`
}

// NewRouteEntry creates a new RouteEntry with the given domain, gateway and resolvedIPs
func NewRouteEntry(domain, gateway string, resolvedIPs []string) *RouteEntry {
	return &RouteEntry{
		Domain:      domain,
		Gateway:     gateway,
		ResolvedIPs: resolvedIPs,
	}
}

// AddEntry adds a new RouteEntry to the State. If the entry already exists, it updates the RouteEntry.ResolvedIPs
func (s *State) AddEntry(entry *RouteEntry) error {
	for _, e := range s.Entries {
		if e.Domain == entry.Domain {
			if utils.SlicesEqual(e.ResolvedIPs, entry.ResolvedIPs) {
				return errors.New(constants.EntryAlreadyExists)
			}

			e.ResolvedIPs = entry.ResolvedIPs
			return s.Write(constants.StateFilePath)
		}
	}

	s.Entries = append(s.Entries, entry)

	return s.Write(constants.StateFilePath)
}

// RemoveEntry removes a RouteEntry from the State
func (s *State) RemoveEntry(domain string) error {
	for i, entry := range s.Entries {
		if entry.Domain == domain {
			s.Entries = append(s.Entries[:i], s.Entries[i+1:]...)
			return s.Write(constants.StateFilePath)
		}
	}

	// target entry not found
	return errors.New(constants.EntryNotFound)
}

// GetEntry returns the RouteEntry for the given domain from the State
func (s *State) GetEntry(domain string) *RouteEntry {
	for i := range s.Entries {
		if s.Entries[i].Domain == domain {
			return s.Entries[i]
		}
	}

	return nil
}

// Read reads the State from the given path
func (s *State) Read(path string) error {
	// Attempt to get the file status
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, create an empty state and write to new file
			return s.Write(path)
		}
		// Some other error occurred
		return err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, s)
	return err
}

// Write writes the State to the given path
func (s *State) Write(path string) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// FromStringSlice deserializes the input string into a slice of DomainInfo structs.
func FromStringSlice(input string) ([]RouteEntry, error) {
	var domainInfos []RouteEntry
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
