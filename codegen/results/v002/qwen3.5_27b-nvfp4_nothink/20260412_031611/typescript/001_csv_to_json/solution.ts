import * as fs from 'fs';
import * as path from 'path';

interface DataRecord {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function main() {
  const filePath = 'input/data.csv';

  try {
    const data = fs.readFileSync(filePath, 'utf-8');
    const lines = data.trim().split(/\r?\n/);

    if (lines.length <= 1) {
      console.log('[]');
      return;
    }

    // Skip header line
    const records: DataRecord[] = [];

    for (let i = 1; i < lines.length; i++) {
      const line = lines[i];
      if (!line.trim()) continue;

      const parts = line.split(',');
      if (parts.length < 4) continue;

      records.push({
        Name: parts[0].trim(),
        Age: parseInt(parts[1], 10),
        Email: parts[2].trim(),
        Score: parseFloat(parts[3])
      });
    }

    console.log(JSON.stringify(records));
  } catch (error) {
    // If the file does not exist or cannot be read, output an empty array
    console.log('[]');
  }
}

main();