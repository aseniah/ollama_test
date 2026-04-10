import * as fs from "fs";
import { parse } from 'csv-parse/async';

async function main() {
  const filePath = process.argv[2] || 'input/data.csv';
  
  let content: string;
  try {
    content = await fs.promises.readFile(filePath, 'utf-8');
  } catch (err) {
    console.error(`Error reading file: ${err}`);
    return;
  }

  const records = await parse(content, {
    columns: true,
    trim: true,
    skipEmpty_lines: true
  });

  const output: Array<{
    name: string;
    age: number;
    email: string;
    score: number;
  }> = [];

  for (const row of records) {
    if (!row.name || !row.age || !row.email || !row.score) continue; // Skip rows with missing data

    output.push({
      name: row.name.trim(),
      age: Number(row.age),
      email: row.email.trim(),
      score: Number(row.score)
    });
  }

  console.log(JSON.stringify(output, null, 2));
}

main().catch(console.error);