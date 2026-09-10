import { readFileSync } from "fs";
import { resolve } from "path";

const dataPath = resolve("input/data.json");
const raw = readFileSync(dataPath, "utf-8");
const data = JSON.parse(raw) as Array<{
  name: string;
  age: number;
  active: boolean;
  score: number;
}>;

const filtered = data
  .filter((record) => record.active === true && record.age >= 30)
  .sort((a, b) => a.name.localeCompare(b.name));

process.stdout.write(JSON.stringify(filtered, null, 2) + "\n");