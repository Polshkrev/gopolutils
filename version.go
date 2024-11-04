package gopolutils

type Version struct {
	name        string
	description string
	major       uint8
	minor       uint8
	patch       uint8
}

func NewVersion() *Version {
	return new(Version)
}

func VersionConvert(major, minor, patch uint8) *Version {
	var version *Version = NewVersion()
	version.major = major
	version.minor = minor
	version.patch = patch
	return version
}

func NewNamedVersion(name string) *Version {
	var version *Version = NewVersion()
	version.name = name
	return version
}

func NewStringVersion(name, description string) *Version {
	var version *Version = NewNamedVersion(name)
	version.description = description
	return version
}

func NewFullVersion(name, description string, major, minor, patch uint8) *Version {
	var version *Version = VersionConvert(major, minor, patch)
	version.name = name
	version.description = description
	return version
}
