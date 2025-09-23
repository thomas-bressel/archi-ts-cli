package files

import "fmt"

// GetEnvTemplate returns the template for a .env file
func GetEnvTemplate(port int) string {
	return fmt.Sprintf(`# API Configuration
NODE_ENV=development
LISTEN_PORT=%d
LISTEN_PORT_TEST=%d
SERVER_NAME="ArchiTS API"
VERSION=1.0.0
API_PREFIX=/api
API_VERSION=v1

# Database Configuration for sqlite or mysql
DB_TYPE=sqlite
SQLITE_FILE=./storage/database/database.sqlite
# DB_TYPE=mysql      
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=my-username
DB_PASSWORD=my-password
DB_NAME=my-database
DB_CONNEXION_LIMIT=100

# MongoDB Configuration
MONGODB_PORT=27017
MONGODB_DATA=my-database
MONGODB_USER=my-username
MONGODB_PASS=my-password
MONGODB_HOST=mongodb

# Redis Configuration
REDIS_PORT=6379
REDIS_HOST=localhost
REDIS_PASSWORD=my-super-password
REDIS_EXPIRES_IN=3600

# TypeORM Configuration
TYPEORM_SYNCHRONIZE=true
TYPEORM_DROP_SCHEMA=false
TYPEORM_LOGGING=false

# JWT Configuration
JWT_SECRET_KEY=your-secret-key
JWT_REFRESH_SECRET_KEY=your-refresh-secret-key
JWT_DURATION=2h
JWT_REFRESH_TOKEN_TIME=20h

# CORS Configuration
CORS_ALLOWED_ORIGINS=http://localhost:8080
CORS_CREDENTIALS=true

# Nodemailer SMTP configuration
MAIL_HOST=mail.domain.fr
MAIL_PORT=465
MAIL_SECURE=true
MAIL_AUTH_USER=exemple@domain.fr
MAIL_AUTH_PASSWORD=my-super-password`, port, port+2000)
}
