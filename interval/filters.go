package interval

import (
	"github.com/go-muse/muse/degree"
)

// GetOptFunc is a function type used to apply options to FilteringOptions.
type GetOptFunc func(*FilteringOptions)

// FilteringOptions holds the filtering criteria for intervals.
type FilteringOptions struct {
	FilterByDegreeCharacteristicName map[degree.CharacteristicName]struct{}
	FilterBySonance                  map[Sonance]struct{}
	FilterByAbsoluteModalPosition    map[degree.ModalPositionName]struct{}
}

// NewIntervalGetOptions creates a new FilteringOptions instance with the provided options.
func NewIntervalGetOptions(opts ...GetOptFunc) *FilteringOptions {
	o := &FilteringOptions{
		FilterByDegreeCharacteristicName: make(map[degree.CharacteristicName]struct{}),
		FilterBySonance:                  make(map[Sonance]struct{}),
		FilterByAbsoluteModalPosition:    make(map[degree.ModalPositionName]struct{}),
	}

	return o.WithOpts(opts...)
}

// WithOpts applies the given options to the FilteringOptions instance.
func (o *FilteringOptions) WithOpts(opts ...GetOptFunc) *FilteringOptions {
	t := *o

	for _, opt := range opts {
		opt(&t)
	}

	return &t
}

// AddFilterByDegreeCharacteristicName adds a filter for degree characteristic names.
func AddFilterByDegreeCharacteristicName(dcNames []degree.CharacteristicName) GetOptFunc {
	return func(o *FilteringOptions) {
		for _, dcName := range dcNames {
			o.FilterByDegreeCharacteristicName[dcName] = struct{}{}
		}
	}
}

// AddFilterByAbsoluteModalPosition adds a filter for absolute modal positions.
func AddFilterByAbsoluteModalPosition(modalPositionNames []degree.ModalPositionName) GetOptFunc {
	return func(o *FilteringOptions) {
		for _, modalPositionName := range modalPositionNames {
			o.FilterByAbsoluteModalPosition[modalPositionName] = struct{}{}
		}
	}
}

// AddFilterBySonance adds a filter for sonances.
func AddFilterBySonance(sonances []Sonance) GetOptFunc {
	return func(o *FilteringOptions) {
		for _, sonance := range sonances {
			o.FilterBySonance[sonance] = struct{}{}
		}
	}
}

// HasFilterByDegreeCharacteristicName checks if a degree characteristic name filter exists.
func (o *FilteringOptions) HasFilterByDegreeCharacteristicName(dcn degree.CharacteristicName) bool {
	_, ok := o.FilterByDegreeCharacteristicName[dcn]

	return ok
}

// HasFilterBySonance checks if a sonance filter exists.
func (o *FilteringOptions) HasFilterBySonance(s Sonance) bool {
	_, ok := o.FilterBySonance[s]

	return ok
}

// HasFilterByAbsoluteModalPosition checks if an absolute modal position filter exists.
func (o *FilteringOptions) HasFilterByAbsoluteModalPosition(s degree.ModalPositionName) bool {
	_, ok := o.FilterByAbsoluteModalPosition[s]

	return ok
}
