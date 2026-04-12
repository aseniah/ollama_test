import { readFileSync } from 'fs';

const rawInput = readFileSync('input/data.json', 'utf-8');
const data: Array<{
  name: string;
  age: number;
  active: boolean;
  score: number;
}>[] = JSON.parse(rawInput);

const filtered = data.filter((record) => record.active && record.age >= 30).sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(filtered, null, 2));