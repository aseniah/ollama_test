// Import necessary Node.js modules
import fs from 'fs';
import path from 'path';

// Function to read the CSV file and convert it to a JSON array
async function readCsvToJson() {
  // Get the file path
  const filePath = path.resolve(__dirname, 'input/data.csv');
  
  try {
    // Read the content of the CSV file
    const csvContent = fs.readFileSync(filePath, 'utf8');
    
    // Split the content by newline to get each row
    const rows = csvContent.trim().split('\n');
    
    // Skip the header
    let jsonArray: any[] = [];
    
    for (const row of rows.slice(1)) {
      // Split each row by comma to get individual fields
      const fields = row.split(',');
      
      // Validate and convert each field to the appropriate type
      const name = fields[0].trim();
      const age = parseInt(fields[1].trim(), 10);
      const email = fields[2].trim();
      const score = parseFloat(fields[3].trim());
      
      // Create an object with the extracted fields
      const obj = {
        Name: name,
        Age: age,
        Email: email,
        Score: score
      };
      
      // Append the object to the JSON array
      jsonArray.push(obj);
    }
    
    // Write the JSON array to stdout
    console.log(JSON.stringify(jsonArray, null, 2));
  } catch (error) {
    console.error('An error occurred:', error);
  }
}

// Execute the function
readCsvToJson();