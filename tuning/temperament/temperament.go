// Package temperament provides implementations of various musical temperaments.
//
// A temperament defines how musical intervals are tuned relative to each other.
// Different temperaments have different characteristics and are suited for
// different musical styles and historical periods.
//
// Available temperaments:
//   - Equal: modern equal temperament (12-TET)
//   - Just: just intonation based on pure ratios
//   - Pythagorean: based on pure perfect fifths
//   - Meantone: quarter-comma meantone
//   - WerckmeisterIII: well temperament from 1691
//   - KirnbergerIII: well temperament from 1779
//   - Vallotti: well temperament c. 1730
//   - Young: Thomas Young's well temperament from 1799
package temperament

// standardTwelveTone is the number of steps in standard 12-tone equal temperament.
const standardTwelveTone = 12
