package software

import "github.com/jpeizer/Vectorworks-Utility-Refresh/internal/utils"

// HomeDirectory : Home Directory based on the running operating system.
var HomeDirectory = utils.GetHomeDirectory()

// Software is all information that makes up a supported piece of software
// Name provides the SoftwareName
// Installations is a slice of Installation
type Software struct {
	Name          SoftwareName
	Installations []Installation
}

// SoftwareName illustrates all Software Names
type SoftwareName = string

// All possible SoftwareName (s)
const (
	SoftwareVectorworks   SoftwareName = "Vectorworks"
	SoftwareVision        SoftwareName = "Vision"
	SoftwareCloudServices SoftwareName = "VCS"
)

// AllActiveSoftwareNames is used to turn on and off software to track
// slice of SoftwareName
var AllActiveSoftwareNames = []SoftwareName{
	SoftwareVectorworks,
	SoftwareVision,
	SoftwareCloudServices,
}

func (s Software) String() string {
	return s.Name
}

func (s Software) ClearRMCache() {
	//for _, installation := range s.Installations {
	//	installation.ClearRMCache()
	//}
}