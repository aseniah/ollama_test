process.argv[2]
  ? parseInt(process.argv[2])
  : undefined
);

if (n === undefined || isNaN(n)) {
  process.exit(1);
}

if (n < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

if (n >= 1) {
  console.log(a);
}

while (true) {
  let next = a + b;
  if (next > n) {
    break;
  }
  console.log(next);
  a = b;
  b = next;
}

process.exit(0);