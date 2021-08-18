module github.com/sonkwo/EncodingModifier

go 1.12

require (
	github.com/gogs/chardet v0.0.0-20191104214054-4b6791f73a28
	github.com/wsw364321644/go-botil v0.1.6
	golang.org/x/text v0.3.0
)

replace (
	github.com/wsw364321644/go-botil v0.1.6 => ../go-botil
)