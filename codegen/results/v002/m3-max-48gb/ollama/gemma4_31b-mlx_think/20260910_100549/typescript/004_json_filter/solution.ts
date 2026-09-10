import * as fs from 'fs';
import * as path from 'path';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

function main() {
  try {
    const filePath = path.join('input', 'data.json');
    const fileContent = fs.readFileSync(filePath, 'utf8');
    const data: Record[] = JSON.parse(fileContent);

    // Filter: active is true AND age is 30 or older
    const filtered = data.filter(item => item.active === true && item.age >= 30);

    // Sort: name ascending
    filtered.sort((a, b) => a.name.localeCompare(b.name));

    // Output JSON array to stdout
    process.stdout.write(JSON.stringify(filtered, null, 2) + '\n');
  } catch (error) {
    // Silent fail or handle error as per requirements
    // Since the prompt says "Do not output anything else", 
    // we avoid printing error logs to stdout if they occur.
  }
}

main();