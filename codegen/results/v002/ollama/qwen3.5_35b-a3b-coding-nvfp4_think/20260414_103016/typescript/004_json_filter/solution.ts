import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'data.json');
const data: Array<{ name: string; age: number; active: boolean; score: number }> = JSON.parse(fs.readFileSync(filePath, 'utf-8'));

const result = data
    .filter(record => record.active === true && record.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));