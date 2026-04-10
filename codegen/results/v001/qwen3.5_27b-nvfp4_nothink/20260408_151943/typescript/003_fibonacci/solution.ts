import { readFileSync } from 'fs';

// If running via tsx with CLI arguments, use process.argv
const n = parseInt(process.argv[2], 10);

if (isNaN(n) || n < 1) {
  // Print nothing
  return;
}

const fibs: number[] = [];
let a = 1;
let b = 1;

while (a <= n) {
  fibs.push(a);
  [a, b] = [b, a + b];
}

fibs.forEach(fib => console.log(fib));