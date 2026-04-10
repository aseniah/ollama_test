import * as fs from 'fs';

const inputFile = process.argv[2];

if (!inputFile) {
  console.error('Please provide a path to the input file as an argument.');
  process.exit(1);
}

try {
  const content = fs.readFileSync(inputFile, 'utf-8');
  const lines = content.split(/\r?\n/).filter(line => line.trim() !== '');
  
  if (lines.length === 0) {
    console.error('The file is empty.');
    process.exit(1);
  }

  // Assuming the first line is a header, we skip it for data processing.
  const header = lines[0].split(','); 
  
  const data: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

  // Determine column indices from the header (assuming simple comma separation)
  // Example Header: Name,Age,Email,Score
  
  for (let i = 1; i < lines.length; i++) {
    const values = lines[i].split(',');
    
    if (values.length !== header.length) {
      continue; 
    }

    let name = values[0];
    let email = values[2];
    
    // Handling Age: might be a string "35" or float "35.5", parse as integer or number based on input spec?
    // Spec asks for Age: integer, Score: float. Let's attempt to parse as strings and then convert.
    
    try {
      const ageVal = parseInt(values[1]);
      if (isNaN(ageVal)) throw new Error('Invalid age format');
      let scoreVal = parseFloat(values[3]);
      if (isNaN(scoreVal)) throw new Error('Invalid score format');

      data.push({ Name: name, Age: ageVal, Email: email, Score: scoreVal });
    } catch (err) {
      // If conversion fails for required types, skip or warn (silently skipping here as per common CSV robustness).
      console.error(`Skipping malformed line ${i}:`, values);
      continue; 
    }
  }

  console.log(JSON.stringify(data, null, 2));
} catch (err) {
  console.error('Error reading file or parsing data:', err.message);
  process.exit(1);
}