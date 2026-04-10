import * as fs from 'fs';
import * as path from 'path';

function main() {
  const filePath = path.join('input', 'data.csv');
  
  let content: string;
  try {
    content = fs.readFileSync(filePath, 'utf-8');
  } catch (err) {
    if (err instanceof Error && 'code' in err && (err as any).code === 'ENOENT') {
      console.error('Error: File not found.');
      process.exit(1);
    }
    throw err;
  }

  const lines = content.split('\n');
  const result: any[] = [];

  // Skip header (index 0), iterate from 1
  for (let i = 1; i < lines.length; i++) {
    const line = lines[i].trim();
    if (!line) continue;

    // Handle potential quoted fields with commas by using a simple CSV parser
    // This handles basic cases: "field", field, "field, with comma"
    const fields: string[] = [];
    let currentField = '';
    let inQuotes = false;

    for (let j = 0; j < line.length; j++) {
      const char = line[j];
      if (char === '"') {
        inQuotes = !inQuotes;
      } else if (char === ',' && !inQuotes) {
        fields.push(currentField.replace(/^"|"$/g, '').trim());
        currentField = '';
      } else {
        currentField += char;
      }
    }
    fields.push(currentField.replace(/^"|"$/g, '').trim());

    if (fields.length >= 4) {
      const [name, ageStr, email, scoreStr] = fields;
      
      const age = parseInt(ageStr, 10);
      const score = parseFloat(scoreStr);

      // Only add if parsing was successful
      if (!isNaN(age) && !isNaN(score)) {
        result.push({
          Name: name,
          Age: age,
          Email: email,
          Score: score
        });
      }
    }
  }

  console.log(JSON.stringify(result));
}

main();