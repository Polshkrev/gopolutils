package fayl

import "github.com/Polshkrev/gopolutils"

// Finite — incomplete — operating system values.
type OS = gopolutils.StringEnum

const (
	// Windows operating system.
	WINDOWS OS = "windows"
	// Mac operating system.
	MAC OS = "darwin"
	// Linux operating system.
	LINUX OS = "linux"
	// Android operating system.
	ANDROID OS = "android"
	// FreeBSD operating system.
	FREEBSD OS = "freebsd"
	// Representation of Apple's IOS.
	IOS OS = "ios"
	// NetBSD operating system.
	NETBSD OS = "netbsd"
	// OpenBSD operating system.
	OPENBSD OS = "openbsd"
)
