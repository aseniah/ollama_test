import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const filePath = path.resolve('input', 'data.json');

try {
  const dataRaw = fs.readFileSync(filePath, 'utf-8');
  const records: Record[] = JSON.parse(dataRaw);

  const result = records
    .filter((r) => r.active && r.age >= 30)
    .sort((a, b) => a.name.localeCompare(b.name));

  console.log(JSON.stringify(result));
} catch (error) {
  // If an error occurs, exit with non-zero code but output nothing to stdout
  process.exit(1);
}