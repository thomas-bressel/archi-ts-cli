
// Layer imports
import { UserRepository } from '../repositories/user.repository';
import User from '../entities/user.entity';

export class UserService {
  private userRepository: UserRepository;

  constructor(userRepository: UserRepository) {
    this.userRepository = userRepository;
  }

  async findAll(): Promise<User[]> {
    return await this.userRepository.findAll();
  }

  
}
