package folder

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/create/folders/clean"
	"archi-ts-cli/internal/templates/create/folders/hexagonal"
	"archi-ts-cli/internal/templates/create/folders/layered"
)

// GetArchitecture return the list of folders to create
func GetArchitecture(architecture models.Architecture) []string {
	switch architecture {
	case models.LayeredArchitecture:
		return layered.GetDirectories()
	case models.CleanArchitecture:
		return clean.GetDirectories()
	case models.HexagonalArchitecture:
		return hexagonal.GetDirectories()
	default:
		return layered.GetDirectories()
	}
}
