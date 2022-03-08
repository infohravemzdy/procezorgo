package types

import "time"

type IPositionTerm interface {
	Contract() ContractCode
	Position() PositionCode
	BaseTerm() IContractTerm
	DateFrom() time.Time
	DateStop() time.Time
	TermDayFrom() byte
	TermDayStop() byte
	IsActive() bool
}

type PositionTerm struct {
	contract ContractCode
	position PositionCode
	baseTerm IContractTerm
	dateFrom time.Time
	dateStop time.Time
	termDayFrom byte
	termDayStop byte
}

func (p PositionTerm) Contract() ContractCode {
	return p.contract
}
func (p PositionTerm) Position() PositionCode {
	return p.position
}
func (p PositionTerm) BaseTerm() IContractTerm {
	return p.baseTerm
}
func (p PositionTerm) DateFrom() time.Time {
	return p.dateFrom
}
func (p PositionTerm) DateStop() time.Time {
	return p.dateStop
}
func (p PositionTerm) TermDayFrom() byte {
	return p.termDayFrom
}
func (p PositionTerm) TermDayStop() byte {
	return p.termDayStop
}

func (p PositionTerm) isPositionActive() bool {
	return p.termDayFrom < TERM_BEG_FINISHED && p.termDayStop > TERM_END_FINISHED
}
func (p PositionTerm) IsActive() bool {
	if p.baseTerm != nil {
		return p.baseTerm.IsActive() && p.isPositionActive()
	}
	return p.isPositionActive()
}

func NewPositionTerm(_contract ContractCode, _position PositionCode, _baseTerm IContractTerm,
	_dateFrom time.Time, _dateStop time.Time,
	_termDayFrom byte, _termDayStop byte) IPositionTerm {
	return PositionTerm {
		contract: _contract, position: _position, baseTerm: _baseTerm,
		dateFrom: _dateFrom, dateStop: _dateStop,
		termDayFrom: _termDayFrom, termDayStop: _termDayStop,
	}
}

type IPositionTermList = []IPositionTerm
