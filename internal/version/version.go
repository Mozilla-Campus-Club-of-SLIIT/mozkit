package version

// Override at build time with:
//	-ldflags "-X github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/version.Version=0.1.0
//	          -X github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/version.CodeName=तुर"

var (
	CodeName = "sixSeven"
	Version  = "67.67.67"
)
