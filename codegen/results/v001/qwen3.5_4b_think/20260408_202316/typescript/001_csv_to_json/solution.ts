import { readFileSync } from 'fs';

const csvContent = readFileSync('input/data.csv', 'utf-8');
const lines = csvContent.trim().split('\n');

const rows = lines.slice(1);

const result = rows.map(line => {
    const [name, age, email, score] = line.split(',');
    return {
        Name: name.trim(),
        Age: parseInt(age.trim()),
        Email: email.trim(),
        Score: parseFloat(score.trim())
    };
});

console.log(JSON.stringify(result));