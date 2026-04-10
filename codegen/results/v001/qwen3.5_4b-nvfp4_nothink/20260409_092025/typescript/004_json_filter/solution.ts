import * as fs from 'fs';

const data = JSON.parse(fs.readFileSync('input/data.json', 'utf8'));
const result = data.filter(rec => rec.active === true && rec.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));
console.log(JSON.stringify(result, null, 2));