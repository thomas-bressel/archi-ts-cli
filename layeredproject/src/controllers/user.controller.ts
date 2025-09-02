
// Express importation
import { Request, Response } from 'express';

// Services importation
import { UserService } from '../services/user.service';

// Libraries importation
import { validate } from 'class-validator';
import { plainToInstance } from 'class-transformer';

export class UserController {
  private userService: UserService;

  constructor(userService: UserService) {
    this.userService = userService;
  }

  /**
   * Get the list of all user
   * @param req 
   * @param res 
   * 
   * @returns Promise<Response> - Express response object with appropriate status code and data
   * @throws {500} When an internal server error occurs during retrieval
   */
  public async getAll(req: Request, res: Response): Promise<Response> {
    try {
      const users = await this.userService.findAll();
      if (!users) throw new Error("Empty result");
      return res.status(200).json({ success: true, data: users });
    } catch (error) {
      return res.status(500).json(error instanceof Error ? error.message : "Internal server error");
    }
  }
}
