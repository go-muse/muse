package muse

type GetOptFunc func(*FilteringOptions)

type FilteringOptions struct {
	FilterByDegreeCharacteristicName map[DegreeCharacteristicName]struct{}
	FilterBySonance                  map[Sonance]struct{}
	FilterByAbsoluteModalPosition    map[ModalPositionName]struct{}
}

func NewGetOptions(opts ...GetOptFunc) *FilteringOptions {
	o := &FilteringOptions{
		FilterByDegreeCharacteristicName: make(map[DegreeCharacteristicName]struct{}),
		FilterBySonance:                  make(map[Sonance]struct{}),
		FilterByAbsoluteModalPosition:    make(map[ModalPositionName]struct{}),
	}

	return o.WithOpts(opts...)
}

func (o *FilteringOptions) WithOpts(opts ...GetOptFunc) *FilteringOptions {
	t := *o

	for _, opt := range opts {
		opt(&t)
	}

	return &t
}

func AddFilterByDegreeCharacteristicName(dcNames []DegreeCharacteristicName) GetOptFunc {
	return func(o *FilteringOptions) {
		for _, dcName := range dcNames {
			o.FilterByDegreeCharacteristicName[dcName] = struct{}{}
		}
	}
}

func AddFilterByAbsoluteModalPosition(modalPositionNames []ModalPositionName) GetOptFunc {
	return func(o *FilteringOptions) {
		for _, modalPositionName := range modalPositionNames {
			o.FilterByAbsoluteModalPosition[modalPositionName] = struct{}{}
		}
	}
}

func AddFilterBySonance(sonances []Sonance) GetOptFunc {
	return func(o *FilteringOptions) {
		for _, sonance := range sonances {
			o.FilterBySonance[sonance] = struct{}{}
		}
	}
}

func (o *FilteringOptions) FilterByDegreeCharacteristicNameExist(dcn DegreeCharacteristicName) bool {
	if _, ok := o.FilterByDegreeCharacteristicName[dcn]; ok {
		return true
	}

	return false
}

func (o *FilteringOptions) FilterBySonanceExist(s Sonance) bool {
	if _, ok := o.FilterBySonance[s]; ok {
		return true
	}

	return false
}

func (o *FilteringOptions) FilterByAbsoluteModalPositionExist(s ModalPositionName) bool {
	if _, ok := o.FilterByAbsoluteModalPosition[s]; ok {
		return true
	}

	return false
}
