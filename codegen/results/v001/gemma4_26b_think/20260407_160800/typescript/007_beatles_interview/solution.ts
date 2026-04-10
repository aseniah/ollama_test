import * as fs from 'fs';
import * as path from 'path';

/**
 * Calculates age as of July 1, 2025.
 * 
 * @param birthDateString - The date string representing the person's birth date.
 * @param targetDate - The reference date for the calculation.
 * @returns The calculated age as an integer.
 */
function calculateAge(birthDateString: string, targetDate: Date): number {
  const birthDate = new Date(birthDateString);
  if (isNaN(birthDate.getTime())) {
    return 0;
  }

  let age = targetDate.getFullYear() - birthDate.getFullYear();
  const monthDiff = targetDate.getMonth() - birthDate.getMonth();
  const dayDiff = targetDate.getDate() - birthDate.getDate();

  // If the birthday hasn't occurred yet in the target year, subtract one year.
  if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
    age--;
  }

  return age;
}

/**
 * Parses a simple CSV string into headers and rows.
 * Note: This implementation handles basic CSVs without complex escaped commas.
 */
function parseCSV(csvContent: string): { headers: string[]; rows: string[][] } {
  const lines = csvContent.split(/\r?\n/).filter((line) => line.trim() !== '');
  if (lines.length === 0) {
    return { headers: [], rows: [] };
  }

  const headers = lines[0]
    .split(',')
    .map((h) => h.trim().replace(/^"|"$/g, ''));

  const rows = lines.slice(1).map((line) =>
    line.split(',').map((cell) => cell.trim().replace(/^"|"$/g, ''))
  );

  return { headers, rows };
}

/**
 * Main function to execute the transformation logic.
 */
function main() {
  const csvFilePath = path.join('input', 'input.csv');
  const expectedFormatPath = path.join('input', 'expected_format.json');
  const targetDate = new Date('2025-07-01');

  // Ensure input file exists
  if (!fs.existsSync(csvFilePath)) {
    return;
  }

  const csvContent = fs.readFileSync(csvFilePath, 'utf-8');
  const { headers, rows } = parseCSV(csvContent);

  // Load the expected format to infer the required output structure and keys
  let expectedTemplate: any[] = [];
  if (fs.existsSync(expectedFormatPath)) {
    try {
      expectedTemplate = JSON.parse(fs.readFileSync(expectedFormatPath, 'utf-8'));
    } catch (err) {
      // If the template is unreadable, we default to empty to avoid crashing
      expectedTemplate = [];
    }
  }

  // If no template is found, we cannot infer transformation rules. 
  // We'll attempt to output the CSV rows as simple objects if the template is empty.
  if (expectedTemplate.length === 0) {
    const fallbackResult = rows.map((row) => {
      const obj: Record<string, string> = {};
      headers.forEach((header, index) => {
        obj[header] = row[index] || '';
      });
      return obj;
    });
    process.stdout.write(JSON.stringify(fallbackResult, null, 2) + '\n');
    return;
  }

  // The keys of the first object in the template represent the desired output keys.
  const targetKeys = Object.keys(expectedTemplate[0]);

  /**
   * Mapping dictionary: targetKey -> sourceCSVHeader
   */
  const keyMapping: Record<string, string> = {};

  targetKeys.forEach((targetKey) => {
    const lowerTarget = targetKey.toLowerCase();

    // 1. Try exact match (case-insensitive)
    let matchedHeader = headers.find((h) => h.toLowerCase() === lowerTarget);

    if (!matchedHeader) {
      // 2. Heuristic: If the target key is 'age', look for columns containing 'date', 'birth', or 'dob'
      if (lowerTarget === 'age') {
        matchedHeader = headers.find((h) =>
          h.toLowerCase().includes('date') ||
          h.toLowerCase().includes('birth') ||
          h.toLowerCase().includes('dob')
        );
      } 
      // 3. Heuristic: If the target key is 'name', look for 'fullname' or similar
      else if (lower
      ) {
        matchedHeader = headers.find((h) => h.toLowerCase().includes('name'));
      }
    }

    if (matchedHeader) {
      keyMapping[targetKey] = matchedHeader;
    }
  });

  // Transform each CSV row into the expected JSON object format
  const transformedData = rows.map((row) => {
    const transformedObj: Record<string, any> = {};

    targetKeys.forEach((targetKey) => {
      const sourceHeader = keyMapping[targetKey];

      if (sourceHeader) {
        const colIndex = headers.indexOf(sourceHeader);
        const rawValue = row[colIndex];

        if (targetKey.toLowerCase() === 'age' && rawValue) {
          // Apply specific age calculation rule
          transformedObj[targetKey] = calculateAge(rawValue, targetDate);
        } else {
          // Direct assignment
          transformedObj[targetKey] = rawValue;
        }
      }
    });

    return transformedObj;
  });

  // Output the final JSON array to stdout
  process.stdout.write(JSON.stringify(transformedData, null, 2) + '\n');
}

// Execute the application
main();