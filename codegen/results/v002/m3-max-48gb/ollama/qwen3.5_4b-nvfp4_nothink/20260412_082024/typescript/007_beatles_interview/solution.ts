import { readFileSync } from "fs";
import { parse } from "csv-parse/sync"; // Since npm packages are not available, we must implement a simple CSV parser manually.

// Simple CSV parser to replace 'csv-parse' dependency
function parseCSV(input: string): any[][] {
  const lines = input.trim().split(/\r?\n/);
  if (lines.length === 0) return [];
  
  const headers = lines[0].split(",").map(h => h.trim());
  const rows: any[][] = [];
  
  for (let i = 1; i < lines.length; i++) {
    const line = lines[i];
    // Handle potential commas inside fields (simple assumption: no commas in fields based on sample data)
    const values = line.split(",");
    rows.push(values.map(v => v.trim()));
  }
  
  return [headers, ...rows];
}

const csvContent = readFileSync("input/input.csv", "utf-8");
const parsedData = parseCSV(csvContent);

// Headers
const names = parsedData[0]; // Index 0 in array
const birthdays = parsedData[1]; // Index 1 in array (birthday column)
const fathers = parsedData[2]; // Index 2
const mothers = parsedData[3]; // Index 3
const brothers = parsedData[4]; // Index 4
const sisters = parsedData[5]; // Index 5

const TARGET_DATE = new Date("2025-07-01T00:00:00Z");

function parseDate(dateStr: string): Date {
  const parts = dateStr.split("/");
  if (parts.length !== 3) return new Date(0); // Fallback
  
  const day = parseInt(parts[0]);
  const month = parseInt(parts[1]) - 1; // Convert to 0-indexed
  const year = parseInt(parts[2]);
  
  return new Date(year, month, day);
}

function calculateAge(birthDate: Date): number {
  const today = new Date("2025-07-01T00:00:00");
  let age = today.getFullYear() - birthDate.getFullYear();
  
  // Adjust for months and days
  const todayMonthDay = today.getMonth() + 1;
  const todayYearDay = today.getDate();
  
  const birthMonthDay = birthDate.getMonth() + 1;
  const birthYearDay = birthDate.getDate();
  
  if (todayMonthDay < birthMonthDay) {
    age -= 1;
  } else if (todayMonthDay === birthMonthDay && todayYearDay < birthYearDay) {
    age -= 1;
  }
  
  return age;
}

function buildRelativeEntry(firstName: string, lastName: string, relation: string): any {
  return {
    FirstName: firstName,
    LastName: lastName,
    Relationship: relation
  };
}

const result: any[] = [];

