package gopolutils

import "fmt"

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

func (version Version) CompareMajor(major uint8) bool {
	return version.Major() >= major
}

func (version Version) CompareMinor(minor uint8) bool {
	return version.Minor() >= minor
}

func (version Version) ComparePatch(patch uint8) bool {
	return version.Patch() >= patch
}

func (version Version) Compare(operand Version) bool {
	return version.CompareMajor(operand.Major()) && version.CompareMinor(operand.Minor()) && version.ComparePatch(operand.Patch())
}

func (version Version) IsPublic() bool {
	return version.CompareMajor(1)
}

func (version *Version) Release() {
	version.major++
	version.SetMinor(0)
	version.SetPatch(0)
}

func (version *Version) Update() {
	version.minor++
	version.SetPatch(0)
}

func (version *Version) Fix() {
	version.patch++
}

func (version Version) ToString() string {
	if len(version.Name()) == 0 && len(version.Description()) == 0 {
		return fmt.Sprintf("%d.%d.%d", version.Major(), version.Minor(), version.Patch())
	} else if len(version.Name()) == 0 && len(version.Description()) != 0 {
		return fmt.Sprintf("%d.%d.%d - %s", version.Major(), version.Minor(), version.Patch(), version.Description())
	} else if len(version.Name()) != 0 && len(version.Description()) == 0 {
		return fmt.Sprintf("%s: %d.%d.%d", version.Name(), version.Major(), version.Minor(), version.Patch())
	}
	return fmt.Sprintf("%s: %d.%d.%d - %s", version.Name(), version.Major(), version.Minor(), version.Patch(), version.Description())
}
