import { argv } from 'node:process';
import { writeFileSync, mkdtempSync, unlinkSync } from 'node:fs';
import { rm } from 'node:fs/promises';

function main() {
  const args = argv.slice(2);
  if (args.length === 0) {
    console.error('Usage: tsx fibonacci.ts <N>');
    process.exit(1);
  }

  const n = parseInt(args[0], 10);

  if (n < 1) {
    return;
  }

  const fibs: number[] = [];
  
  let a = 1;
  let b = 1; // F(1) and F(2) both 1
  
  while (true) {
    if (a <= n) {
      fibs.push(a);
    } else {
      break;
    }

    const nextFib = a + b;
    
    if (b <= n) {
      // We need to track the current 'b' value too because it might be <= n
      // Actually, let's rethink the loop logic.
      // F sequence: 1, 1, 2, 3, 5...
      // Current 'a' is the latest added number. 'b' is the next candidate.
    }

    a = b;
    b = nextFib;
  }

  // Re-evaluate the loop logic to ensure correctness for N < 1 (handled) and sequence generation
  fibs.length = 0;
  
  let prev = 0;
  let curr = 1;
  
  while (true) {
    if (curr > n) break;
    
    fibs.push(curr);
    const nextVal = prev + curr;
    prev = curr;
    curr = nextVal;
  }

  for (const num of fibs) {
    console.log(num.toString());
  }
}

main();