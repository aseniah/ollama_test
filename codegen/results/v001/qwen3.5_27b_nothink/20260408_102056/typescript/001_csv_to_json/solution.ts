import * as fs from 'fs';
import * as path from 'path';

const filePath = path.join(process.cwd(), 'input', 'data.csv');

try {
  const data = fs.readFileSync(filePath, 'utf-8');
  const lines = data.split('\n');

  // Skip header line
  const rows = lines.slice(1).filter(line => line.trim() !== '');

  const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = rows.map(row => {
    // Split by comma, handling potential edge cases
    const columns = row.split(',');

    // Basic validation to ensure we have at least 4 columns
    if (columns.length < 4) {
      return {
        Name: '',
        Age: 0,
        Email: '',
        Score: 0
      };
    }

    return {
      Name: columns[0].trim(),
      Age: parseInt(columns[1].trim(), 10),
      Email: columns[2].trim(),
      Score: parseFloat(columns[3].trim())
    };
  });

  console.log(JSON.stringify(result));
} catch (error) {
  console.error('Error reading file:', error);
  process.exit(1);
}