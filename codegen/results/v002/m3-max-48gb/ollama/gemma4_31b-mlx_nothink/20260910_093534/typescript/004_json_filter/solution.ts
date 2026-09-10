import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function solve() {
  try {
    const filePath = path.join('input', 'data.json');
    const fileContent = fs.readFileSync(filePath, 'utf8');
    const data: Record[] = JSON.parse(fileContent);

    // Filter: active is true AND age >= 30
    const filtered = data.filter(item => item.active === true && item.age >= 30);

    // Sort by name ascending
    filtered.sort((a, b) => a.name.localeCompare(b.name));

    process.stdout.write(JSON.stringify(filtered, null, 2) + '\n');
  } catch (error) {
    // The prompt specifies not to output anything else, 
    // but for robustness in a real scenario, you'd handle errors.
    process.exit(1);
  }
}

solve();