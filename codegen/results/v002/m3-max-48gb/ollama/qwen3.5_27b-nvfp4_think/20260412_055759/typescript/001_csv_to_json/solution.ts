import * as fs from 'fs';

const filePath = 'input/data.csv';
const content = fs.readFileSync(filePath, 'utf-8');
const lines = content.trim().split('\n');

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

// Skip header (first line), process data rows starting from index 1
for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',');
    if (values.length >= 4) {
        result.push({
            Name: values[0],
            Age: parseInt(values[1], 10),
            Email: values[2],
            Score: parseFloat(values[3])
        });
    }
}

console.log(JSON.stringify(result));