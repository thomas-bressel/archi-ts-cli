
// Layer importations
import User from '../entities/user.entity';

export class UserRepository {
  // TODO: Add your database connection/ORM here
  // private userQueries: UserQueries(); if you have a query class

  constructor() {
    // Initialize your database connection/ORM here
    // this.userQueries = new UserQueries(); if you have a query class
  }
  
  async findAll(): Promise<User[]> {
    // TODO: Implement database query
    return [];
  }
}
