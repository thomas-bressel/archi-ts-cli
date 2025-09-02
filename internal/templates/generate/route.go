package generate

import (
	"fmt"
	"strings"
)

// GetRouteTemplate génère le template des routes

// GetRouteTemplate génère le template des routes
func GetRouteTemplate(cfg EntityConfig) string {
	lowerName := strings.ToLower(cfg.Name)
	if cfg.Variant == "typescript" {
		return fmt.Sprintf(`import { Router } from 'express';
import { %sController } from '../controllers/%s.controller';

const router = Router();
const %sController = new %sController();

// GET /%s
router.get('/', %sController.getAll.bind(%sController));

// GET /%s/:id
router.get('/:id', %sController.getById.bind(%sController));

// POST /%s
router.post('/', %sController.create.bind(%sController));

// PUT /%s/:id
router.put('/:id', %sController.update.bind(%sController));

// DELETE /%s/:id
router.delete('/:id', %sController.delete.bind(%sController));

export { router as %sRoutes };
`, cfg.Name, lowerName, lowerName, cfg.Name, lowerName+"s", lowerName, lowerName, lowerName+"s", lowerName, lowerName, lowerName+"s", lowerName, lowerName, lowerName+"s", lowerName, lowerName, lowerName+"s", lowerName, lowerName, cfg.Name)
	}

	return fmt.Sprintf(`const { Router } = require('express');
const %sController = require('../controllers/%s.controller');

const router = Router();
const %sControllerInstance = new %sController();

// GET /%s
router.get('/', %sControllerInstance.getAll.bind(%sControllerInstance));

// GET /%s/:id
router.get('/:id', %sControllerInstance.getById.bind(%sControllerInstance));

// POST /%s
router.post('/', %sControllerInstance.create.bind(%sControllerInstance));

// PUT /%s/:id
router.put('/:id', %sControllerInstance.update.bind(%sControllerInstance));

// DELETE /%s/:id
router.delete('/:id', %sControllerInstance.delete.bind(%sControllerInstance));

module.exports = router;
`, cfg.Name, lowerName, lowerName, cfg.Name, lowerName+"s", lowerName, lowerName, lowerName+"s", lowerName, lowerName, lowerName+"s", lowerName, lowerName, lowerName+"s", lowerName, lowerName, lowerName+"s", lowerName, lowerName)
}

// GetQueryTemplate génère le template pour les queries
func GetQueryTemplate(cfg EntityConfig) string {
	tableName := strings.ToLower(cfg.Name) + "s"
	if cfg.Variant == "typescript" {
		return fmt.Sprintf(`// TODO: Replace with your actual database client (e.g., knex, prisma, typeorm)
// import db from '../../config/database';

export class %sQuery {
  private tableName = '%s';

  async selectAll(): Promise<any[]> {
    // TODO: Implement actual database query
    // Example with Knex:
    // return await db(this.tableName).select('*');
    
    // Mock data for now
    return [];
  }

  async selectById(id: number): Promise<any | null> {
    // TODO: Implement actual database query
    // Example with Knex:
    // const result = await db(this.tableName).where({ id }).first();
    // return result || null;
    
    // Mock data for now
    return null;
  }

  async insert(data: any): Promise<any> {
    // TODO: Implement actual database query
    // Example with Knex:
    // const [id] = await db(this.tableName).insert(data);
    // return await this.selectById(id);
    
    // Mock data for now
    return { id: 1, ...data, createdAt: new Date(), updatedAt: new Date() };
  }

  async update(id: number, data: any): Promise<any | null> {
    // TODO: Implement actual database query
    // Example with Knex:
    // await db(this.tableName).where({ id }).update({ ...data, updatedAt: new Date() });
    // return await this.selectById(id);
    
    // Mock data for now
    return { id, ...data, updatedAt: new Date() };
  }

  async delete(id: number): Promise<boolean> {
    // TODO: Implement actual database query
    // Example with Knex:
    // const deleted = await db(this.tableName).where({ id }).delete();
    // return deleted > 0;
    
    // Mock data for now
    return true;
  }
}
`, cfg.Name, tableName)
	}

	return fmt.Sprintf(`// TODO: Replace with your actual database client (e.g., knex, prisma, typeorm)
// const db = require('../../config/database');

class %sQuery {
  constructor() {
    this.tableName = '%s';
  }

  async selectAll() {
    // TODO: Implement actual database query
    // Example with Knex:
    // return await db(this.tableName).select('*');
    
    // Mock data for now
    return [];
  }

  async selectById(id) {
    // TODO: Implement actual database query
    // Example with Knex:
    // const result = await db(this.tableName).where({ id }).first();
    // return result || null;
    
    // Mock data for now
    return null;
  }

  async insert(data) {
    // TODO: Implement actual database query
    // Example with Knex:
    // const [id] = await db(this.tableName).insert(data);
    // return await this.selectById(id);
    
    // Mock data for now
    return { id: 1, ...data, createdAt: new Date(), updatedAt: new Date() };
  }

  async update(id, data) {
    // TODO: Implement actual database query
    // Example with Knex:
    // await db(this.tableName).where({ id }).update({ ...data, updatedAt: new Date() });
    // return await this.selectById(id);
    
    // Mock data for now
    return { id, ...data, updatedAt: new Date() };
  }

  async delete(id) {
    // TODO: Implement actual database query
    // Example with Knex:
    // const deleted = await db(this.tableName).where({ id }).delete();
    // return deleted > 0;
    
    // Mock data for now
    return true;
  }
}
`, cfg.Name, tableName)
}
