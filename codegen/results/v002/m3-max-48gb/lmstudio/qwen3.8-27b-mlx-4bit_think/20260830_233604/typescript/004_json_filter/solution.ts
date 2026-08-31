import * as fs from "fs";
import * as path from "path";

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const filePath = path.join("input", "data.json");
const data: Record[] = JSON.parse(fs.readFileSync(filePath, "utf-8"));

const result = data
  .filter((r) => r.active === true && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result, null, 2));