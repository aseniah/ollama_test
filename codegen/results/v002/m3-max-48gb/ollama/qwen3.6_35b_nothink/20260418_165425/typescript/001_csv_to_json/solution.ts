import * as fs from 'fs';
import * as path from 'path';

const filePath = 'input/data.csv';
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

// Skip the header line
const dataLines = lines.slice(1);

const result = dataLines.map(line => {
    const parts = line.split(',');
    return {
        Name: parts[0],
        Age: parseInt(parts[1], 10),
        Email: parts[2],
        Score: parseFloat(parts[3])
    };
});

console.log(JSON.stringify(result, null, 2));