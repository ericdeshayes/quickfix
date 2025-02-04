package relatedpartydetail

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relatedaltptyssubgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/relatedcontextptyssubgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/relationshipriskinstrumentscope"
	"github.com/quickfixgo/quickfix/fix50sp2/relationshipriskwarninglevels"
)

//NoRelatedPartySubIDs is a repeating group in RelatedPartyDetail
type NoRelatedPartySubIDs struct {
	//RelatedPartySubID is a non-required field for NoRelatedPartySubIDs.
	RelatedPartySubID *string `fix:"1567"`
	//RelatedPartySubIDType is a non-required field for NoRelatedPartySubIDs.
	RelatedPartySubIDType *int `fix:"1568"`
}

//NoRelatedPartyAltIDs is a repeating group in RelatedPartyDetail
type NoRelatedPartyAltIDs struct {
	//RelatedPartyAltID is a non-required field for NoRelatedPartyAltIDs.
	RelatedPartyAltID *string `fix:"1570"`
	//RelatedPartyAltIDSource is a non-required field for NoRelatedPartyAltIDs.
	RelatedPartyAltIDSource *string `fix:"1571"`
	//RelatedAltPtysSubGrp Component
	relatedaltptyssubgrp.RelatedAltPtysSubGrp
}

//NoRelatedContextPartyIDs is a repeating group in RelatedPartyDetail
type NoRelatedContextPartyIDs struct {
	//RelatedContextPartyID is a non-required field for NoRelatedContextPartyIDs.
	RelatedContextPartyID *string `fix:"1576"`
	//RelatedContextPartyIDSource is a non-required field for NoRelatedContextPartyIDs.
	RelatedContextPartyIDSource *string `fix:"1577"`
	//RelatedContextPartyRole is a non-required field for NoRelatedContextPartyIDs.
	RelatedContextPartyRole *int `fix:"1578"`
	//RelatedContextPtysSubGrp Component
	relatedcontextptyssubgrp.RelatedContextPtysSubGrp
}

//NoRelationshipRiskLimits is a repeating group in RelatedPartyDetail
type NoRelationshipRiskLimits struct {
	//RelationshipRiskLimitType is a non-required field for NoRelationshipRiskLimits.
	RelationshipRiskLimitType *int `fix:"1583"`
	//RelationshipRiskLimitAmount is a non-required field for NoRelationshipRiskLimits.
	RelationshipRiskLimitAmount *float64 `fix:"1584"`
	//RelationshipRiskLimitCurrency is a non-required field for NoRelationshipRiskLimits.
	RelationshipRiskLimitCurrency *string `fix:"1585"`
	//RelationshipRiskLimitPlatform is a non-required field for NoRelationshipRiskLimits.
	RelationshipRiskLimitPlatform *string `fix:"1586"`
	//RelationshipRiskInstrumentScope Component
	relationshipriskinstrumentscope.RelationshipRiskInstrumentScope
	//RelationshipRiskWarningLevels Component
	relationshipriskwarninglevels.RelationshipRiskWarningLevels
}

//RelatedPartyDetail is a fix50sp2 Component
type RelatedPartyDetail struct {
	//RelatedPartyID is a non-required field for RelatedPartyDetail.
	RelatedPartyID *string `fix:"1563"`
	//RelatedPartyIDSource is a non-required field for RelatedPartyDetail.
	RelatedPartyIDSource *string `fix:"1564"`
	//RelatedPartyRole is a non-required field for RelatedPartyDetail.
	RelatedPartyRole *int `fix:"1565"`
	//NoRelatedPartySubIDs is a non-required field for RelatedPartyDetail.
	NoRelatedPartySubIDs []NoRelatedPartySubIDs `fix:"1566,omitempty"`
	//NoRelatedPartyAltIDs is a non-required field for RelatedPartyDetail.
	NoRelatedPartyAltIDs []NoRelatedPartyAltIDs `fix:"1569,omitempty"`
	//NoRelatedContextPartyIDs is a non-required field for RelatedPartyDetail.
	NoRelatedContextPartyIDs []NoRelatedContextPartyIDs `fix:"1575,omitempty"`
	//NoRelationshipRiskLimits is a non-required field for RelatedPartyDetail.
	NoRelationshipRiskLimits []NoRelationshipRiskLimits `fix:"1582,omitempty"`
}

func (m *RelatedPartyDetail) SetRelatedPartyID(v string)       { m.RelatedPartyID = &v }
func (m *RelatedPartyDetail) SetRelatedPartyIDSource(v string) { m.RelatedPartyIDSource = &v }
func (m *RelatedPartyDetail) SetRelatedPartyRole(v int)        { m.RelatedPartyRole = &v }
func (m *RelatedPartyDetail) SetNoRelatedPartySubIDs(v []NoRelatedPartySubIDs) {
	m.NoRelatedPartySubIDs = v
}
func (m *RelatedPartyDetail) SetNoRelatedPartyAltIDs(v []NoRelatedPartyAltIDs) {
	m.NoRelatedPartyAltIDs = v
}
func (m *RelatedPartyDetail) SetNoRelatedContextPartyIDs(v []NoRelatedContextPartyIDs) {
	m.NoRelatedContextPartyIDs = v
}
func (m *RelatedPartyDetail) SetNoRelationshipRiskLimits(v []NoRelationshipRiskLimits) {
	m.NoRelationshipRiskLimits = v
}
