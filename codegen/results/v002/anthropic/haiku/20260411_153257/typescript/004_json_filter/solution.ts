const fs = require('fs');
const path = require('path');

const data = JSON.parse(fs.readFileSync(path.join('input', 'data.json'), 'utf8'));

const result = data
  .filter((r: any) => r.active === true && r.age >= 30)
  .sort((a: any, b: any) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));
