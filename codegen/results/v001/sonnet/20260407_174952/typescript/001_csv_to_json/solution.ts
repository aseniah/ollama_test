import { readFileSync } from "fs";
import { resolve } from "path";

const filePath = resolve("input/data.csv");
const content = readFileSync(filePath, "utf-8");
const lines = content.trim().split("\n");

const headers = lines[0].split(",").map((h) => h.trim());
const nameIdx = headers.indexOf("Name");
const ageIdx = headers.indexOf("Age");
const emailIdx = headers.indexOf("Email");
const scoreIdx = headers.indexOf("Score");

const result = lines.slice(1).map((line) => {
  const cols = line.split(",").map((c) => c.trim());
  return {
    Name: cols[nameIdx],
    Age: parseInt(cols[ageIdx], 10),
    Email: cols[emailIdx],
    Score: parseFloat(cols[scoreIdx]),
  };
});

console.log(JSON.stringify(result));
