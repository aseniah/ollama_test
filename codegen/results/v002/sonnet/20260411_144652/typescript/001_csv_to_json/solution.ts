import fs from "fs";
import path from "path";

const filePath = path.resolve("input/data.csv");
const content = fs.readFileSync(filePath, "utf-8");
const lines = content.trim().split("\n");
const headers = lines[0].split(",");

const result = lines.slice(1).map((line) => {
  const values = line.split(",");
  return {
    Name: values[0].trim(),
    Age: parseInt(values[1].trim(), 10),
    Email: values[2].trim(),
    Score: parseFloat(values[3].trim()),
  };
});

process.stdout.write(JSON.stringify(result));
