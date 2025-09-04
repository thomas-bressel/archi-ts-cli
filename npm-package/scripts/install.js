// scripts/install.js
const https = require('https');
const fs = require('fs');
const path = require('path');
const os = require('os');

const platform = os.platform();
const arch = os.arch();

// Map des platforms
const platformMap = {
  'linux': 'linux',
  'darwin': 'darwin',  
  'win32': 'windows'
};

const archMap = {
  'x64': 'amd64',
  'arm64': 'arm64'
};

const goos = platformMap[platform];
const goarch = archMap[arch];

if (!goos || !goarch) {
  console.error(`❌ Platform ${platform}/${arch} not supported`);
  console.error('Supported platforms: linux/amd64, linux/arm64, darwin/amd64, darwin/arm64, windows/amd64');
  process.exit(1);
}

const version = require('../package.json').version;
const binaryName = platform === 'win32' ? 'archits.exe' : 'archits';
const downloadUrl = `https://github.com/thomas-bressel/archi-ts-cli/releases/download/v${version}/archits-${goos}-${goarch}${platform === 'win32' ? '.exe' : ''}`;

console.log(`🚀 Installing archits CLI v${version} for ${platform}/${arch}...`);

// Créer le dossier bin
const binDir = path.join(__dirname, '..', 'bin');
if (!fs.existsSync(binDir)) {
  fs.mkdirSync(binDir, { recursive: true });
}

const binaryPath = path.join(binDir, binaryName);

// Fonction de téléchargement
function downloadFile(url, destination) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(destination);
    
    const request = https.get(url, (response) => {
      // Suivre les redirections
      if (response.statusCode === 302 || response.statusCode === 301) {
        return downloadFile(response.headers.location, destination)
          .then(resolve)
          .catch(reject);
      }
      
      if (response.statusCode !== 200) {
        reject(new Error(`Failed to download: HTTP ${response.statusCode}`));
        return;
      }
      
      const totalBytes = parseInt(response.headers['content-length'], 10);
      let downloadedBytes = 0;
      
      response.on('data', (chunk) => {
        downloadedBytes += chunk.length;
        if (totalBytes) {
          const percent = Math.round((downloadedBytes / totalBytes) * 100);
          process.stdout.write(`\r📦 Downloading... ${percent}%`);
        }
      });
      
      response.pipe(file);
      
      file.on('finish', () => {
        file.close();
        console.log('\n✅ Download completed!');
        resolve();
      });
    });
    
    request.on('error', (error) => {
      fs.unlink(destination, () => {}); // Supprimer le fichier partiel
      reject(error);
    });
    
    file.on('error', (error) => {
      fs.unlink(destination, () => {}); // Supprimer le fichier partiel
      reject(error);
    });
  });
}

// Installation
downloadFile(downloadUrl, binaryPath)
  .then(() => {
    // Rendre exécutable (Unix seulement)
    if (platform !== 'win32') {
      fs.chmodSync(binaryPath, '755');
    }
    
    console.log('✅ archits CLI installed successfully!');
    console.log('');
    console.log('Usage:');
    console.log('  npx archits create          # Create new project');
    console.log('  npx archits generate entity # Generate entity files');
    console.log('  npx archits --help          # Show help');
    console.log('');
    console.log('Or install globally:');
    console.log('  npm install -g archits-cli');
    console.log('  archits --help');
  })
  .catch((error) => {
    console.error('❌ Failed to install archits CLI:', error.message);
    console.error('');
    console.error('Please check:');
    console.error('- Your internet connection');
    console.error('- GitHub releases availability');
    console.error('- File permissions');
    process.exit(1);
  });