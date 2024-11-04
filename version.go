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
