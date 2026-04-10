import * as fs from 'fs';
import * as path from 'path';

const filePath = process.argv[2] || 'input/data.csv';

if (!fs.existsSync(filePath)) {
  throw new Error(`File not found: ${filePath}`);
}

const rawData = fs.readFileSync(filePath, 'utf-8');
const lines = rawData.trim().split('\n');

if (lines.length < 1) {
  console.log('[]');
  process.exit(0);
}

// Skip header (first line), process the rest
const results: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (let i = 1; i < lines.length; i++) {
  const line = lines[i].trim();
  if (!line) continue; // Skip empty lines

  // Split by comma, assuming no commas inside fields
  const parts = line.split(',');

  // Ensure we have enough columns (Name, Age, Email, Score)
  if (parts.length < 4) {
    throw new Error(`Invalid data in line ${i + 1}: expected 4 columns`);
  }

  const name = parts[0].trim();
  const age = parseInt(parts[1].trim(), 10);
  const email = parts[2].trim();
  const score = parseFloat(parts[3].trim());

  if (isNaN(age)) {
    throw new Error(`Invalid Age value in line ${i + 1}`);
  }
  if (isNaN(score)) {
    throw new Error(`Invalid Score value in line ${i + 1}`);
  }

  results.push({
    Name: name,
    Age: age,
    Email: email,
    Score: score
  });
}

console.log(JSON.stringify(results));