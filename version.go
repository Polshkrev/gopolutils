package gopolutils

import "fmt"

// Representation of a semantic versioning object.
type Version struct {
	name        string
	description string
	major       uint8
	minor       uint8
	patch       uint8
}

// Construct a new zero-initialized version object.
// All the string properties are empty and the numeric properties are set to zero.
// Returns a pointer to a new version object.
func NewVersion() *Version {
	return new(Version)
}

// Construct a new version object with given numeric values.
// The string properties of the version object are empty.
// Returns a pointer to a new version object with each of the numeric properties initialized with the given parametres.
func VersionConvert(major, minor, patch uint8) *Version {
	var version *Version = NewVersion()
	version.major = major
	version.minor = minor
	version.patch = patch
	return version
}

// Construct a new version object initialized with a given name.
// Returns a pointer to a new version object initialized with the given name.
func NewNamedVersion(name string) *Version {
	var version *Version = NewVersion()
	version.name = name
	return version
}

// Construct a new version object with each of its string properties initialized.
// Returns a pointer to a new version object initialized with the given string parametres.
func NewStringVersion(name, description string) *Version {
	var version *Version = NewNamedVersion(name)
	version.description = description
	return version
}

// Construct a full initialized version object.
// Returns a pointer to a new fully initialized version object.
func NewFullVersion(name, description string, major, minor, patch uint8) *Version {
	var version *Version = VersionConvert(major, minor, patch)
	version.name = name
	version.description = description
	return version
}

// Access the name property of the version object.
// Returns the name property of the version object.
func (version Version) Name() string {
	return version.name
}

// Access the description property of the version object.
// Returns the description property of the version object.
func (version Version) Description() string {
	return version.description
}

// Access the major property of the version object.
// Returns the major property of the version object.
func (version Version) Major() uint8 {
	return version.major
}

// Access the minor property of the version object.
// Returns the minor property of the version object.
func (version Version) Minor() uint8 {
	return version.minor
}

// Access the patch property of the version object.
// Returns the patch property of the version object.
func (version Version) Patch() uint8 {
	return version.patch
}

// Set the name property of the version object.
func (version *Version) SetName(name string) {
	version.name = name
}

// Set the description property of the version object.
func (version *Version) SetDescription(description string) {
	version.description = description
}

// Set the major property of the version object.
func (version *Version) SetMajor(major uint8) {
	version.major = major
}

// Set the minor property of the version object.
func (version *Version) SetMinor(minor uint8) {
	version.minor = minor
}

// Set the patch property of the version object.
func (version *Version) SetPatch(patch uint8) {
	version.patch = patch
}

// Determine if the version object's major property is greater than or equal to the given operand.
// Returns true if the version object's major property is greater than or equal to the given operand.
func (version Version) CompareMajor(major uint8) bool {
	return version.Major() >= major
}

// Determine if the version object's minor property is greater than or equal to the given operand.
// Returns true if the version object's minor property is greater than or equal to the given operand.
func (version Version) CompareMinor(minor uint8) bool {
	return version.Minor() >= minor
}

// Determine if the version object's patch property is greater than or equal to the given operand.
// Returns true if the version object's patch property is greater than or equal to the given operand.
func (version Version) ComparePatch(patch uint8) bool {
	return version.Patch() >= patch
}

// Compare each of the numeric properties of the version object to a given operand.
// Returns true if each of the version object's properties are greater than or equal to the given operand's numeric properties.
func (version Version) Compare(operand Version) bool {
	return version.CompareMajor(operand.Major()) && version.CompareMinor(operand.Minor()) && version.ComparePatch(operand.Patch())
}

// Determine if the version object is public.
// Returns true if the version object's major property is evaluated greater than or equal to 1.
func (version Version) IsPublic() bool {
	return version.CompareMajor(1)
}

// Publish a version object.
// Set the version object's major property to 1.
// Zero-out all other numeric properties.
// If the version object is evaluated to have already been published, a ValueError is returned and no properties are modified.
func (version *Version) Publish() *Exception {
	if version.IsPublic() {
		return NewNamedException("ValueError", "Version is already public.")
	}
	version.SetMajor(1)
	version.SetMinor(0)
	version.SetPatch(0)
	return nil
}

// Increment the version object's major property.
// The version object's minor and patch properties are set to 0.
func (version *Version) Release() {
	version.major++
	version.SetMinor(0)
	version.SetPatch(0)
}

// Increment the version object's minor property.
// The version object's patch property is set to 0.
// The version object's major version is not modified.
func (version *Version) Update() {
	version.minor++
	version.SetPatch(0)
}

// Increment the version object's patch property.
// The version object's major and minor property are not modified.
func (version *Version) Fix() {
	version.patch++
}

// Render a string representation of the version object.
// Returns a version object represented as a string.
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
