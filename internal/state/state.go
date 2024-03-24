package state

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/bilalcaliskan/split-the-tunnel/internal/utils"

	"github.com/bilalcaliskan/split-the-tunnel/internal/constants"

	"github.com/pkg/errors"
)

// State is the struct that holds the state of the application
type State struct {
	Entries []*RouteEntry `json:"entries"`
	logger  zerolog.Logger
	path    string
}

// NewState creates a new State with an empty list of RouteEntry
func NewState(logger zerolog.Logger, path string) *State {
	return &State{
		[]*RouteEntry{},
		logger,
		path,
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

func (s *State) CheckIPChanges() error {
	if len(s.Entries) == 0 {
		s.logger.Info().Msg("no entries found in the state, skipping ip check")
		return nil
	}

	if applyNeeded := s.updateEntries(); applyNeeded {
		s.logger.Info().Msg("ip changes detected, applying internal state")
		return s.Write()
	}

	s.logger.Info().Msg("no change, skipping state update")

	return nil
}

func (s *State) updateEntries() bool {
	var applyNeeded bool
	for _, entry := range s.Entries {
		ipList, err := utils.ResolveDomain(entry.Domain)
		if err != nil {
			s.logger.Error().Err(err).Str("domain", entry.Domain).Msg("failed to resolve domain")
			continue
		}

		if !utils.SlicesEqual(ipList, entry.ResolvedIPs) {
			s.logger.Info().Str("domain", entry.Domain).Msg("ip changes detected, applying changes to the routing table")
			s.removeOldRoutes(entry)
			entry.ResolvedIPs = ipList
			applyNeeded = true
			s.addNewRoutes(entry)
		}
	}

	return applyNeeded
}

func (s *State) removeOldRoutes(entry *RouteEntry) {
	for _, ip := range entry.ResolvedIPs {
		if err := utils.RemoveRoute(ip); err != nil {
			fmt.Println(err)
		}
	}
}

func (s *State) addNewRoutes(entry *RouteEntry) {
	for _, ip := range entry.ResolvedIPs {
		if err := utils.AddRoute(ip, entry.Gateway); err != nil {
			fmt.Println(err)
		}
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
			return s.Write()
		}
	}

	s.Entries = append(s.Entries, entry)

	return s.Write()
}

// RemoveEntry removes a RouteEntry from the State
func (s *State) RemoveEntry(domain string) error {
	for i, entry := range s.Entries {
		if entry.Domain == domain {
			s.Entries = append(s.Entries[:i], s.Entries[i+1:]...)
			return s.Write()
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

// Reload reads the State from the given path
func (s *State) Reload() error {
	// Attempt to get the file status
	if _, err := os.Stat(s.path); err != nil {
		if os.IsNotExist(err) {
			// File does not exist, create an empty state and write to new file
			return s.Write()
		}
		// Some other error occurred
		return err
	}

	content, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, s)
	return err
}

// Write writes the State to the given path
func (s *State) Write() error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}
