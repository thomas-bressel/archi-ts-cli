package layered

func GetTsconfigPaths() string {
	return `{
      "@src/*": ["src/*"],
      "@config/*": ["src/common/config/*"],
      "@constants/*": ["src/common/constants/*"],
      "@errors/*": ["src/common/errors/*"],
      "@logging/*": ["src/common/logging/*"],
      "@utils/*": ["src/common/utils/*"],

      "@controllers/*": ["src/presentation/controllers/*"],
      "@middlewares/*": ["src/presentation/middlewares/*"],
      "@routes/*": ["src/presentation/routes/*"],

      "@services/*": ["src/business/services/*"],
      "@interfaces/*": ["src/business/interfaces/*"],
      "@businessmodels/*": ["src/business/models/*"],

      "@repositories/*": ["src/data/repositories/*"],
      "@datamodels/*": ["src/data/models/*"],
      "@connection/*": ["src/data/database/connection/*"],
      "@migration/*": ["src/data/database/migration/*"],
      "@seeds/*": ["src/data/database/seeds/*"],
      "@storage/*": ["storage/*"],
    }`
}
