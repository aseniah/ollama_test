import * as fs from 'fs';
import * as path from 'path';

// Define the structure of the data
interface User {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

// Read the input file
const filePath = path.join(__dirname, 'input', 'data.json');
const rawData = fs.readFileSync(filePath, 'utf-8');

// Parse the JSON data
let users: User[] = [];
try {
  users = JSON.parse(rawData);
} catch (error) {
  console.error('Error parsing JSON:', error);
  process.exit(1);
}

// Filter and sort the users
const filteredUsers = users.filter(user => user.active && user.age >= 30).sort((a, b) => a.name.localeCompare(b.name));

// Output the result as JSON
console.log(JSON.stringify(filteredUsers, null, 2));