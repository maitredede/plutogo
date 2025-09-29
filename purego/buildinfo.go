package plutopure

var (
	libBuildInfo func() string
)

func BuildInfo() string {
	libInit()

	return libBuildInfo()
}
