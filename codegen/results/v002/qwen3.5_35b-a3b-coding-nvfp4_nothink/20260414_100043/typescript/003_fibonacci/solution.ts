const args = process.argv.slice(2);
if (args.length < 1) {
  process.exit(0);
}

const n = parseInt(args[0], 10);
if (isNaN(n) || n < 1) {
  process.exit(0);
}

let a = 1;
let b = 1;

// Print the first 1 if it's <= N
if (a <= n) {
  console.log(a);
}

// Print the second 1 if it's <= N
if (b <= n && b !== a) {
  // This case only happens if we started with different values, but here a=b=1.
  // Since we already printed a (1), we need to handle the second 1 explicitly.
  // However, the sequence is 1, 1, 2...
  // If N >= 1, we print 1. Then we need to print the next 1 if N >= 1.
  // Let's restructure the loop to handle the sequence 1, 1, 2, 3... cleanly.
}

// Reset and use a cleaner approach for 1, 1, 2, 3...
let prev = 1;
let curr = 1;
let first = true;

while (prev <= n) {
  if (first) {
    console.log(prev);
    first = false;
  } else {
    if (prev <= n) {
      console.log(prev);
    }
  }
  
  const next = prev + curr;
  prev = curr;
  curr = next;
}

// The above logic is slightly flawed for the second 1. Let's rewrite completely.