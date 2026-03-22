package fayl

import "github.com/Polshkrev/gopolutils"

// Finite — incomplete — operating system values.
type Os gopolutils.StringEnum

const (
	// Windows operating system.
	Windows Os = "windows"
	// Mac operating system.
	Mac Os = "darwin"
	// Linux operating system.
	Linux Os = "linux"
	// Android operating system.
	Android Os = "android"
	// FreeBSD operating system.
	Freebsd Os = "freebsd"
	// Representation of Apple's IOS.
	Ios Os = "ios"
	// NetBSD operating system.
	Netbsd Os = "netbsd"
	// OpenBSD operating system.
	Openbsd Os = "openbsd"
)
