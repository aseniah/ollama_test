import { readFileSync } from 'fs';

// Read CSV file
const csvContent = readFileSync('input/input.csv', 'utf8');
const lines = csvContent.trim().split('\n');

// Helper function to parse dates (assume YYYY-MM-DD format)
function parseDate(dateStr: string): Date {
  const [year, month, day] = dateStr.split('-').map(Number);
  return new Date(year, month - 1, day);
}

// Calculate age as of July 1, 2025
const targetDate = new Date(2025, 6, 1); // July is month index 6

function calculateAge(dateStr: string): number {
  const birthDate = parseDate(dateStr);
  let years = targetDate.getFullYear() - birthDate.getFullYear();
  const targetMonth = targetDate.getMonth() + 1;
  const birthMonth = birthDate.getMonth() + 1;

  // Adjust if birthday hasn't occurred yet in the current year
  if (targetMonth < birthMonth || (targetMonth === birthMonth && targetDate.getDate() < birthDate.getDate())) {
    years -= 1;
  }
  return years;
}

// Parse CSV: assume first line is header, rest are rows. 
// We'll flatten the structure to match expected format if needed, but for now we follow a safe assumption.
// Let's parse assuming first row is header and data starts from second row.
const lines = csvContent.trim().split('\n');

if (lines.length < 2) {
  console.log('[]'); // Empty array if only header or no data
  process.exit(0);
}

const headers = lines[0].split(',');
const rowCount = lines.length - 1;

// Build rows using index arrays, but without knowing schema we assume simple mapping.
// However, this approach will break down if the format is complex (nested). 
// Let's assume first line is header and rest are data.
const columns = headers.map((h) => h.trim().split(',')); // Split by commas might be needed for nested

const jsonRows: any[] = [];

for (let i = 1; i < lines.length; i++) {
  const rowStr = lines[i];
  
  // Assuming simple CSV without quotes for safety, but this is simplistic.
  // We split each field by comma and trim whitespace.
  const fields: string[] = [];
  let currentField = '';
  let inQuotes = false;

  // Manual parsing with quote handling
  for (const char of rowStr) {
    if (char === '"') {
      inQuotes = !inQuotes;
      fields.push(currentField + char);
    } else if (char === ',') {
      if (!inQuotes) fields.push(currentField.trim());
      currentField = '';
    } else if (char === '\n') {
        // Shouldn't happen here, but safety catch
        if (!inQuotes) fields.push(currentField.trim());
        break;
    } else {
      currentField += char;
    }
  }
  
  // Add last field if present
  if (!inQuotes || !fields.includes('')) {
      fields.push(currentField.trim());
  }

  const row = {
    'id': parseInt(fields[0]) || null,
    'name': fields[1] || null,
    'age': calculateAge(fields[2] || ''),
    'status': fields[3] || ''
  };

  // Assuming schema based on typical expected format: [id, name, date_of_birth, status]
  jsonRows.push(row);
}

console.log(JSON.stringify(jsonRows, null, 2));