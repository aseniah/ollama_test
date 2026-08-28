import * as fs from 'fs';
import * as path from 'path';

interface UserRecord {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function main() {
  const inputPath = path.join('input', 'data.json');

  try {
    // Read the file synchronously
    const rawData = fs.readFileSync(inputPath, 'utf8');
    const data: UserRecord[] = JSON.parse(rawData);

    // Filter: active is true AND age >= 30
    const filtered = data.filter((record) => record.active === true && record.age >= 30);

    // Sort: name ascending
    const sorted = filtered.sort((a, b) => a.name.localeCompare(b.name));

    // Output to stdout as JSON
    process.stdout.write(JSON.stringify(sorted, null, 2) + '\n');
  } catch (error) {
    // The prompt asks to output only the resulting JSON array.
    // In a production environment, we would handle errors more gracefully,
    // but based on the strict output constraints, we avoid printing error details
    // unless the system allows for stderr.
    process.exit(1);
  }
}

main();