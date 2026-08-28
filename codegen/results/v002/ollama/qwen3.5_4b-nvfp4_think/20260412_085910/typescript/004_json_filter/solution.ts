import fs from 'fs';

const inputData = JSON.parse(fs.readFileSync('input/data.json', 'utf8'));

const filteredData = inputData.filter(record => 
  record.active === true && record.age >= 30
);

filteredData.sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filteredData));