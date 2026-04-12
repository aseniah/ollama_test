import * as fs from 'fs';
import * as path from 'path';

// Helper to parse dates in M/D/YYYY or MM/DD/YYYY format to YYYY-MM-DD
function formatDate(dateStr: string): string | null {
  if (!dateStr || dateStr === 'null') return null;
  const parts = dateStr.split('/');
  if (parts.length !== 3) return null;

  // Assume M/D/YYYY or MM/DD/YYYY
  let [month, day, year] = parts.map((p) => parseInt(p.trim(), 10));

  // If month > 12, assume it was actually day and swap (though input suggests standard M/D/Y)
  // Given the input format is clearly Month/Day/Year based on "10/9/1940", we trust that order.
  // However, to be safe against ambiguity if numbers are ambiguous, we stick to M/D/Y as per typical US format for Beatles bios.

  const yearStr = String(year);
  const monthStr = String(month).padStart(2, '0');
  const dayStr = String(day).padStart(2, '0');

  return `${yearStr}-${monthStr}-${dayStr}`;
}

// Calculate age as of July 1, 2025
function calculateAge(birthday: string): number {
  if (!birthday) return 0; // Should not happen based on input but good for safety
  
  const birthDate = new Date(birthday);
  const referenceDate = new Date('2025-07-01');

  let age = referenceDate.getFullYear() - birthDate.getFullYear();
  
  // Check if the birthday has occurred this year before July 1st
  const monthCheck = referenceDate.getMonth() < birthDate.getMonth() || 
                     (referenceDate.getMonth() === birthDate.getMonth() && referenceDate.getDate() < birthDate.getDate());

  if (monthCheck) {
    age--;
  }

  return age;
}

// Parse CSV line handling basic comma separation (no quoted commas expected in this specific input)
function parseLine(line: string): Record<string, string> {
  const headers = [
    "Name", "Birthday", "Died", "Father", "Mother", "Brother", "Sister"
  ];
  
  // Split by comma. Note: The input doesn't seem to have commas inside names like "Lennon, John".
  // It uses "John Winston Lennon". Simple split is sufficient here.
  const values = line.split(',').map(v => v.trim());

  const row: Record<string, string> = {};
  headers.forEach((header, index) => {
    row[header] = values[index] || '';
  });

  return row;
}

function main() {
  const inputPath = path.join(process.cwd(), 'input', 'input.csv');
  
  try {
    const fileContent = fs.readFileSync(inputPath, 'utf8');
    const lines = fileContent.split(/\r?\n/).filter(line => line.trim() !== '');
    
    // Skip header row if present (first line)
    const dataLines = lines.slice(1);
    
    const result: any[] = [];

    for (const line of dataLines) {
      const row = parseLine(line);
      
      // Split full name into First and Last (assuming space separation, taking first word as First)
      // Example: "John Winston Lennon" -> First: John, Last: Lennon
      const nameParts = row['Name'].trim().split(' ');
      const firstName = nameParts[0];
      const lastName = nameParts[nameParts.length - 1];

      const birthdayStr = formatDate(row['Birthday']);
      
      // Calculate Age as of July 1, 2025
      const age = calculateAge(birthdayStr || '');

      // Build Relatives Array
      const relatives: { FirstName: string; LastName: string; Relationship: string }[] = [];

      // Helper to add relative if name is provided and not 'null'
      const addRelative = (fullName: string, relationship: string) => {
        if (fullName && fullName !== 'null') {
          const parts = fullName.trim().split(' ');
          relatives.push({
            FirstName: parts[0],
            LastName: parts.length > 1 ? parts[parts.length - 1] : parts[0],
            Relationship: relationship
          });
        }
      };

      addRelative(row['Father'], 'Father');
      addRelative(row['Mother'], 'Mother');
      addRelative(row['Brother'], 'Brother');
      addRelative(row['Sister'], 'Sister');

      result.push({
        FirstName: firstName,
        LastName: lastName,
        Birthday: birthdayStr || '', // Should be YYYY-MM-DD based on expected format
        Age: age,
        Relatives: relatives
      });
    }

    console.log(JSON.stringify(result));

  } catch (error) {
    // If an error occurs, we might want to output empty array or null, 
    // but the instructions say "Do not output anything other than the JSON array".
    // We'll assume valid input as per task description.
    console.log('[]');
  }
}

main();