import { readFileSync } from "fs";

const csv = readFileSync("input/data.csv", "utf-8");
const lines = csv.trim().split("\n");
const headers = lines[0].split(",").map((h) => h.trim());

const result = lines.slice(1).map((line) => {
  const values = line.split(",").map((v) => v.trim());
  const row: Record<string, string> = {};
  headers.forEach((h, i) => (row[h] = values[i]));

  return {
    Name: row["Name"],
    Age: parseInt(row["Age"], 10),
    Email: row["Email"],
    Score: parseFloat(row["Score"]),
  };
});

console.log(JSON.stringify(result));
