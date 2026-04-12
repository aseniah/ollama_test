import * as fs from 'fs';
import * as path from 'path';

/**
 * This script reads a CSV file from 'input/data.csv', parses its content,
 * and outputs a JSON array of objects to stdout.
 * 
 * Fields:
 * - Name (string)
 * - Age (integer)
 * - Email (string)
 * - Score (float)
 */

interface UserRecord {
  Name: string;
  Age: number;
  Email: string;
  Score: number;
}

function parseCSV(filePath: string): UserRecord[] {
  if (!fs.existsSync(filePath)) {
    throw new Error(`File not found: ${filePath}`);
  }

  const content = fs.readFileSync(filePath, 'utf-8');
  const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');

  if (lines.length <= 1) {
    return [];
  }

  // The first line is the header: Name,Age,Email,Score
  // We skip the first line and map the rest.
  const dataLines = lines.slice(1);

  return dataLines.map(line => {
    const columns = line.split(',');

    if (columns.length < 4) {
      throw new Error(`Malformed CSV row: ${line}`);
    }

    const name = columns[0].trim();
    const age = parseInt(columns[1].trim(), 10);
    const email = columns[2].trim();
    const score = parseFloat(columns[3].trim());

    if (isNaN(age)) {
      throw new Error(`Invalid Age in row: ${line}`);
    }
    if (isNaN(score)) {
      throw new Error(`Invalid Score in row: ${line}`);
    }

    return {
      Name: name,
      Age: age,
      Email: email,
      Score: score
    };
  });
}

function main() {
  // Based on instructions, the file is located at input/data.csv
  // We look for it relative to the current working directory
  const filePath = path.join(process.cwd(), 'input', 'data.csv');

  try {
    const result = parseCSV(filePath);
    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
  } catch (error) {
    // If there is an error, we write it to stderr so stdout remains pure JSON
    if (error instanceof Error) {
      process.stderr.write(`Error: ${error.message}\n`);
    }
    process.exit(1);
  }
}

main();