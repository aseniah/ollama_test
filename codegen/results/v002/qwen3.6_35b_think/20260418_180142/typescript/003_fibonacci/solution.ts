const input = process.argv[2];
const n: number = parseInt(input || '', 10);

if (isNaN(n) || n < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

while (a <= n) {
  console.log(a);
  const next = a + b;
  a = b;
  b = next;
}