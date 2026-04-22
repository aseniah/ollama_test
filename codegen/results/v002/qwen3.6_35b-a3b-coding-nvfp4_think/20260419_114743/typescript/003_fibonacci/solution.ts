const n = parseInt(process.argv[2], 10);

if (isNaN(n) || n < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

while (a <= n) {
  console.log(a);
  [a, b] = [b, a + b];
}