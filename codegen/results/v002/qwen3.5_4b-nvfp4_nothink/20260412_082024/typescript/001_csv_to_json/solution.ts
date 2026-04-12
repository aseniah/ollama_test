const fs = require('fs');

// Read the CSV file content
const content = fs.readFileSync('input/data.csv', 'utf-8');

// Split by newlines and filter out empty lines if any
const rows = content.trim().split('\n').filter(line => line.length > 0);

if (rows.length === 0) {
  console.log('[]');
  process.exit(0);
}

// Skip the header row and process data rows
const outputRows = rows.slice(1).map(row => {
  const [name, age, email, score] = row.split(',').map(str => str.trim());
  
  return {
    Name: name,
    Age: parseInt(age, 10),
    Email: email,
    Score: parseFloat(score)
  };
});

// Output the result as a JSON array
console.log(JSON.stringify(outputRows, null, 2));