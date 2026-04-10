import { readFileSync } from 'fs';
import { join } from 'path';

const data = JSON.parse(readFileSync(join('input', 'data.json'), 'utf-8'));

const result = data
  .filter((r: { name: string; age: number; active: boolean; score: number }) =>
    r.active && r.age >= 30
  )
  .sort((a: { name: string }, b: { name: string }) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));
