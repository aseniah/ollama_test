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
    // Read file content
    const fileContent = fs.readFileSync(filePath, 'utf8');
    const data: Record[] = JSON.parse(fileContent);

    // Filter: active is true AND age >= 30
    const filtered = data.filter(item => item.active === true && item.age >= 30);

    // Sort: name ascending
    filtered.sort((a, b) => a.name.localeCompare(b.name));

    // Output as JSON array
    process.stdout.write(JSON.stringify(filtered, null, 2) + '\n');
  } catch (error) {
    // In a real scenario, handle errors appropriately. 
    // For this task, we assume the file exists and is valid JSON.
    process.exit(1);
  }
}

main();