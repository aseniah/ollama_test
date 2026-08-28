import * as fs from 'fs';
import * as path from 'path';

// Read the input CSV file
const csvPath = path.join(__dirname, 'input', 'input.csv');
const csvContent = fs.readFileSync(csvPath, 'utf-8');

// Parse the CSV
const lines = csvContent.trim().split('\n');
const headers = lines[0].split(',');
const dataRows = lines.slice(1);

// Reference date for age calculation
const referenceDate = new Date('2025-07-01');

const result: any[] = [];

for (const line of dataRows) {
  const values = line.split(',');
  
  // Create a map from header to value
  const row: Record<string, string> = {};
  headers.forEach((header, index) => {
    row[header.trim()] = (values[index] || '').trim();
  });

  // Parse name into first and last
  const fullName = row['Name'];
  const lastSpaceIndex = fullName.lastIndexOf(' ');
  let firstName = fullName;
  let lastName = '';
  if (lastSpaceIndex !== -1) {
    firstName = fullName.slice(0, lastSpaceIndex);
    lastName = fullName.slice(lastSpaceIndex + 1);
  }

  // Parse birthday
  const birthdayStr = row['Birthday'];
  const birthdayParts = birthdayStr.split('/');
  const birthday = new Date(
    parseInt(birthdayParts[2], 10),
    parseInt(birthdayParts[0], 10) - 1, // Month is 0-indexed
    parseInt(birthdayParts[1], 10)
  );

  // Calculate age as of July 1, 2025
  let age = referenceDate.getFullYear() - birthday.getFullYear();
  const monthDiff = referenceDate.getMonth() - birthday.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && referenceDate.getDate() < birthday.getDate())) {
    age--;
  }

  // Build relatives array
  const relatives: any[] = [];
  
  const fatherName = row['Father'];
  if (fatherName && fatherName !== 'null') {
    const fLastSpaceIndex = fatherName.lastIndexOf(' ');
    let fFirstName = fatherName;
    let fLastName = '';
    if (fLastSpaceIndex !== -1) {
      fFirstName = fatherName.slice(0, fLastSpaceIndex);
      fLastName = fatherName.slice(fLastSpaceIndex + 1);
    }
    relatives.push({
      FirstName: fFirstName,
      LastName: fLastName,
      Relationship: "Father"
    });
  }

  const motherName = row['Mother'];
  if (motherName && motherName !== 'null') {
    const mLastSpaceIndex = motherName.lastIndexOf(' ');
    let mFirstName = motherName;
    let mLastName = '';
    if (mLastSpaceIndex !== -1) {
      mFirstName = motherName.slice(0, mLastSpaceIndex);
      mLastName = motherName.slice(mLastSpaceIndex + 1);
    }
    relatives.push({
      FirstName: mFirstName,
      LastName: mLastName,
      Relationship: "Mother"
    });
  }

  const brotherName = row['Brother'];
  if (brotherName && brotherName !== 'null') {
    const bLastSpaceIndex = brotherName.lastIndexOf(' ');
    let bFirstName = brotherName;
    let bLastName = '';
    if (bLastSpaceIndex !== -1) {
      bFirstName = brotherName.slice(0, bLastSpaceIndex);
      bLastName = brotherName.slice(bLastSpaceIndex + 1);
    }
    relatives.push({
      FirstName: bFirstName,
      LastName: bLastName,
      Relationship: "Brother"
    });
  }

  const sisterName = row['Sister'];
  if (sisterName && sisterName !== 'null') {
    const sLastSpaceIndex = sisterName.lastIndexOf(' ');
    let sFirstName = sisterName;
    let sLastName = '';
    if (sLastSpaceIndex !== -1) {
      sFirstName = sisterName.slice(0, sLastSpaceIndex);
      sLastName = sisterName.slice(sLastSpaceIndex + 1);
    }
    relatives.push({
      FirstName: sFirstName,
      LastName: sLastName,
      Relationship: "Sister"
    });
  }

  result.push({
    FirstName: firstName,
    LastName: lastName,
    Birthday: birthdayStr,
    Age: age,
    Relatives: relatives
  });
}

console.log(JSON.stringify(result, null, 2));