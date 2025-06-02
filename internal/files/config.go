package files

/*
GetEnvTemplate return the template for a .env file
*/
func GetEnvTemplate() string {
	return `
# Environment
NODE_ENV=development

# Server
LISTEN_PORT="3000"
SERVER_NAME="ArchiTS API"
VERSION="1.0.0"

# Database
DB_HOST="localhost"
DB_PORT="3306"
DB_NAME="archi_db"
DB_USER="root"
DB_PASSWORD="my-super-password"
DB_CONNEXION_LIMIT="100"

# Redis
REIDS_PORT="6379"
REDIS_HOST="localhost"
REDIS_PASSWORD="my-super-password"
REDIS_EXPIRES_IN="3600"

# JWT
WT_SECRET_KEY="your-secret-key"
JWT_REFRESH_SECRET_KEY="your-refresh-secret-key"
JWT_DURATION="2h"
JWT_REFRESH_TOKEN_TIME="20h"

# CORS
CORS_ALLOWED_ORIGINS="http://localhost:3000"
CORS_CREDENTIALS="true"

# SMTP
MAIL_HOST="mail.domain.fr"
MAIL_PORT="465"
MAIL_SECURE="true"
MAIL_AUTH_USER="exemple@domain.fr"
MAIL_AUTH_PASSWORD="my-super-password"

`
}
