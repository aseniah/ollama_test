const arg = process.argv[2];

if (arg === undefined) {
  process.exit(0);
}

const N = BigInt(arg);

if (N < 1n) {
  process.exit(0);
}

let a: bigint = 1n;
let b: bigint = 1n;

// Print first 1
console.log(a.toString());

// Print second 1 (since N >= 1, both 1s are valid)
console.log(b.toString());

let next: bigint = a + b;
while (next <= N) {
  console.log(next.toString());
  a = b;
  b = next;
  next = a + b;
}