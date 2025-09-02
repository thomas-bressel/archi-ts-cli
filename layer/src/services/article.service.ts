
// Layer imports
import { ArticleRepository } from '../repositories/article.repository';
import Article from '../entities/article.entity';

export class ArticleService {
  private articleRepository: ArticleRepository;

  constructor(articleRepository: ArticleRepository) {
    this.articleRepository = articleRepository;
  }

  async findAll(): Promise<Article[]> {
    return await this.articleRepository.findAll();
  }

  
}
