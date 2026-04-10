import { readFileSync } from 'fs';

const csvContent = readFileSync('input/data.csv', 'utf-8');
const lines = csvContent.split('\n').filter(line => line.trim() !== '');

if (lines.length < 2) {
    console.log(JSON.stringify([]));
    process.exit(0);
}

const header = lines[0].split(',').map(h => h.trim());
const data = lines.slice(1).map(row => {
    const values = row.split(',').map(v => v.trim());

    return {
        Name: values[0],
        Age: parseInt(values[1]) || 0,
        Email: values[2],
        Score: parseFloat(values[3]) || 0.0
    };
});

console.log(JSON.stringify(data));