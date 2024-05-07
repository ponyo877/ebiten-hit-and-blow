package drawable

type Feedback struct {
	hand *Cards
	hint *Cards
}

func NewFeedback(ha *Cards, hi *Cards) *Feedback {
	return &Feedback{ha, hi}
}

func (f *Feedback) Hand() *Cards {
	return f.hand
}

func (f *Feedback) Hint() *Cards {
	return f.hint
}
