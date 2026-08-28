import * as fs from 'fs';
import * as path from 'path';

const filePath = path.resolve('input/data.csv');
const content = fs.readFileSync(filePath, 'utf-8');

const lines = content.split('\n').filter(line => line.trim() !== '');
const header = lines[0].split(',');
const rows = lines.slice(1);

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (const row of rows) {
    const parts = row.split(',');
    if (parts.length < 4) continue;
    
    const name = parts[0].trim();
    const age = parseInt(parts[1].trim(), 10);
    const email = parts[2].trim();
    const score = parseFloat(parts[3].trim());
    
    result.push({ Name: name, Age: age, Email: email, Score: score });
}

console.log(JSON.stringify(result));