package clean

func GetJestConfigPaths() string {
	return `{
      "^@src/(.*)$": "<rootDir>/src/$1",
      "^@entities/(.*)$": "<rootDir>/src/domain/entities/$1",
      "^@errors/(.*)$": "<rootDir>/src/domain/errors/$1",
      "^@usecases/(.*)$": "<rootDir>/src/application/use-cases/$1",
      "^@services/(.*)$": "<rootDir>/src/application/use-cases/$1",
      "^@interfaces/(.*)$": "<rootDir>/src/application/interfaces/$1",
      "^@dtos/(.*)$": "<rootDir>/src/application/dtos/$1",
      "^@controllers/(.*)$": "<rootDir>/src/presentation/controllers/$1",
      "^@routes/(.*)$": "<rootDir>/src/presentation/routes/$1",
      "^@middlewares/(.*)$": "<rootDir>/src/presentation/middlewares/$1",
      "^@validators/(.*)$": "<rootDir>/src/presentation/validators/$1",
      "^@repositories/(.*)$": "<rootDir>/src/infrastructure/repositories/$1",
      "^@datamodels/(.*)$": "<rootDir>/src/infrastructure/models/$1",
      "^@cache/(.*)$": "<rootDir>/src/infrastructure/cache/$1",
      "^@email/(.*)$": "<rootDir>/src/infrastructure/email/$1",
      "^@mappers/(.*)$": "<rootDir>/src/infrastructure/mappers/$1",
      "^@config/(.*)$": "<rootDir>/src/infrastructure/database/config/$1",
      "^@utils/(.*)$": "<rootDir>/src/shared/utils/$1",
      "^@storage/(.*)$": "<rootDir>/storage/$1",
    }`
}
