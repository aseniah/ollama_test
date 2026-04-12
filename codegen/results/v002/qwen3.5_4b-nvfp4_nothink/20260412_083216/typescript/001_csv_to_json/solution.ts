import * as fs from "fs";
import { readFileSync } from "fs";

const filePath = process.argv[2] || "input/data.csv";

// Read the file contents
const content = fs.readFileSync(filePath, "utf-8");
const lines = content.split("\n").filter((line) => line.trim() !== "");

if (lines.length === 0) {
  console.log("[]");
  process.exit(0);
}

// Parse header to find column indices
const header = lines[0].split(",");
const indexByName = header.indexOf("Name");
const indexByAge = header.indexOf("Age");
const indexByEmail = header.indexOf("Email");
const indexByScore = header.indexOf("Score");

if (
  indexByName === -1 ||
  indexByAge === -1 ||
  indexByEmail === -1 ||
  indexByScore === -1
) {
  console.log([]);
  process.exit(0);
}

const result: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

// Parse data rows starting from the second line
for (let i = 1; i < lines.length; i++) {
  const fields = lines[i].split(",");
  
  // Ensure enough fields exist before parsing, though basic assumptions are made as per prompt
  if (fields[indexByAge] === undefined || fields[indexByEmail] === undefined) {
    continue;
  }

  result.push({
    Name: fields[indexByName],
    Age: parseInt(fields[indexByAge], 10), // Assuming integer format for age (or parseFloat)
    Email: fields[indexByEmail],
    Score: parseFloat(fields[indexByScore]),
  });
}

// Output the JSON array as required
console.log(JSON.stringify(result));