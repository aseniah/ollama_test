import { readFileSync } from 'fs';
import { join } from 'path';

const filePath = process.argv[2] || 'input/data.csv';

const content = readFileSync(filePath, 'utf-8');
const lines = content.split('\n').filter(line => line.trim() !== '');

// Assume first line is header with format: Name,Age,Email,Score
const headerLine = lines.shift();
const headers = headerLine.split(',');

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (const line of lines) {
    const values = line.split(',');
    
    let name = '', age = 0, email = '', score = 0;
    
    for (let i = 0; i < headers.length; i++) {
        if (headers[i] === 'Name') {
            name = values[i];
        } else if (headers[i] === 'Age') {
            age = parseInt(values[i], 10);
        } else if (headers[i] === 'Email') {
            email = values[i];
        } else if (headers[i] === 'Score') {
            score = parseFloat(values[i]);
        }
    }
    
    result.push({ Name: name, Age: age, Email: email, Score: score });
}

console.log(JSON.stringify(result));