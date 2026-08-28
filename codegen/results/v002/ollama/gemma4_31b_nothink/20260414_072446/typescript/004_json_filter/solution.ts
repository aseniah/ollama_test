import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function solve() {
  const inputPath = path.join(process.cwd(), 'input', 'data.json');
  
  try {
    const rawData = fs.readFileSync(inputPath, 'utf8');
    const records: Record[] = JSON.parse(rawData);

    const filteredAndSorted = records
      .filter(r => r.active === true && r.age >= 30)
      .sort((a, b) => a.name.localeCompare(b.name));

    process.stdout.write(JSON.stringify(filteredAndSorted, null, 2) + '\n');
  } catch (error) {
    // Per instructions, we only output the resulting JSON array.
    // If the file doesn't exist or is invalid, we exit silently or handle as needed.
    process.exit(1);
  }
}

solve();