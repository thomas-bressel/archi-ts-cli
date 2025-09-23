package clean

func GetTsconfigPaths() string {
	return `{
      "@src/*": ["src/*"],
      "@entities/*": ["src/domain/entities/*"],
      "@errors/*": ["src/domain/errors/*"],
      "@usecases/*": ["src/application/use-cases/*"],
      "@services/*": ["src/application/use-cases/*"],
      "@interfaces/*": ["src/application/interfaces/*"],
      "@dtos/*": ["src/application/dtos/*"],
      "@controllers/*": ["src/presentation/controllers/*"],
      "@routes/*": ["src/presentation/routes/*"],
      "@middlewares/*": ["src/presentation/middlewares/*"],
      "@validators/*": ["src/presentation/validators/*"],
      "@datamodels/*": ["src/infrastructure/models/*"],
      "@repositories/*": ["src/infrastructure/repositories/*"],
      "@cache/*": ["src/infrastructure/cache/*"],
      "@email/*": ["src/infrastructure/email/*"],
      "@mappers/*": ["src/infrastructure/mappers/*"],
      "@config/*": ["src/infrastructure/database/config/*"],
      "@utils/*": ["src/shared/utils/*"],
      "@storage/*": ["storage/*"],
    }`
}
