package files

import (
	"archi-ts-cli/internal/models"
	"archi-ts-cli/internal/templates/create/files/typeorm"
)

// GetTypeORMDataSourceTemplate return template for data-source.ts file
func GetTypeORMDataSourceTemplate(architecture string) string {
	switch architecture {
	case string(models.LayeredArchitecture):
		return typeorm.GetDataSourceTemplate("src/data")
	case string(models.CleanArchitecture):
		return typeorm.GetDataSourceTemplate("src/infrastructure")
	case string(models.HexagonalArchitecture):
		return typeorm.GetDataSourceTemplate("storage")
	default:
		return ""
	}
}

// GetTypeORMCreateDatabaseTemplate return template for create-database.ts file
func GetTypeORMCreateDatabaseTemplate() string {
	return typeorm.GetCreateDatabaseTemplate()
}

// GetHeloperORMScriptTemplate return template for generate-migration.ts file
func GetTypeORMHelperScriptTemplate(architecture string) string {
	switch architecture {
	case string(models.LayeredArchitecture):
		return typeorm.GetHelperScriptTemplate("src/data/database/migrations", "src/data/database/connection")
	case string(models.CleanArchitecture):
		return typeorm.GetHelperScriptTemplate("src/infrastructure/database/migrations", "src/infrastructure/database/config")
	case string(models.HexagonalArchitecture):
		return typeorm.GetHelperScriptTemplate("storage/database/migrations", "src/adapters/secondary/persistence/orm")
	default:
		return ""
	}
}
