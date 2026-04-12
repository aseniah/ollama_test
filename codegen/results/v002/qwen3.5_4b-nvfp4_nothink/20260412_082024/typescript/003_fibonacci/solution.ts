// Read command-line argument
if (process.argv.length !== 3) {
  console.error('Usage: node script.js <N>');
  process.exit(1);
}

const n = parseInt(process.argv[2], 10);

// If N is less than 1, print nothing and exit
if (n < 1) {
  process.exit(0);
}

let a = 1; // First Fibonacci number (F(1))
let b = 1; // Second Fibonacci number (F(2))

// Print the first two numbers if they are within range
console.log(a);
if (a <= n) {
  console.log(b);
}

// Generate subsequent Fibonacci numbers starting from the third one (2, 3, 5...)
while (true) {
  const next = a + b; // F(3), F(4), etc.
  
  // If next is within range, print it and update for next iteration
  if (next <= n) {
    console.log(next);
    // Shift a and b forward: a becomes the current b, b becomes the new 'next'
    a = b;
    b = next;
  } else {
    // The next Fibonacci number exceeds N, so we are done.
    // We do not print 'next', as it should be strictly <= N.
    break;
  }
}