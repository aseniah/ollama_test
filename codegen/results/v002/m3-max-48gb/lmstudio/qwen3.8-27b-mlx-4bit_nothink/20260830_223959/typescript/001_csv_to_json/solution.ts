import * as fs from 'fs';

const csvContent = fs.readFileSync('input/data.csv', 'utf-8');
const lines = csvContent.split('\n').filter(line => line.trim() !== '');

if (lines.length === 0) {
    console.log('[]');
    process.exit(0);
}

const header = lines[0].split(',').map(h => h.trim());
const rows = lines.slice(1).map(line => {
    const values = line.split(',').map(v => v.trim());
    const obj: Record<string, string | number> = {};
    for (let i = 0; i < header.length; i++) {
        const key = header[i];
        const value = values[i] ?? '';
        if (key === 'Age') {
            obj[key] = parseInt(value, 10);
        } else if (key === 'Score') {
            obj[key] = parseFloat(value);
        } else {
            obj[key] = value;
        }
    }
    return {
        Name: String(obj['Name'] ?? ''),
        Age: Number(obj['Age'] ?? 0),
        Email: String(obj['Email'] ?? ''),
        Score: Number(obj['Score'] ?? 0),
    };
});

console.log(JSON.stringify(rows));