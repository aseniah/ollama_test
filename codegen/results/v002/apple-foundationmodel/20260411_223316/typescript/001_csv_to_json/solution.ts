import fs from 'fs';
import path from 'path';

// Define the CSV file path
const csvFilePath = path.join(__dirname, 'input', 'data.csv');

// Define the expected JSON output format
const expectedJsonOutput = [
  {
    Name: 'Alice Johnson',
    Age: 34,
    Email: 'alice@example.com',
    Score: 92.5
  },
  {
    Name: 'Bob Smith',
    Age: 28,
    Email: 'bob@example.com',
    Score: 78.0
  },
  {
    Name: 'Carol White',
    Age: 45,
    Email: 'carol@example.com',
    Score: 88.3
  },
  {
    Name: 'David Lee',
    Age: 31,
    Email: 'david@example.com',
    Score: 95.1
  },
  {
    Name: 'Eva Martinez',
    Age: 52,
    Email: 'eva@example.com',
    Score: 71.8
  }
];

// Function to read CSV file and parse into JSON objects
async function parseCsvToJson(filePath: string): Promise<string> {
  try {
    const csvData = fs.readFileSync(filePath, 'utf8');
    const rows = csvData.trim().split('\n');
    
    // Skip the header row
    const dataRows = rows.slice(1);
    const jsonArray: any[] = dataRows.map((row) => {
      const fields = row.split(',');
      return {
        Name: fields[0],
        Age: parseInt(fields[1], 10),
        Email: fields[2],
        Score: parseFloat(fields[3])
      };
    });

    return JSON.stringify(jsonArray, null, 2);
  } catch (error) {
    console.error('Error reading CSV file:', error);
    throw error;
  }
}

// Main execution
if (process.argv.length < 2) {
  console.error('Usage: tsx input/data.csv');
  process.exit(1);
}

const jsonOutput = parseCsvToJson(csvFilePath);
console.log(jsonOutput);