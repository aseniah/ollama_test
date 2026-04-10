import fs from "fs";
import path from "path";

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const data: Record[] = JSON.parse(fs.readFileSync(path.join("input", "data.json"), "utf-8"));

const result = data
  .filter((r) => r.active && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

console.log(JSON.stringify(result));
