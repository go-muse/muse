package muse

type IntervalGetOptFunc func(*IntervalFilteringOptions)

type IntervalFilteringOptions struct {
	FilterByDegreeCharacteristicName map[DegreeCharacteristicName]struct{}
	FilterBySonance                  map[Sonance]struct{}
	FilterByAbsoluteModalPosition    map[ModalPositionName]struct{}
}

func NewIntervalGetOptions(opts ...IntervalGetOptFunc) *IntervalFilteringOptions {
	o := &IntervalFilteringOptions{
		FilterByDegreeCharacteristicName: make(map[DegreeCharacteristicName]struct{}),
		FilterBySonance:                  make(map[Sonance]struct{}),
		FilterByAbsoluteModalPosition:    make(map[ModalPositionName]struct{}),
	}

	return o.WithOpts(opts...)
}

func (o *IntervalFilteringOptions) WithOpts(opts ...IntervalGetOptFunc) *IntervalFilteringOptions {
	t := *o

	for _, opt := range opts {
		opt(&t)
	}

	return &t
}

func AddFilterByDegreeCharacteristicName(dcNames []DegreeCharacteristicName) IntervalGetOptFunc {
	return func(o *IntervalFilteringOptions) {
		for _, dcName := range dcNames {
			o.FilterByDegreeCharacteristicName[dcName] = struct{}{}
		}
	}
}

func AddFilterByAbsoluteModalPosition(modalPositionNames []ModalPositionName) IntervalGetOptFunc {
	return func(o *IntervalFilteringOptions) {
		for _, modalPositionName := range modalPositionNames {
			o.FilterByAbsoluteModalPosition[modalPositionName] = struct{}{}
		}
	}
}

func AddFilterBySonance(sonances []Sonance) IntervalGetOptFunc {
	return func(o *IntervalFilteringOptions) {
		for _, sonance := range sonances {
			o.FilterBySonance[sonance] = struct{}{}
		}
	}
}

func (o *IntervalFilteringOptions) FilterByDegreeCharacteristicNameExist(dcn DegreeCharacteristicName) bool {
	if _, ok := o.FilterByDegreeCharacteristicName[dcn]; ok {
		return true
	}

	return false
}

func (o *IntervalFilteringOptions) FilterBySonanceExist(s Sonance) bool {
	if _, ok := o.FilterBySonance[s]; ok {
		return true
	}

	return false
}

func (o *IntervalFilteringOptions) FilterByAbsoluteModalPositionExist(s ModalPositionName) bool {
	if _, ok := o.FilterByAbsoluteModalPosition[s]; ok {
		return true
	}

	return false
}
