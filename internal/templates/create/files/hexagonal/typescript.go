package hexagonal

func GetTsconfigPaths() string {
	return `{
      "@src/*": ["src/*"],
      "@entities/*": ["src/domain/entities/*"], 
      "@objects/*": ["src/domain/value-objects/*"], 
      "@events/*": ["src/domain/events/*"],
      "@exceptions/*": ["src/domain/exceptions/*"],
      "@usecases/*": ["src/application/use-cases/*"], 
      "@services/*": ["src/application/use-cases/*"],
      "@ports/*": ["src/application/ports/*"], 
      "@dtos/*": ["src/application/dtos/*"],
      "@http/*": ["src/adapters/primary/http/*"],
      "@controllers/*": ["src/adapters/primary/http/controllers/*"], 
      "@routes/*": ["src/adapters/primary/http/routes/*"],  
      "@middlewares/*": ["src/adapters/primary/http/middlewares/*"],
      "@cli/*": ["src/adapters/primary/cli/*"],
      "@orm/*": ["src/adapters/secondary/persistence/orm/*"], 
      "@models/*": ["src/adapters/secondary/persistence/models/*"], 
      "@repositories/*": ["src/adapters/secondary/persistence/repositories/*"],
      "@email/*": ["src/adapters/secondary/email/*"],
      "@cache/*": ["src/adapters/secondary/cache/*"],
      "@config/*": ["src/config/*"],  
      "@storage/*": ["storage/*"],
    }`
}
