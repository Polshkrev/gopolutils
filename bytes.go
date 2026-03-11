package gopolutils

// Standardization of a unit byte scaler.
type ByteSize = Size

const (
	Byte ByteSize = 1 << (10 * iota) // Unit size in bytes.
	KB                               // Unit size in kilobytes.
	MB                               // Unit size in megabytes.
	GB                               // Unit size in gigabytes.
	TB                               // Unit size in terabytes.
	PB                               // Unit size in petabytes.
	EB                               // Unit size in exabytes.
)

// Represent a [ByteSize] as a string.
// Returns a suffix string representation of a unit byte scaler.
func ByteSizeToString(size ByteSize) string {
	switch size {
	case KB:
		return "KB"
	case MB:
		return "MB"
	case GB:
		return "GB"
	case TB:
		return "TB"
	case PB:
		return "PB"
	case EB:
		return "EB"
	default:
		return "B"
	}
}