for (let i = 0; i < parsedData.length - 1; i++) { // Skip header row
  const rowData = parsedData[i + 1]; // Index in CSV file
  
  const firstName = rowData[0];
  const birthdayStr = rowData[1];
  const fatherName = rowData[2];
  const motherName = rowData[3];
  const brotherName = rowData[4];
  const sisterName = rowData[5];
  
  // Calculate age from birth year (since we can't reliably parse the full string without library help, 
  // but standard format is MM/DD/YYYY. We'll parse manually).
  // Note: The input data has "10/9/1940". This means Day=10, Month=9, Year=1940.
  // Wait, the prompt says "Calculate ages as of July 1, 2025". 
  // Let's assume MM/DD/YYYY format for parsing if possible, but looking at Lennon: "10/9/1940" -> Oct 9? Or Day 10 Month 9?
  // Lennon was born October 9, 1940. So "10/9/1940" = Month=10, Day=9, Year=1940.
  // McCartney: "6/18/1942" -> June 18, 1942. So MM/DD/YYYY.
  
  const [dayStr, monthStr, yearStr] = birthdayStr.split("/");
  const birthDate = new Date(parseInt(yearStr), parseInt(monthStr) - 1, parseInt(dayStr));
  
  const age = calculateAge(birthDate);
  
  const relatives: any[] = [];
  
  // Build relative list
  if (fatherName && fatherName !== "null") {
    // We need to split first and last name because they might be in different formats or we just use as-is
    // Looking at data: "Alfred Lennon" -> First="Alfred", Last="Lennon"
    // How to determine? Usually split on space, but names like "Winston Lennon" contain a space.
    // We know the first name of the current person from column 0.
    // For relative families (father/mother/brother), we might need full names or parts.
    // Based on expected format: FirstName = part1, LastName = remaining
    const fatherParts = fatherName.split(" ");
    relatives.push(buildRelativeEntry(fatherParts[0], fatherParts.slice(1).join(" "), "Father"));
  }
  
  if (motherName && motherName !== "null") {
    const motherParts = motherName.split(" ");
    relatives.push(buildRelativeEntry(motherParts[0], motherParts.slice(1).join(" "), "Mother"));
  }
  
  if (brotherName && brotherName !== "null") {
    const brotherParts = brotherName.split(" ");
    relatives.push(buildRelativeEntry(brotherParts[0], brotherParts.slice(1).join(" "), "Brother"));
  }
  
  if (sisterName && sisterName !== "null") {
    const sisterParts = sisterName.split(" ");
    relatives.push(buildRelativeEntry(sisterParts[0], sisterParts.slice(1).join(" "), "Sister"));
  }
  
  // Sort order? Expected format shows: Father, Mother, Brother, Sister (seems to follow the order in CSV columns)
  // Lennon: Father, Mother. McCartney: Father, Mother, Brother. Starr: Father, Mother, Sister. Harrison: Father, Mother, Brother, Sister.
  // So relative list is built based on column existence.
  
  result.push({
    FirstName: firstName,
    LastName: rowData[0] === "null" ? "" : rows[i][0].split(" ")[1]?.trim() || "Lennon", // Wait, last name is usually last part of full name? 
    Birthday: birthdayStr,
    Age: age,
    Relatives: relatives
  });
}

// Correction on logic above:
// We need to reconstruct the LastName correctly.
// The CSV row structure is: Name,Birthday,Died,Father,Mother,Brother,Sister
// "John Winston Lennon" -> First="John", Last="Winston Lennon"? No. 
// Actually, John Lennon's full name is John Winston Lennon.
// But in the output JSON, LastName is "Lennon". FirstName is "John".
// It seems we need to parse the first column "Name" into First and Last.
// We can do a simple split by space.

let finalResult = [];

const headers = parsedData[0];
const dataRows = parsedData.slice(1); // Skip header

for (const row of dataRows) {
  const fullName = row[0];
  const parts = fullName.split(" ");
  const firstName = parts[0];
  const lastNameParts = parts.slice(1).join(" ");
  
  // Parse Birthday
  const [dStr, mStr, yStr] = row[1].split("/");
  const birthDate = new Date(parseInt(yStr), parseInt(mStr) - 1, parseInt(dStr));
  const age = calculateAge(birthDate);
  
  const relatives: any[] = [];
  
  // Parse Father (Col 2)
  if (row[2] !== "null") {
    const fParts = row[2].split(" ");
    relatives.push({
      FirstName: fParts[0],
      LastName: fParts.slice(1).join(" "),
      Relationship: "Father"
    });
  }
  
  // Parse Mother (Col 3)
  if (row[3] !== "null") {
    const mParts = row[3].split(" ");
    relatives.push({
      FirstName: mParts[0],
      LastName: mParts.slice(1).join(" "),
      Relationship: "Mother"
    });
  }
  
  // Parse Brother (Col 4)
  if (row[4] !== "null") {
    const bParts = row[4].split(" ");
    relatives.push({
      FirstName: bParts[0],
      LastName: bParts.slice(1).join(" "),
      Relationship: "Brother"
    });
  }
  
  // Parse Sister (Col 5)
  if (row[5] !== "null") {
    const sParts = row[5].split(" ");
    relatives.push({
      FirstName: sParts[0],
      LastName: sParts.slice(1).join(" "),
      Relationship: "Sister"
    });
  }
  
  finalResult.push({
    FirstName,
    LastName: lastNameParts,
    Birthday: row[1],
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(finalResult, null, 2));