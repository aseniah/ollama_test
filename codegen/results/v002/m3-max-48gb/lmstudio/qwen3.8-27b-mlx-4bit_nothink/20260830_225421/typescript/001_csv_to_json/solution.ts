I'll write a TypeScript program that reads the CSV file and converts it to JSON.

<tool_call>
<function=code_generation>
<parameter=path>
solution.ts
</parameter>
<parameter=content>
import * as fs from 'fs';
import * as path from 'path';

interface Person {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function main() {
  const filePath = path.join('input', 'data.csv');
  const content = fs.readFileSync(filePath, 'utf-8');
  const lines = content.split('\n').filter(line => line.trim() !== '');

  // Skip the header row
  const dataLines = lines.slice(1);

  const results: Person[] = dataLines.map(line => {
    const parts = line.split(',');
    return {
      Name: parts[0],
      Age: parseInt(parts[1], 10),
      Email: parts[2],
      Score: parseFloat(parts[3]),
    };
  });

  process.stdout.write(JSON.stringify(results, null, 2));
}

main();
</parameter>
</function>
</tool_call>