const args = process.argv.slice(2);

if (args.length === 0) {
  process.exit(0);
}

const n = parseInt(args[0], 10);

if (isNaN(n) || n < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

while (a <= n) {
  console.log(a);
  [a, b] = [b, a + b];
}