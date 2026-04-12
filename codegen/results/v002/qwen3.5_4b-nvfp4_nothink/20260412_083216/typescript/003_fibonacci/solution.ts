import { parseArgs } from "node:util";

const args = parseArgs({
  stopAt: ["--", "-"],
  options: undefined,
  strict: false,
});

const nStr = args.raw[0];
if (!nStr) {
  process.stderr.write(`Usage: ts node script.ts <N>\n`);
  process.exit(1);
}

try {
  const n = parseInt(nStr, 10);
} catch {
  process.stderr.write(`Argument must be an integer: ${nStr}\n`);
  process.exit(1);
}

if (n < 1) {
  process.stdout.write("");
  process.exit(0);
}

let a = 1;
let b = 1;
const seenN = Math.max(n, 1);

// First output: if n >= 1, print 1
process.stdout.write(`${a}\n`);

while (b <= seenN) {
  process.stdout.write(`${b}\n`);
  [a, b] = [b, a + b];
}