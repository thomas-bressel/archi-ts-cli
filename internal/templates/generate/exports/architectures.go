package exports

import "fmt"

func GetCleanServiceLayerImports(upperName string, lowerName string) string {
	return fmt.Sprintf(`// Layer import
import { %sRepository } from '@repositories/%s.repository';
import  %s  from '@entities/%s.entity';`, upperName, lowerName, upperName, lowerName)
}

func GetLayeredServiceLayerImports(upperName string, lowerName string) string {
	return fmt.Sprintf(`// Layer import
import { %sRepository } from '@repositories/%s.repository';
import { %s } from '@businessmodels/%s.model';`, upperName, lowerName, upperName, lowerName)
}

func GetCleanRepositoryLayerImports(upperName string, lowerName string) string {
	return fmt.Sprintf(`// Layer import
import  %s  from '@entities/%s.entity';`, upperName, lowerName)
}

func GetLayeredRepositoryLayerImports(upperName string, lowerName string) string {
	return fmt.Sprintf(`// Layer import
import { %s } from '@businessmodels/%s.model';`, upperName, lowerName)
}

func GetTestLayeredImports(upperName string, lowerName string) string {
	return fmt.Sprintf(`// Layer import
	import %s from '@businessmodels/%s.model';`, upperName, lowerName)
}

func GetTestCleanImports(upperName string, lowerName string) string {
	return fmt.Sprintf(`// Layer import
		import %s from '@entities/%s.entity';`, upperName, lowerName)
}
func GetTestRepositoryLayeredImports(upperName string, lowerName string) string {
	return fmt.Sprintf(`// Layer import
	import %s from '@datamodels/%s.model';`, upperName, lowerName)
}

func GetTestRepositoryCleanImports(upperName string, lowerName string) string {
	return fmt.Sprintf(`// Layer import
		import %s from '@entities/%s.entity';`, upperName, lowerName)
}
