import fs from "fs";
import path from "path";

const AS_OF = new Date("2025-07-01");

function parseDate(mmddyyyy: string): Date {
  const [m, d, y] = mmddyyyy.split("/").map(Number);
  return new Date(y, m - 1, d);
}

function formatDate(d: Date): string {
  const y = d.getFullYear();
  const m = String(d.getMonth() + 1).padStart(2, "0");
  const day = String(d.getDate()).padStart(2, "0");
  return `${y}-${m}-${day}`;
}

function calcAge(birthday: Date, asOf: Date): number {
  let age = asOf.getFullYear() - birthday.getFullYear();
  const hadBirthday =
    asOf.getMonth() > birthday.getMonth() ||
    (asOf.getMonth() === birthday.getMonth() && asOf.getDate() >= birthday.getDate());
  if (!hadBirthday) age--;
  return age;
}

function splitName(full: string): { FirstName: string; LastName: string } {
  const parts = full.trim().split(/\s+/);
  return { FirstName: parts[0], LastName: parts[parts.length - 1] };
}

const csvPath = path.join(path.dirname(process.argv[1]), "input", "input.csv");
const raw = fs.readFileSync(csvPath, "utf8").trim();
const lines = raw.split("\n");
const headers = lines[0].split(",");

const result = lines.slice(1).map((line) => {
  const cols = line.split(",");
  const row: Record<string, string> = {};
  headers.forEach((h, i) => (row[h.trim()] = (cols[i] ?? "").trim()));

  const { FirstName, LastName } = splitName(row.Name);
  const birthday = parseDate(row.Birthday);
  const died = row.Died && row.Died !== "null" ? parseDate(row.Died) : null;
  const ageRef = died ?? AS_OF;
  const age = calcAge(birthday, ageRef);

  const relativeFields = ["Father", "Mother", "Brother", "Sister"] as const;
  const relatives = relativeFields
    .filter((f) => row[f] && row[f] !== "null")
    .map((f) => {
      const { FirstName: fn, LastName: ln } = splitName(row[f]);
      return { FirstName: fn, LastName: ln, Relationship: f };
    });

  return {
    FirstName,
    LastName,
    Birthday: formatDate(birthday),
    Age: age,
    Relatives: relatives,
  };
});

console.log(JSON.stringify(result, null, 2));
