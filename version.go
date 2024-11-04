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

func (version Version) Name() string {
	return version.name
}

func (version Version) Description() string {
	return version.description
}

func (version Version) Major() uint8 {
	return version.major
}

func (version Version) Minor() uint8 {
	return version.minor
}

func (version Version) Patch() uint8 {
	return version.patch
}

func (version *Version) SetName(name string) {
	version.name = name
}

func (version *Version) SetDescription(description string) {
	version.description = description
}

func (version *Version) SetMajor(major uint8) {
	version.major = major
}

func (version *Version) SetMinor(minor uint8) {
	version.minor = minor
}

func (version *Version) SetPatch(patch uint8) {
	version.patch = patch
}
