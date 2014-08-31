package wu

type UpdateType int

const (
	UPDATETYPE_SOFTWARE UpdateType = iota
	UPDATETYPE_DRIVERS
)

func (ut UpdateType) String() string {
	switch ut {
	case UPDATETYPE_SOFTWARE:
		return "Software"
	case UPDATETYPE_DRIVERS:
		return "Drivers"
	default:
		return "Unknown"
	}
}

type AutoDownload int

const (
	AUTODOWNLOAD_STANDARD AutoDownload = iota
	AUTODOWNLOAD_NEVER
	AUTODOWNLOAD_ALWAYS
)

func (ad AutoDownload) String() string {
	switch ad {
	case AUTODOWNLOAD_STANDARD:
		return "Standard"
	case AUTODOWNLOAD_NEVER:
		return "Never"
	case AUTODOWNLOAD_ALWAYS:
		return "Always"
	default:
		return "Unknown"
	}
}

type AutoSelection int

const (
	AUTOSELECTION_STANDARD AutoSelection = iota
	AUTOSELECTION_IFDOWNLOADED
	AUTOSELECTION_NEVER
	AUTOSELECTION_ALWAYS
)

func (as AutoSelection) String() string {
	switch as {
	case AUTOSELECTION_STANDARD:
		return "Standard"
	case AUTOSELECTION_IFDOWNLOADED:
		return "IfDownloaded"
	case AUTOSELECTION_NEVER:
		return "Never"
	case AUTOSELECTION_ALWAYS:
		return "Always"
	default:
		return "Unknown"
	}
}
