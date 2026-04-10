import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'data.csv');

try {
  const content = fs.readFileSync(filePath, 'utf-8');
  const lines = content.split('\n').filter(line => line.trim() !== '');

  if (lines.length === 0) {
    console.log('[]');
    process.exit(0);
  }

  // Skip header (first line)
  const dataLines = lines.slice(1);

  const result: any[] = [];

  for (const line of dataLines) {
    // Split by comma, handling potential whitespace
    const parts = line.split(',').map(p => p.trim());

    if (parts.length < 4) continue;

    const name = parts[0];
    const age = parseInt(parts[1], 10);
    const email = parts[2];
    const score = parseFloat(parts[3]);

    if (isNaN(age) || isNaN(score)) continue;

    result.push({
      Name: name,
      Age: age,
      Email: email,
      Score: score
    });
  }

  console.log(JSON.stringify(result));
} catch (error) {
  // In case of error (e.g., file not found), output empty array or handle gracefully
  // But typically for these tasks, we assume the file exists as per requirements.
  console.log('[]');
}