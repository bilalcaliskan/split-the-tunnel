package state

import (
	"encoding/json"
	"os"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/pkg/errors"
)

// State is the struct that holds the state of the application
type State struct {
	Entries []*RouteEntry `json:"entries"`
}

func NewState() *State {
	return &State{
		Entries: []*RouteEntry{},
	}
}

// RouteEntry is the struct that holds the state of a single route entry
type RouteEntry struct {
	Domain     string `json:"domain"`
	ResolvedIP string `json:"resolvedIP"`
	Gateway    string `json:"gateway"`
}

func NewRouteEntry(domain, resolvedIP, gateway string) *RouteEntry {
	return &RouteEntry{
		Domain:     domain,
		ResolvedIP: resolvedIP,
		Gateway:    gateway,
	}
}

// AddEntry adds a new route entry to the state. If the entry already exists, it updates the ResolvedIP.
func (s *State) AddEntry(entry *RouteEntry) error {
	for _, e := range s.Entries {
		if e.Domain == entry.Domain {
			if e.ResolvedIP == entry.ResolvedIP {
				return errors.New(constants.EntryAlreadyExists)
			}

			e.ResolvedIP = entry.ResolvedIP
			return s.Write("/tmp/state.json")
		}
	}

	s.Entries = append(s.Entries, entry)

	return s.Write("/tmp/state.json")
}

// RemoveEntry removes a route entry from the state.
func (s *State) RemoveEntry(domain string) error {
	for i, entry := range s.Entries {
		if entry.Domain == domain {
			s.Entries = append(s.Entries[:i], s.Entries[i+1:]...)
			return s.Write("/tmp/state.json")
		}
	}

	// target entry not found
	return errors.New(constants.EntryNotFound)
}

func (s *State) GetEntry(domain string) *RouteEntry {
	for i := range s.Entries {
		if s.Entries[i].Domain == domain {
			return s.Entries[i]
		}
	}

	return nil
}

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

func (s *State) Write(path string) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
