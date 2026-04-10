import * as fs from 'fs';
import * as path from 'path';

const inputCsvPath = path.join(__dirname, 'input', 'input.csv');
const expectedFormatPath = path.join(__dirname, 'input', 'expected_format.json');

// Read the CSV file
const csvContent = fs.readFileSync(inputCsvPath, 'utf-8');

// Parse the CSV content into an array of objects
function parseCsv(csv: string): { [key: string]: string }[] {
  const lines = csv.split('\n');
  const header = lines[0].split(',');
  const data = [];

  for (let i = 1; i < lines.length; i++) {
    if (lines[i]) {
      const obj: { [key: string]: string } = {};
      lines[i].split(',').forEach((value, index) => {
        obj[header[index]] = value;
      });
      data.push(obj);
    }
  }

  return data;
}

const csvData = parseCsv(csvContent);

// Read the expected format
const expectedFormat = JSON.parse(fs.readFileSync(expectedFormatPath, 'utf-8'));

// Transform the data according to the expected format
function transformData(data: { [key: string]: string }[], format: any[]): any[] {
  return data.map(item => {
    const transformedItem: { [key: string]: any } = {};

    for (const field of format) {
      if (field.type === 'string') {
        transformedItem[field.name] = item[field.source];
      } else if (field.type === 'number') {
        if (field.operation === 'increment' && field.amount !== undefined) {
          transformedItem[field.name] = parseInt(item[field.source], 10) + field.amount;
        } else if (field.operation === 'decrement' && field.amount !== undefined) {
          transformedItem[field.name] = parseInt(item[field.source], 10) - field.amount;
        }
      } else if (field.type === 'date') {
        const date = new Date(item[field.source]);
        date.setFullYear(date.getFullYear() + 25);
        transformedItem[field.name] = date.toISOString().split('T')[0];
      } else if (field.type === 'age') {
        const birthDate = new Date(item[field.source]);
        const today = new Date('2025-07-01');
        const age = today.getFullYear() - birthDate.getFullYear();
        transformedItem[field.name] = age;
      }
    }

    return transformedItem;
  });
}

const transformedData = transformData(csvData, expectedFormat);

// Output the transformed data as a JSON array
console.log(JSON.stringify(transformedData, null, 2));