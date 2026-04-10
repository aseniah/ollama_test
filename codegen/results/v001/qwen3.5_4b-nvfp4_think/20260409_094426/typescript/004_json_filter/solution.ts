import { readFileSync } from 'fs';

const filePath = 'input/data.json';
let data: any[] = [];

try {
  const content = readFileSync(filePath, 'utf-8');
  data = JSON.parse(content);
} catch {
  process.exitCode = 1;
  return;
}

const filtered = data.filter((record) => {
  return record.active === true && record.age >= 30;
});

filtered.sort((a, b) => a.name.localeCompare(b.name));

process.stdout.write(JSON.stringify(filtered) + '\n');