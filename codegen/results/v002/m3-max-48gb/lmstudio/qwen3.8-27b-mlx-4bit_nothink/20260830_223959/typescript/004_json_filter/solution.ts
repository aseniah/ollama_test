I'll write a TypeScript program that reads the JSON file, filters and sorts the records, and outputs the result.

<tool_call>
<function=write_file>
<parameter=path>
solution.ts
</parameter>
<parameter=content>
import * as fs from 'fs';

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const data: Record[] = JSON.parse(fs.readFileSync('input/data.json', 'utf-8'));

const result = data
  .filter(r => r.active === true && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

process.stdout.write(JSON.stringify(result, null, 2) + '\n');

</parameter>
</function>
</tool_call>