const n = parseInt(process.argv[2], 10);

if (n < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

console.log(a);

while (b <= n) {
  console.log(b);
  [a, b] = [b, a + b];
}
