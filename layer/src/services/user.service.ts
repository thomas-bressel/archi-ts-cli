
// Layer imports
import { UserRepository } from '@repositories/user.repository';
import User from '@entities/user.entity';

export class UserService {
  private userRepository: UserRepository;

  constructor(userRepository: UserRepository) {
    this.userRepository = userRepository;
  }

  /**
   * Get the list of all users
   * * @returns Promise<User[]> - Object response from the database with raw entity data
   */
  public async findAll(): Promise<User[]> {
    try {
      return await this.userRepository.findAll();
    } catch (error) {
      throw new Error("Failed to fetch users: " + (error instanceof Error ? error.message : String(error)));
    }
  }
}
