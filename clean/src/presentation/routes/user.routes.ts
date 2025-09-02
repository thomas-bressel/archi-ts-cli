// Express importation
import express, {Request, Response} from "express";

// Layers importation
import { UserController }  from '@controllers/user.controller'; 
import { UserService } from '@services/user.service';
import { UserRepository } from "@repositories/user.repository";


const router = express.Router();

const userRepository = new UserRepository()
const userService = new UserService(userRepository)
const userController = new UserController(userService);

// GET /api/v1/admin/users
router.get("/api/v1/admin/users", async (req: Request, res: Response) => { 
    userController.getAll(req, res)
});


export default router;
