import * as fs from 'fs';
import * as path from 'path';

interface OutputPerson {
  FirstName: string;
  LastName: string;
  Birthday: string; // ISO format YYYY-MM-DD
  Age: number;
  Relatives: Array<{
    FirstName: string;
    LastName: string;
    Relationship: string;
  }>;
}

function parseName(fullName: string): { first: string; last: string } {
  const parts = fullName.split(' ');
  // Handle the case where there might be a middle name (like "John Winston Lennon")
  // Based on the expected output, it seems we take the first word as First Name
  // and the last word as Last Name.
  const firstName = parts[0];
  const lastName = parts[parts.length - 1];
  return { first: firstName, last: lastName };
}

function formatDate(dateStr: string): string {
  // Expected input format: M/D/YYYY (e.g., "10/9/1940")
  const [monthStr, dayStr, yearStr] = dateStr.split('/').map(Number);
  
  // Create a Date object. Note: month is 0-indexed in JS Date
  const date = new Date(yearStr, monthStr - 1, dayStr);

  // Format to YYYY-MM-DD
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');

  return `${year}-${month}-${day}`;
}

function calculateAge(birthday: Date, referenceDate: Date): number {
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  
  // Check if the birthday has occurred in the current year yet
  const monthCheck = referenceDate.getMonth() < birthday.getMonth();
  const dayCheck = referenceDate.getMonth() === birthday.getMonth() && 
                    referenceDate.getDate() < birthday.getDate();

  if (monthCheck || dayCheck) {
    age--;
  }

  return age;
}

function main() {
  // Reference date: July 1, 2025
  const referenceDate = new Date(2025, 6, 1); // Month is 0-indexed, so 6 is July

  // Read input file
  const inputFile = path.join(process.cwd(), 'input', 'input.csv');
  const rawContent = fs.readFileSync(inputFile, 'utf-8');
  
  const lines = rawContent.trim().split('\n');
  if (lines.length === 0) {
    console.log('[]');
    return;
  }

  // Parse CSV Header
  // Name,Birthday,Died,Father,Mother,Brother,Sister
  // We will map columns by index based on the provided input structure
  
  const results: OutputPerson[] = [];

  for (let i = 1; i < lines.length; i++) {
    const line = lines[i];
    if (!line.trim()) continue;

    const [fullName, birthdayStr, diedStr, father, mother, brother, sister] = line.split(',');

    // Parse Name
    const { first: firstName, last: lastName } = parseName(fullName);

    // Parse Birthday
    let birthDate: Date | null = null;
    let formattedBirthday: string = "";
    
    if (birthdayStr) {
      try {
        birthDate = new Date(formatDate(birthdayStr));
        formattedBirthday = formatDate(birthdayStr);
      } catch (e) {
        // Fallback or skip if invalid
      }
    }

    // Determine Age
    let age: number | null = null;
    if (birthDate) {
      // If the person has died, we technically shouldn't calculate age as of 2025 
      // in a literal sense for living status, but the prompt asks "Calculate ages as of July 1, 2025".
      // The example shows John Lennon (died 1980) having Age 40. 
      // Wait, 2025 - 1940 = 85. Why is the expected age 40?
      // Looking closely at the provided expected_output.json for John Lennon: "Age": 40.
      // Born Oct 9, 1940. Died Dec 8, 1980. Age at death: 40.
      // The prompt says "Calculate ages as of July 1, 2025".
      // However, the expected output explicitly shows 40 for Lennon.
      // This implies: If 'Died' is present, calculate age at death. If not, calculate as of July 1, 2025.
      
      let endOfLifeDate = referenceDate;
      if (diedStr && diedStr !== 'null') {
        try {
          const dateAtDeath = new Date(formatDate(diedStr));
          // We subtract 1 day from death date to get age *at* death 
          // Or simply calculate age as of the death date?
          // Let's calculate age at the date of death.
          endOfLifeDate = dateAtDeath;
        } catch (e) {
          // Ignore invalid death date
        }
      }
      
      age = calculateAge(birthDate, endOfLifeDate);
    }

    // Parse Relatives
    const relatives: OutputPerson['Relatives'] = [];

    // Helper to parse relative string
    const addRelative = (nameStr: string | null, relationship: string) => {
      if (nameStr && nameStr !== 'null' && nameStr.trim() !== '') {
        const { first, last } = parseName(nameStr);
        relatives.push({
          FirstName: first,
          LastName: last,
          Relationship: relationship
        });
      }
    };

    addRelative(father, "Father");
    addRelative(mother, "Mother");
    addRelative(brother, "Brother");
    addRelative(sister, "Sister");

    if (firstName && formattedBirthday !== "" && age !== null) {
      results.push({
        FirstName: firstName,
        LastName: lastName,
        Birthday: formattedBirthday,
        Age: age,
        Relatives: relatives
      });
    }
  }

  console.log(JSON.stringify(results, null, ' '));
}

main();