// scripts/uninstall.js
const fs = require('fs');
const path = require('path');

console.log('🧹 Cleaning up archits CLI...');

const binDir = path.join(__dirname, '..', 'bin');

if (fs.existsSync(binDir)) {
  try {
    fs.rmSync(binDir, { recursive: true, force: true });
    console.log('✅ archits CLI cleaned up successfully!');
  } catch (error) {
    console.error('❌ Error during cleanup:', error.message);
  }
} else {
  console.log('✅ Nothing to clean up.');
}