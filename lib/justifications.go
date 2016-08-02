package lib

import (
	"sync"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

// Verification struct holds data for a specific component and verification
// This is an internal data structure that helps map standards and controls to components
type Verification struct {
	ComponentKey  string
	SatisfiesData common.Satisfies
}

// Verifications is a slice of type Verifications
type Verifications []Verification

// Justifications struct contains the mapping that links controls to specific components
type Justifications struct {
	mapping map[string]map[string]Verifications
	sync.RWMutex
}

// Len returns the length of the GeneralReferences slice
func (slice Verifications) Len() int {
	return len(slice)
}

// Less returns true if a GeneralReference is less than another reference
func (slice Verifications) Less(i, j int) bool {
	return slice[i].ComponentKey < slice[j].ComponentKey
}

// Swap swaps the two GeneralReferences
func (slice Verifications) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// NewJustifications creates a new justification
func NewJustifications() *Justifications {
	return &Justifications{mapping: make(map[string]map[string]Verifications)}
}

// Add methods adds a new mapping to the justification while locking
func (justifications *Justifications) Add(standardKey string, controlKey string, componentKey string, satisfies common.Satisfies) {
	justifications.Lock()
	newVerification := Verification{componentKey, satisfies}
	_, standardKeyExists := justifications.mapping[standardKey]
	if !standardKeyExists {
		justifications.mapping[standardKey] = make(map[string]Verifications)
	}
	_, controlKeyExists := justifications.mapping[standardKey][controlKey]
	if !controlKeyExists {
		justifications.mapping[standardKey][controlKey] = Verifications{}
	}
	justifications.mapping[standardKey][controlKey] = append(
		justifications.mapping[standardKey][controlKey], newVerification,
	)
	justifications.Unlock()
}

// LoadMappings loads a set of mappings from a component
func (justifications *Justifications) LoadMappings(component common.Component) {
	for _, satisfies := range component.GetAllSatisfies() {
		justifications.Add(satisfies.GetStandardKey(), satisfies.GetControlKey(), component.GetKey(), satisfies)
	}
}

// Get retrieves justifications for a specific standard and control
func (justifications *Justifications) Get(standardKey string, controlKey string) Verifications {
	_, standardKeyExists := justifications.mapping[standardKey]
	if !standardKeyExists {
		return nil
	}
	controlJustifications, controlKeyExists := justifications.mapping[standardKey][controlKey]
	if !controlKeyExists {
		return nil
	}
	return controlJustifications
}
