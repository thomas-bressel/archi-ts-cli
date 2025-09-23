package layered

func GetJestConfigPaths() string {
	return `{
    "^@src/(.*)$": "<rootDir>/src/$1",
    "^@config/(.*)$": "<rootDir>/src/common/config/$1",
    "^@constants/(.*)$": "<rootDir>/src/common/constants/$1",
    "^@errors/(.*)$": "<rootDir>/src/common/errors/$1",
    "^@logging/(.*)$": "<rootDir>/src/common/logging/$1",
    "^@utils/(.*)$": "<rootDir>/src/common/utils/$1",
    "^@controllers/(.*)$": "<rootDir>/src/presentation/controllers/$1",
    "^@middlewares/(.*)$": "<rootDir>/src/presentation/middlewares/$1",
    "^@routes/(.*)$": "<rootDir>/src/presentation/routes/$1",
    "^@services/(.*)$": "<rootDir>/src/business/services/$1",
    "^@interfaces/(.*)$": "<rootDir>/src/business/interfaces/$1",
    "^@businessmodels/(.*)$": "<rootDir>/src/business/models/$1",
    "^@repositories/(.*)$": "<rootDir>/src/data/repositories/$1",
    "^@datamodels/(.*)$": "<rootDir>/src/data/models/$1",
    "^@connection/(.*)$": "<rootDir>/src/data/database/connection/$1",
    "^@migration/(.*)$": "<rootDir>/src/data/database/migration/$1",
    "^@seeds/(.*)$": "<rootDir>/src/data/database/seeds/$1",
    "^@storage/(.*)$": "<rootDir>/storage/$1",
	}`
}
