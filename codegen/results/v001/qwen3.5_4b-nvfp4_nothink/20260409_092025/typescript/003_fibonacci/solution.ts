import { stdout } from 'node:process';

const args = process.argv.slice(2);
if (args.length !== 1) {
  console.error('Usage: tsx script.ts N');
  process.exit(1);
}

const nStr = args[0];
if (!/^\-?\d+$/.test(nStr)) {
  console.error('Argument must be an integer.');
  process.exit(1);
}

const n = parseInt(nStr, 10);

if (n < 1) {
  return;
}

let a = 1;
let b = 1;

// Print the first term explicitly if N >= 1
stdout.write(a + '\n');

if (n >= 2) {
  // The sequence starts 1, 1, 2, 3, 5...
  // We already printed the first '1'. Now we need to generate subsequent terms.
  // b is the second '1'.
  
  // Loop while the current term (b + a) does not exceed n
  while (true) {
    const next = a + b;
    
    if (next > n) {
      break;
    }
    
    stdout.write(next + '\n');
    
    a = b;
    b = next;
  }
}