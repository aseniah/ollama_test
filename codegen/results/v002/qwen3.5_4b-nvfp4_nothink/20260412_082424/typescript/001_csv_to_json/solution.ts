import * as fs from 'fs';

const filePath = process.argv[2];
if (!filePath) {
  console.error('Usage: tsx script.ts input/data.csv');
  process.exit(1);
}

const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.split('\n').slice(1); // skip header

const result: any[] = [];
for (const line of lines) {
  if (!line.trim()) continue;
  
  const parts = line.split(',');
  const name = parts[0].trim();
  const age = parseInt(parts[1], 10);
  const email = parts[2].trim();
  const score = parseFloat(parts[3]);

  result.push({ Name: name, Age: age, Email: email, Score: score });
}

console.log(JSON.stringify(result));