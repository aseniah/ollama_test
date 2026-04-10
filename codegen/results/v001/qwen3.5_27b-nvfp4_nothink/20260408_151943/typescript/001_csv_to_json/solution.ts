import * as fs from 'fs';
import * as path from 'path';

function main() {
  const filePath = path.resolve('input/data.csv');
  
  try {
    const fileContent = fs.readFileSync(filePath, 'utf-8');
    
    // Split by lines and trim empty ones
    const lines = fileContent.split('\n').filter(line => line.trim().length > 0);
    
    if (lines.length === 0) {
      console.log('[]');
      return;
    }

    // Skip header (first line)
    const dataLines = lines.slice(1);
    const result: any[] = [];

    for (const line of dataLines) {
      // Simple CSV parsing (assuming no commas inside quoted values for this basic implementation)
      // If the file has quoted fields containing commas, a more complex parser would be needed.
      const columns = line.split(',');

      if (columns.length < 4) continue;

      result.push({
        Name: columns[0].trim(),
        Age: parseInt(columns[1].trim(), 10),
        Email: columns[2].trim(),
        Score: parseFloat(columns[3].trim())
      });
    }

    console.log(JSON.stringify(result));
    
  } catch (error) {
    // If file is not found or other error, output empty JSON array as per instructions
    console.log('[]');
  }
}

main();