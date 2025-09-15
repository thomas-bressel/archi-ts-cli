package models

// EntityPaths contains all paths for generated files INCLUDING TESTS
type EntityPaths struct {
	Model              string
	Entity             string
	Controller         string
	Service            string
	Repository         string
	RawModel           string
	Route              string
	ControllerUnitTest string
	ServiceUnitTest    string
	RepositoryUnitTest string
	ApiIntegrationTest string
}

// EntityConfig configuration to generate an entity
type EntityConfig struct {
	Name    string
	Express bool
	Orm     Orm
}
