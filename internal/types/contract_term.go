package types

import "time"

const (
	TERM_BEG_NOTLIMIT byte = 0
	TERM_BEG_FINISHED byte = 32
	TERM_END_NOTLIMIT byte = 32
	TERM_END_FINISHED byte = 0
)

type IContractTerm interface {
	Contract() ContractCode
	DateFrom() time.Time
	DateStop() time.Time
	TermDayFrom() byte
	TermDayStop() byte
	IsActive() bool
}

type ContractTerm struct {
	contract ContractCode
	dateFrom time.Time
	dateStop time.Time
	termDayFrom byte
	termDayStop byte
}

func (c ContractTerm) Contract() ContractCode {
	return c.contract
}
func (c ContractTerm) DateFrom() time.Time {
	return c.dateFrom
}
func (c ContractTerm) DateStop() time.Time {
	return c.dateStop
}
func (c ContractTerm) TermDayFrom() byte {
	return c.termDayFrom
}
func (c ContractTerm) TermDayStop() byte {
	return c.termDayStop
}

func (c ContractTerm) IsActive() bool {
	return c.termDayFrom < TERM_BEG_FINISHED && c.termDayStop > TERM_END_FINISHED
}

func NewContractTerm(_contract ContractCode,
	_dateFrom time.Time, _dateStop time.Time,
	_termDayFrom byte, _termDayStop byte) IContractTerm {
	return ContractTerm {
		contract: _contract,
		dateFrom: _dateFrom, dateStop: _dateStop,
		termDayFrom: _termDayFrom, termDayStop: _termDayStop,
	}
}

type IContractTermList = []IContractTerm