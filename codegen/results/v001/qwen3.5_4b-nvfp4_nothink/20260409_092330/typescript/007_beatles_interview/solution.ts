import { readFileSync } from 'fs';
import { writeFileSync } from 'fs';
import { resolve } from 'path';

// Define paths
const inputDir = resolve(import.meta.dirname, 'input');
const inputFile = resolve(inputDir, 'input.csv');
const outputFormatFile = resolve(inputDir, 'expected_format.json');

function parseDate(dateStr: string): Date {
  const d = new Date(dateStr);
  if (isNaN(d.getTime())) {
    console.error(`Invalid date format in: ${dateStr}`);
    process.exit(1);
  }
  return d;
}

function calculateAges(dateStrs: string[]): number[] {
  // Convert all dates to JS Date objects.
  const dates = dateStrs.map(parseDate);
  
  // Define the fixed reference date for age calculation
  const REFERENCE_DATE = new Date('2025-07-01');

  // Calculate ages as of Reference Date
  const ages: number[] = [];
  
  for (let i = 0; i < dates.length; i++) {
    const birthDate = dates[i];
    const ageYears = REFERENCE_DATE.getFullYear() - birthDate.getFullYear();
    
    // If the person's birthday has already passed in Reference Year, it's their age.
    // Otherwise, they are currently (Age - 1).
    if (REFERENCES_DATE >= Date) {
      ages[i] = ageYears;
    } else {
      ages[i] = ageYears - 1;
    }
  }

  return ages;
}