package match

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/errors"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

// Registry aggregates ongoing matches
// of all types
type Registry struct {
	dm   DMMatches
	ctf  CTFMatches
	lms  LMSMatches
	duel DuelMatches
}

// NewDM creates a new DeathMatch-type match
// and stores it in the registry
func (r *Registry) NewDM(pl player.List) *DMMatch {
	m := &DMMatch{
		ID: GenerateNumber(),
	}

	for _, p := range pl {
		m.Add(p)
	}

	r.dm[m.ID] = m
	return m
}

// NewCTF creates a new match and populates
// it with the given set of players
func (r *Registry) NewCTF(pl player.List) *CTFMatch {
	m := &CTFMatch{
		ID: GenerateNumber(),
	}

	for _, p := range pl {
		m.Add(p)
	}

	r.ctf[m.ID] = m
	return m
}

// NewLMS creates a new match and populates
// it with the given set of players
func (r *Registry) NewLMS(pl player.List) *LMSMatch {
	m := &LMSMatch{
		ID: GenerateNumber(),
	}

	for _, p := range pl {
		m.Add(p)
	}

	r.lms[m.ID] = m
	return m
}

// NewDuel creates a new match and populates
// it with the given set of players
func (r *Registry) NewDuel(pl player.List) *DuelMatch {
	m := &DuelMatch{
		ID: GenerateNumber(),
	}

	for _, p := range pl {
		m.Add(p)
	}

	r.duel[m.ID] = m
	return m
}

// GetDM looks up a DeathMatch-type match in the
// match registry and returns it, if the number corresponds
// to an existing match. If the match could not be found,
// nil is returned and an error is raised
func (r Registry) GetDM(nr Number) (*DMMatch, error) {
	m, ok := r.dm[nr]
	if !ok {
		return nil, errors.InvalidArgumentError{Message: "DM-type match #" + NumberToString(nr) + " not found"}
	}

	return m, nil
}

// GetCTF looks up a CTF-type match in the
// match registry and returns it, if the number corresponds
// to an existing match. If the match could not be found,
// nil is returned and an error is raised
func (r Registry) GetCTF(nr Number) (*CTFMatch, error) {
	m, ok := r.ctf[nr]
	if !ok {
		return nil, errors.InvalidArgumentError{Message: "CTF-type match #" + NumberToString(nr) + " not found"}
	}

	return m, nil
}

// GetLMS looks up a LMS-type match in the
// match registry and returns it, if the number corresponds
// to an existing match. If the match could not be found,
// nil is returned and an error is raised
func (r Registry) GetLMS(nr Number) (*LMSMatch, error) {
	m, ok := r.lms[nr]
	if !ok {
		return nil, errors.InvalidArgumentError{Message: "LMS-type match #" + NumberToString(nr) + " not found"}
	}

	return m, nil
}

// GetDuel looks up a Duel-type match in the
// match registry and returns it, if the number corresponds
// to an existing match. If the match could not be found,
// nil is returned and an error is raised
func (r Registry) GetDuel(nr Number) (*DuelMatch, error) {
	m, ok := r.duel[nr]
	if !ok {
		return nil, errors.InvalidArgumentError{Message: "Duel-type match #" + NumberToString(nr) + " not found"}
	}

	return m, nil
}

// DropDM deletes a DeathMatch-type match
// from the registry. True is returned upon
// successful deletion, false is returned in
// case a match with the given ID could not
// be found
func (r *Registry) DropDM(ID Number) bool {
	_, ok := r.dm[ID]
	if !ok {
		return false
	}

	delete(r.dm, ID)
	return true
}

// DropCTF deletes a CTF-type match
// from the registry. True is returned upon
// successful deletion, false is returned in
// case a match with the given ID could not
// be found
func (r *Registry) DropCTF(ID Number) bool {
	_, ok := r.ctf[ID]
	if !ok {
		return false
	}

	delete(r.ctf, ID)
	return true
}

// DropLMS deletes a LMS-type match
// from the registry. True is returned upon
// successful deletion, false is returned in
// case a match with the given ID could not
// be found
func (r *Registry) DropLMS(ID Number) bool {
	_, ok := r.lms[ID]
	if !ok {
		return false
	}

	delete(r.lms, ID)
	return true
}

// DropDuel deletes a Duel-type match
// from the registry. True is returned upon
// successful deletion, false is returned in
// case a match with the given ID could not
// be found
func (r *Registry) DropDuel(ID Number) bool {
	_, ok := r.duel[ID]
	if !ok {
		return false
	}

	delete(r.duel, ID)
	return true
}

// GetAllDM returns a slice of all existing
// DeathMatch-type matches
func (r Registry) GetAllDM() []*DMMatch {
	retval := make([]*DMMatch, len(r.dm))

	for _, m := range r.dm {
		retval = append(retval, m)
	}

	return retval
}

// GetAllCTF returns a slice of all existing
// CTF-type matches
func (r Registry) GetAllCTF() []*CTFMatch {
	retval := make([]*CTFMatch, len(r.ctf))

	for _, m := range r.ctf {
		retval = append(retval, m)
	}

	return retval
}

// GetAllLMS returns a slice of all existing
// DeathMatch-type matches
func (r Registry) GetAllLMS() []*LMSMatch {
	retval := make([]*LMSMatch, len(r.lms))

	for _, m := range r.lms {
		retval = append(retval, m)
	}

	return retval
}

// GetAllDuel returns a slice of all existing
// DeathMatch-type matches
func (r Registry) GetAllDuel() []*DuelMatch {
	retval := make([]*DuelMatch, len(r.duel))

	for _, m := range r.duel {
		retval = append(retval, m)
	}

	return retval
}
