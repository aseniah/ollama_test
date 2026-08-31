const n = parseInt(process.argv[2]);

if (isNaN(n) || n < 1) {
  process.exit(0);
}

let a = 1, b = 1;
console.log(a);
if (b <= n) console.log(b);
let next = a + b;
while (next <= n) {
  console.log(next);
  a = b;
  b = next;
  next = a + b;
}