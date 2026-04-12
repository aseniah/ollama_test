import fs from 'fs';

// Read the input file
const inputFile = process.argv[2] || 'input/data.json';
const rawData = fs.readFileSync(inputFile, 'utf8');
const records: any[] = JSON.parse(rawData);

// Filter and sort
const filteredRecords = records.filter(record => record.active && record.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

// Output to stdout
console.log(JSON.stringify(filteredRecords, null, 2));