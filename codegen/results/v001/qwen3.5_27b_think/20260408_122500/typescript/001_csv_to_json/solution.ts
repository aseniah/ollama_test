import * as fs from 'fs';

const filePath = process.argv[2] || 'input/data.csv';

const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.split('\n').filter(line => line.trim() !== '');

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',').map(v => v.trim());

    const row = {
        Name: values[0],
        Age: parseInt(values[1], 10),
        Email: values[2],
        Score: parseFloat(values[3])
    };

    result.push(row);
}

console.log(JSON.stringify(result));