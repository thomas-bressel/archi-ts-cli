import { Router } from 'express';
import { UserController } from '../controllers/user.controller';

const router = Router();
const userController = new UserController();

// GET /users
router.get('/', userController.getAll.bind(userController));

// GET /users/:id
router.get('/:id', userController.getById.bind(userController));

// POST /users
router.post('/', userController.create.bind(userController));

// PUT /users/:id
router.put('/:id', userController.update.bind(userController));

// DELETE /users/:id
router.delete('/:id', userController.delete.bind(userController));

export { router as UserRoutes };
