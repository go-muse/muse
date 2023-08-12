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
- [x] Comparison
- [x] Alterations
- [x] Octaves
- [x] Durations
- [x] MIDI numbering

### Tracks:
- [x] Adding notes as events
- [x] Note timing calculation
- [x] Getting sorted start/end events
### Chords:
- [x] Adding and removing Notes
- [x] Setting the same duration for all the chord's notes

### Modes:
- [x] Templates* of most commonly used modes
- [x] Creating a mode based on mode template and tonic
- [x] Calculation of modal positions of degrees in seven-degree modes
- [x] Finding modes from an incoming set of degrees or notes
### Scales:
- [x] Generating scales
### Intervals:
- [x] Templates of intervals within an octave
- [x] Identifying the interval between two degrees
- [x] Acoustic and Chromatic intervals

\* In this context, mode templates are a set of intervals that determine the arrangement of degrees in a mode.

## Concept
Definitions may differ from conventional.

**Mode** is a set of degrees within octave spaced apart by a certain number of semitones.

**Degree** is a position of a note within a mode, possessing modal characteristics. 

**Note** is the representation of a musical sound. It's name represent a pitch class.

**Scale** is the set of musical notes got from the mode and ordered by fundamental frequency or pitch.

**Interval** is the ratio of two musical sounds in pitch. 
Chromatic interval means the ratio of two notes, diatonic interval means the ratio of two degrees.

**Chord** is a set of notes of the same duration.

**Track** is a set of events(e.g. notes) with starting time. 

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

Q: How to find a mode by pattern?\
A: There is the storage of mode names and their templates. You can find them by specifying your template. [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-ModeTemplatesStore.FindModeTemplatesByPattern) Then you can sort them by mode names [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-ModeNamesAndTemplates.SortByModeName) or by templates(patterns) [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-ModeNamesAndTemplates.SortByModeTemplate).

Q: How can I find a mode based on a given set of notes?\
A: You can find them in the storage of mode names and their templates by specifying your set of notes. [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-ModeTemplatesStore.FindModeTemplatesByNotes).

## Examples

Q: How to create a note?\
A: It is preferable to create a note by its name, taken from the library [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-NoteName.NewNote). Also, there are other ways to create notes [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Note).

Q: How to specify an octave to a note?\
A: You can create an octave [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Octave) and set it to your note [![example](https://img.shields.io/badge/example-link1-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Octave.SetToNote) [![example](https://img.shields.io/badge/example-link2-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Note.SetOctave).

Q: How to know MIDI number of a note?\
A: For this, you need to assign an octave to the note then call the method to get the midi number [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Note.MIDINumber). Also, you can create a note from known midi number [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-NewNoteFromMIDINumber).

Q: How to set a note's duration?\
A: There are two types of durations. The first type is a relative duration, based on the [note value](https://en.wikipedia.org/wiki/Note_value) [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-NewDuration). Eventually, knowing the bpm, unit, and time signature, you'll be able to get the absolute time duration of the note [![example](https://img.shields.io/badge/example-link1-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Duration.TimeDuration), [![example](https://img.shields.io/badge/example-link2-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Note.TimeDuration). The second type is absolute duration. You can set your time.Duration as absolute duration [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Note.SetDurationAbs). These types are separate values and do not affect each other.

Q: How to make modes?\
A: You can make it by specifying correct names (from the muse package) [![example](https://img.shields.io/badge/example-link1-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-MakeNewMode) [![example](https://img.shields.io/badge/example-link2-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-MustMakeNewMode) or you can create custom mode. [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-MakeNewCustomMode)

Q: How to iterate mode?\
A: You can iterate through degrees  [![example](https://img.shields.io/badge/example-link1-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Mode.IterateOneRound) [![example](https://img.shields.io/badge/example-link2-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-DegreesIterator.GetAllDegrees) or notes [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-DegreesIterator.GetAllNotes) of modes.

Q: Are there already existing scales in the library?\
A: Yes, there is mode templates store [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-InitModeTemplatesStore). It contains mode names and templates.





Q: How to use chords?\
A: Chord is a set of notes with the same duration  [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-NewChord). You can add your notes to the chord, their durations will be set to chord's duration automatically [![example](https://img.shields.io/badge/example-link-blue)](https://pkg.go.dev/github.com/go-muse/muse#example-Chord.AddNotes).

---
<font size="2">This is not very stable version.
Public function names and parameters can be changed.
Bugs are OK.\
Feel free to contact me\
</font>
[![GitHub](https://custom-icon-badges.demolab.com/badge/-email-white?logo=gmail)](mailto:igor.georgia.sx@gmail.com)
[![GitHub](https://custom-icon-badges.demolab.com/badge/-@Igorstx-white?logo=telegram)](https://t.me/Igorstx)
