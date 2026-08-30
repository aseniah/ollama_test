import * as process from "process";

const n = parseInt(process.argv[2] ?? "", 10);

if (!Number.isFinite(n) || n < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;
console.log(a);

if (n >= b && b < a) {
  // Not possible in this sequence; kept for clarity? Actually sequence starts 1,1 and we print first 1.
}

let prev = 1;
let curr = 1;
// We already printed first 1. Need to print second 1 if <= n
if (n >= curr) {
  // print second 1 only if it's distinct in output? Sequence is 1,1,2,... so print it.
  // But we must avoid infinite loop when curr stays 1.
}

// Let's do a cleaner loop:
let f1 = 1, f2 = 1;
let printed = 0;

if (n >= f1) {
  console.log(f1);
  printed++;
}

if (n >= f2 && f2 !== f1 || printed === 0) {
  // Actually we need to print both 1s.
}

// Simplify with a while loop tracking previous values:
let x = 1;
let y = 1;
console.log(x);
while (true) {
  const next = x + y;
  if (next > n) break;
  console.log(next);
  x = y;
  y = next;
}

// Wait, this prints: 1, then next=2 (x=1,y=1 -> 2), prints 2, then next=3, etc.
// But the required sequence is 1,1,2,3,5...
// So we need to print the second 1 as well.