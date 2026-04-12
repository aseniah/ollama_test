import * as fs from 'fs';

const filePath = 'input/data.csv';
const data = fs.readFileSync(filePath, 'utf8');
const lines = data.trim().split('\n');

interface Record {
    Name: string;
    Age: number;
    Email: string;
    Score: number;
}

const result: Record[] = [];

for (let i = 1; i < lines.length; i++) {
    const line = lines[i].trim();
    if (!line) continue;

    const parts = line.split(',');
    result.push({
        Name: parts[0],
        Age: parseInt(parts[1], 10),
        Email: parts[2],
        Score: parseFloat(parts[3])
    });
}

console.log(JSON.stringify(result));