package fayl

import "github.com/Polshkrev/gopolutils"

// Finite — incomplete — operating system values.
type OperatingSystem gopolutils.StringEnum

const (
	// Windows operating system.
	Windows OperatingSystem = "windows"
	// Mac operating system.
	Mac OperatingSystem = "darwin"
	// Linux operating system.
	Linux OperatingSystem = "linux"
	// Android operating system.
	Android OperatingSystem = "android"
	// FreeBSD operating system.
	Freebsd OperatingSystem = "freebsd"
	// Representation of Apple's IOS.
	Ios OperatingSystem = "ios"
	// NetBSD operating system.
	Netbsd OperatingSystem = "netbsd"
	// OpenBSD operating system.
	Openbsd OperatingSystem = "openbsd"
)
