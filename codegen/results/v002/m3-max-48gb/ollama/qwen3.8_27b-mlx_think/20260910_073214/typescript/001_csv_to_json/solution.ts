import * as fs from "fs";

const csv = fs.readFileSync("input/data.csv", "utf-8");
const lines = csv.trim().split("\n");

// Skip header
const rows = lines.slice(1);

const result = rows.map((line) => {
  const parts = line.split(",");
  return {
    Name: parts[0],
    Age: parseInt(parts[1], 10),
    Email: parts[2],
    Score: parseFloat(parts[3]),
  };
});

process.stdout.write(JSON.stringify(result, null, 2) + "\n");