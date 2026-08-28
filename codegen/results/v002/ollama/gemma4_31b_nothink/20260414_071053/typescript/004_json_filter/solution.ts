import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function main() {
  const filePath = path.join('input', 'data.json');

  try {
    // Read the file content
    const rawData = fs.readFileSync(filePath, 'utf8');
    const data: Record[] = JSON.parse(rawData);

    // Filter: active is true AND age >= 30
    const filtered = data.filter(item => item.active === true && item.age >= 30);

    // Sort: name ascending
    filtered.sort((a, b) => a.name.localeCompare(b.name));

    // Output the result as a JSON array
    process.stdout.write(JSON.stringify(filtered, null, 2) + '\n');
  } catch (error) {
    // Silently exit or handle error as per requirements
    process.exit(1);
  }
}

main();