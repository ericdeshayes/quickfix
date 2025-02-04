package instrumentextension

//NoInstrAttrib is a repeating group in InstrumentExtension
type NoInstrAttrib struct {
	//InstrAttribType is a non-required field for NoInstrAttrib.
	InstrAttribType *int `fix:"871"`
	//InstrAttribValue is a non-required field for NoInstrAttrib.
	InstrAttribValue *string `fix:"872"`
}

//InstrumentExtension is a fix50sp1 Component
type InstrumentExtension struct {
	//DeliveryForm is a non-required field for InstrumentExtension.
	DeliveryForm *int `fix:"668"`
	//PctAtRisk is a non-required field for InstrumentExtension.
	PctAtRisk *float64 `fix:"869"`
	//NoInstrAttrib is a non-required field for InstrumentExtension.
	NoInstrAttrib []NoInstrAttrib `fix:"870,omitempty"`
}

func (m *InstrumentExtension) SetDeliveryForm(v int)              { m.DeliveryForm = &v }
func (m *InstrumentExtension) SetPctAtRisk(v float64)             { m.PctAtRisk = &v }
func (m *InstrumentExtension) SetNoInstrAttrib(v []NoInstrAttrib) { m.NoInstrAttrib = v }
