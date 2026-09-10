import { readFileSync } from "fs";

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const data: Record[] = JSON.parse(readFileSync("input/data.json", "utf-8"));

const result = data
  .filter((r) => r.active === true && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

process.stdout.write(JSON.stringify(result, null, 2) + "\n");