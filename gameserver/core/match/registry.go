package match

import (
	"github.com/mrclayman/rest-and-go/gameserver/core/servererrors"
	"github.com/mrclayman/rest-and-go/gameserver/core/player"
)

// Registry aggregates ongoing matches
// of all types
type Registry struct {
	// TODO Hide the members once the dummy data are no longer needed
	DM   DMMatches
	CTF  CTFMatches
	LMS  LMSMatches
	Duel DuelMatches
}

// NewDM creates a new DeathMatch-type match
// and stores it in the registry
func (r *Registry) NewDM(pl player.List) *DMMatch {
	m := newDM(pl)
	r.DM[m.Number] = m
	return m
}

// NewCTF creates a new match and populates
// it with the given set of players
func (r *Registry) NewCTF(pl player.List) *CTFMatch {
	m := newCTF(pl)
	r.CTF[m.Number] = m
	return m
}

// NewLMS creates a new match and populates
// it with the given set of players
func (r *Registry) NewLMS(pl player.List) *LMSMatch {
	m := newLMS(pl)
	r.LMS[m.Number] = m
	return m
}

// NewDuel creates a new match and populates
// it with the given set of players
func (r *Registry) NewDuel(pl player.List) *DuelMatch {
	m := newDuel(pl)
	r.Duel[m.Number] = m
	return m
}

// GetDM looks up a DeathMatch-type match in the
// match registry and returns it, if the number corresponds
// to an existing match. If the match could not be found,
// nil is returned and an error is raised
func (r *Registry) GetDM(nr Number) (*DMMatch, error) {
	m, ok := r.DM[nr]
	if !ok {
		return nil, servererrors.InvalidArgumentError{Message: "DM-type match #" + NumberToString(nr) + " not found"}
	}

	return m, nil
}

// GetCTF looks up a CTF-type match in the
// match registry and returns it, if the number corresponds
// to an existing match. If the match could not be found,
// nil is returned and an error is raised
func (r *Registry) GetCTF(nr Number) (*CTFMatch, error) {
	m, ok := r.CTF[nr]
	if !ok {
		return nil, servererrors.InvalidArgumentError{Message: "CTF-type match #" + NumberToString(nr) + " not found"}
	}

	return m, nil
}

// GetLMS looks up a LMS-type match in the
// match registry and returns it, if the number corresponds
// to an existing match. If the match could not be found,
// nil is returned and an error is raised
func (r *Registry) GetLMS(nr Number) (*LMSMatch, error) {
	m, ok := r.LMS[nr]
	if !ok {
		return nil, servererrors.InvalidArgumentError{Message: "LMS-type match #" + NumberToString(nr) + " not found"}
	}

	return m, nil
}

// GetDuel looks up a Duel-type match in the
// match registry and returns it, if the number corresponds
// to an existing match. If the match could not be found,
// nil is returned and an error is raised
func (r *Registry) GetDuel(nr Number) (*DuelMatch, error) {
	m, ok := r.Duel[nr]
	if !ok {
		return nil, servererrors.InvalidArgumentError{Message: "Duel-type match #" + NumberToString(nr) + " not found"}
	}

	return m, nil
}

// DropDM deletes a DeathMatch-type match
// from the registry. True is returned upon
// successful deletion, false is returned in
// case a match with the given ID could not
// be found
func (r *Registry) DropDM(ID Number) bool {
	_, ok := r.DM[ID]
	if !ok {
		return false
	}

	delete(r.DM, ID)
	return true
}

// DropCTF deletes a CTF-type match
// from the registry. True is returned upon
// successful deletion, false is returned in
// case a match with the given ID could not
// be found
func (r *Registry) DropCTF(ID Number) bool {
	_, ok := r.CTF[ID]
	if !ok {
		return false
	}

	delete(r.CTF, ID)
	return true
}

// DropLMS deletes a LMS-type match
// from the registry. True is returned upon
// successful deletion, false is returned in
// case a match with the given ID could not
// be found
func (r *Registry) DropLMS(ID Number) bool {
	_, ok := r.LMS[ID]
	if !ok {
		return false
	}

	delete(r.LMS, ID)
	return true
}

// DropDuel deletes a Duel-type match
// from the registry. True is returned upon
// successful deletion, false is returned in
// case a match with the given ID could not
// be found
func (r *Registry) DropDuel(ID Number) bool {
	_, ok := r.Duel[ID]
	if !ok {
		return false
	}

	delete(r.Duel, ID)
	return true
}

// GetAllDM returns a slice of all existing
// DeathMatch-type matches
func (r *Registry) GetAllDM() []*DMMatch {

	if len(r.DM) == 0 {
		return make([]*DMMatch, 0)
	}

	var retval []*DMMatch
	for _, m := range r.DM {
		retval = append(retval, m)
	}

	return retval
}

// GetAllCTF returns a slice of all existing
// CTF-type matches
func (r *Registry) GetAllCTF() []*CTFMatch {
	if len(r.CTF) == 0 {
		return make([]*CTFMatch, 0)
	}

	var retval []*CTFMatch
	for _, m := range r.CTF {
		retval = append(retval, m)
	}

	return retval
}

// GetAllLMS returns a slice of all existing
// LMS-type matches
func (r *Registry) GetAllLMS() []*LMSMatch {
	if len(r.LMS) == 0 {
		return make([]*LMSMatch, 0)
	}

	var retval []*LMSMatch
	for _, m := range r.LMS {
		retval = append(retval, m)
	}

	return retval
}

// GetAllDuel returns a slice of all existing
// Duel-type matches
func (r *Registry) GetAllDuel() []*DuelMatch {
	if len(r.Duel) == 0 {
		return make([]*DuelMatch, 0)
	}

	var retval []*DuelMatch
	for _, m := range r.Duel {
		retval = append(retval, m)
	}

	return retval
}
