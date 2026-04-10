import { readFileSync } from 'fs';

const dataPath = process.argv[2];

if (!dataPath) {
  console.error('Usage: node script.js <input_file>');
  process.exit(1);
}

try {
  const rawData = JSON.parse(readFileSync(dataPath, 'utf8'));
  
  if (!Array.isArray(rawData)) {
    console.log('[]');
    process.exit(0);
  }

  const result = rawData
    .filter(record => record.active === true && record.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

  console.log(JSON.stringify(result, null, 2));
} catch (e) {
  console.error('Error:', e.message);
  process.exit(1);
}