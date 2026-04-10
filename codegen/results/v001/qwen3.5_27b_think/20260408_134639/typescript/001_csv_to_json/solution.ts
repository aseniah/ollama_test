import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join('input', 'data.csv');
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n').map(line => line.replace('\r', ''));

// Skip header
const dataLines = lines.slice(1);

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = dataLines
    .filter(line => line.length > 0)
    .map(line => {
        const [name, age, email, score] = line.split(',');
        return {
            Name: name,
            Age: parseInt(age, 10),
            Email: email,
            Score: parseFloat(score)
        };
    });

console.log(JSON.stringify(result));