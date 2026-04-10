const arg = process.argv[2];
if (!arg) {
  console.error("Usage: tsx script.ts <N>");
  process.exit(1);
}

const n = parseInt(arg, 10);

if (isNaN(n) || n < 1) {
  return;
}

let a = 1;
let b = 1;

console.log(a);
a = b;
b = a + b;

while (b <= n) {
  console.log(b);
  const next = a + b;
  a = b;
  b = next;
}