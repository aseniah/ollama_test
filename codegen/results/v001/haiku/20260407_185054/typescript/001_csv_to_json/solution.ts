import fs from "fs";
import path from "path";

const filePath = path.join(process.cwd(), "input/data.csv");
const content = fs.readFileSync(filePath, "utf-8");
const lines = content.trim().split("\n");

const headers = lines[0].split(",").map((h) => h.trim());
const result = lines.slice(1).map((line) => {
  const values = line.split(",").map((v) => v.trim());
  return {
    Name: values[0],
    Age: parseInt(values[1], 10),
    Email: values[2],
    Score: parseFloat(values[3]),
  };
});

console.log(JSON.stringify(result));
