import { readFileSync } from "fs";

const csv = readFileSync("input/data.csv", "utf-8");
const lines = csv.trim().split("\n");

const header = lines[0].split(",");
const rows: Array<{ Name: string; Age: number; Email: string; Score: number }> = [];

for (let i = 1; i < lines.length; i++) {
  const line = lines[i].trim();
  if (!line) continue;
  const fields = line.split(",");
  rows.push({
    Name: fields[0].trim(),
    Age: parseInt(fields[1].trim(), 10),
    Email: fields[2].trim(),
    Score: parseFloat(fields[3].trim()),
  });
}

process.stdout.write(JSON.stringify(rows, null, 2));