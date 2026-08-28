import * as fs from 'fs';

// Read the CSV file
const csvContent = fs.readFileSync('input/data.csv', 'utf-8');

// Parse CSV - split into lines and filter out empty lines
const lines = csvContent.trim().split('\n').filter(line => line.trim() !== '');

// Skip header (first line), process data rows
const dataRows = lines.slice(1);

// Parse each row into objects
const result = dataRows.map(line => {
    const [name, age, email, score] = line.split(',');
    return {
        Name: name.trim(),
        Age: parseInt(age.trim()),
        Email: email.trim(),
        Score: parseFloat(score.trim())
    };
});

// Output valid JSON
console.log(JSON.stringify(result));