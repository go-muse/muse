# muse <img src="https://github.com/go-muse/muse/blob/main/img/muse-logo.png?raw=true" alt="muse" width="100" align="right">
## Toolkit for working with musical theory concepts - modes, scales, degrees, notes etc.
![GitHub release (latest SemVer)](https://custom-icon-badges.demolab.com/github/v/release/go-muse/muse?logo=tag&sort=semver)
[![codecov](https://codecov.io/gh/go-muse/muse/branch/main/graph/badge.svg?token=cVbYqHMT3o)](https://codecov.io/gh/go-muse/muse)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-muse/muse)](https://goreportcard.com/report/github.com/go-muse/muse)
![GitHub go.mod Go version](https://custom-icon-badges.demolab.com/github/go-mod/go-version/go-muse/muse?color=gray&label=%20&logo=go)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-muse/muse.svg)](https://pkg.go.dev/github.com/go-muse/muse)
![GitHub](https://custom-icon-badges.demolab.com/github/license/go-muse/muse?logo=law)
[![Pull requests are welcome](https://custom-icon-badges.demolab.com/badge/PRs-welcome!-green?logo=git-pull-request)](https://github.com/go-muse/muse/pulls)
[![Buy me a coffee](https://custom-icon-badges.demolab.com/badge/buyme-acoffee-brightgreen?logo=buymeacoffee)](https://www.buymeacoffee.com/igormuse)

## Functionality
### Notes:
- [x] Name comparison
- [ ] Octaves
- [ ] Chords

### Modes:
- [x] Templates of most commonly used modes
- [x] Creating a mode based on formula and tonic
- [x] Calculation of modal positions of degrees in seven-degree modes
- [x] Finding modes from an incoming set of degres
### Scales:
- [x] Generating scales
### Intervals:
- [x] Templates of intervals within an octave
- [x] Identifying the interval between two degrees
- [x] Acoustic and Chromatic intervals

## Concept
Definitions may differ from conventional.

**Mode** is a set of degrees within octave spaced apart by a certain number of semitones.

**Degree** is a position of a note within a mode, possessing modal characteristics. 

**Note** is the representation of a musical sound. It's name represent a pitch class.

**Scale** is the set of musical notes got from the mode and ordered by fundamental frequency or pitch.

**Interval** is the ratio of two musical sounds in pitch. 
Chromatic interval means the ratio of two notes, diatonic interval means the ratio of two degrees.

![Concept pic](/img/concept.jpg)

## Musician's usecases

Q: How can I find out the set of notes in a chosen mode constructed from a given tonic?\
A: You need to create new mode and get it's scale. [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Mode.GenerateScale)

Q: How to determine the next note by interval?\
A: You can do it by interval's name. [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-MakeNoteByIntervalName)

Q: How to determine the next degree by interval?\
A: You can do it by interval's name. [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-MakeDegreeByIntervalName)

Q: How to know interval by two notes?\
A: You need to create interval from that notes. [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-NewIntervalChromatic)

Q: How to know interval by two degrees?\
A: You need to create interval from that degrees. [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-NewIntervalByDegrees)

Q: How to get notes from a mode sorted by their modal positions?\
A: You need to create a mode from the selected tonic and call sorting function. [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Mode.SortByAbsoluteModalPositions)

---
<font size="2">This is not very stable version.
Public function names and parameters can be changed.
Bugs are OK.
Feel free to contact me\
</font>
[![GitHub](https://custom-icon-badges.demolab.com/badge/-email-white?logo=gmail)](mailto:igor.georgia.sx@gmail.com)
[![GitHub](https://custom-icon-badges.demolab.com/badge/-@Igorstx-white?logo=telegram)](https://t.me/Igorstx)
