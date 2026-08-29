import * as fs from "node:fs";

interface Person {
  name: string;
  age: number;
  active: boolean;
  score: number;
}

const filePath = process.argv[2] ?? "input/data.json";
const data = JSON.parse(fs.readFileSync(filePath, "utf8")) as Person[];

const result = data
  .filter((person) => person.active === true && person.age >= 30)
  .sort((a, b) => {
    if (a.name < b.name) return -1;
    if (a.name > b.name) return 1;
    return 0;
  });

process.stdout.write(JSON.stringify(result));