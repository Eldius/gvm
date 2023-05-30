package versions

// download page parser selectors
const (
	versionCardsSelector          = `table.downloadtable`
	downloadFileByArchRowSelector = `tbody>tr`

	downloadFileByArchLinkSelector     = `td.filename>a`
	downloadFileByArchOSNameSelector   = `td:nth-child(3)`
	downloadFileByArchArchNameSelector = `td:nth-child(4)`
)

const (
	linuxAmd64ArchName = "Linux-x86-64"
	linuxArm64ArchName = "Linux-ARM64"
)
