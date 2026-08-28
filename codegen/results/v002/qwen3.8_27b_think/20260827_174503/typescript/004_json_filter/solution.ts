import * as fs from "fs";
import * as path from "path";

interface Record {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const dataPath = path.resolve("input/data.json");
const raw: string = fs.readFileSync(dataPath, "utf-8");
const records: Record[] = JSON.parse(raw);

const filtered = records
  .filter((r) => r.active === true && r.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

process.stdout.write(JSON.stringify(filtered) + "\n");