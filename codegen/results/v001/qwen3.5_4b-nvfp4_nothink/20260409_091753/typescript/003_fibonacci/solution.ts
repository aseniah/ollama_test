import { writeFileSync } from 'fs';

// Read N from command line arguments
const args = process.argv.slice(1);
if (args.length !== 1) {
  console.error('Usage: node script.js <number>');
  process.exit(1);
}

const nStr = args[0];
if (!/^\d+$/.test(nStr)) {
  console.error('Error: Argument must be a valid integer.');
  process.exit(1);
}

const n = parseInt(nStr, 10);

if (n < 1) {
  // Nothing to print
  writeFileSync(process.argv[1], '');
  process.exit(0);
}

let a = 1;
let b = 1;
const sequence: number[] = [a];

function printFibonacci(n: number): void {
  let i = 0;
  
  while (true) {
    // Current term is 'i' (sequence index starting at 2 for value 1, then 2, etc.)
    if (i === 0) {
      console.log(a);
      continue;
    }
    
    if (i === 1) {
      console.log(b);
      // Check next term
      const c = a + b;
      if (c <= n) {
        a = b;
        b = c;
        i++;
        sequence.push(c);
      } else {
        break;
      }
    } else {
      console.log(b);
      // Calculate next term
      const next = a + b;
      
      if (next <= n) {
        a = b;
        b = next;
        i++;
        sequence.push(next);
      } else {
        break;
      }
    }
  }
}

// Alternative simpler implementation logic to avoid potential loop confusion:
a = 1;
b = 1;
console.log(a); // Print first '1'
sequence.push(a);

while (true) {
  const nextVal = a + b;
  
  if (nextVal <= n) {
    console.log(nextVal);
    sequence.push(nextVal);
    
    // Update values for next iteration: new sequence index is now 'nextVal' in original logic
    a = b;
    b = nextVal;
    
    // If next step's value exceeds N, we are done
    if (a + b > n) {
      break;
    }
  } else {
    break;
  }
}

writeFileSync(process.argv[1], '');
process.exit(0);