import * as fs from 'fs';
import * as path from 'path';

const csvPath = path.join(process.cwd(), 'input', 'data.csv');
const content = fs.readFileSync(csvPath, 'utf-8');
const lines = content.split('\n').filter(line => line.trim() !== '');
const result: { Name: string; Age: number; Email: string; Score: number }[] = [];

for (let i = 1; i < lines.length; i++) {
    const parts = lines[i].split(',');
    if (parts.length === 4) {
        result.push({
            Name: parts[0].trim(),
            Age: parseInt(parts[1].trim(), 10),
            Email: parts[2].trim(),
            Score: parseFloat(parts[3].trim())
        });
    }
}

console.log(JSON.stringify(result));