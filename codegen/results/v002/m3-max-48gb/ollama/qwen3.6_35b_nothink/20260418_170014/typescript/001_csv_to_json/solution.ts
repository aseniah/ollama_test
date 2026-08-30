import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join('input', 'data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

// Skip header (first line)
const header = lines[0];
const dataLines = lines.slice(1);

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (const line of dataLines) {
  const parts = line.split(',');
  if (parts.length < 4) continue;

  const Name = parts[0].trim();
  const Age = parseInt(parts[1].trim(), 10);
  const Email = parts[2].trim();
  const Score = parseFloat(parts[3].trim());

  result.push({ Name, Age, Email, Score });
}

console.log(JSON.stringify(result, null, 2));