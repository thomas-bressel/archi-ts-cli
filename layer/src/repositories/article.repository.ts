
  // Layer importations
  import  Article  from '../entities/article.entity';

export class ArticleRepository {
  // TODO: Add your database connection/ORM here
  // private articleQueries: ArticleQueries(); if you have a query class

  constructor() {
    // Initialize your database connection/ORM here
    // this.articleQueries = new ArticleQueries(); if you have a query class
  }
  
  async findAll(): Promise<Article[]> {
    // TODO: Implement database query
    return [];
  }

}
