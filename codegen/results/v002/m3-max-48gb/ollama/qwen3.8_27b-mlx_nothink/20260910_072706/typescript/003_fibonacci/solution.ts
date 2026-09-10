const N = parseInt(process.argv[2], 10);

if (isNaN(N) || N < 1) {
  // Print nothing
  process.exit(0);
}

// Start with 1, 1
let a = 1;
let b = 1;

// Print the first 1
console.log(a);

if (N >= 1) {
  // Print the second 1 only if it's <= N
  // But we need to be careful: if N is exactly 1, we print just one "1"
  // Actually, the sequence starts 1, 1, 2, 3, 5...
  // If N < 1, print nothing. If N >= 1, we need to print Fibonacci numbers up to N.
  
  // Let me reconsider: the sequence is 1, 1, 2, 3, 5, ...
  // We print all that don't exceed N.
  
  // If N >= 1, print first 1
  // If N >= 1, we should print the second 1 as well (since 1 <= N when N >= 1)
  
  // Wait, if N = 1, we print "1" once or twice?
  // The sequence is 1, 1, 2, 3, 5. Up to and including the largest Fibonacci number that doesn't exceed N.
  // If N = 1, the largest Fibonacci number that doesn't exceed 1 is 1.
  // But the sequence has two 1s. Do we print both?
  
  // The problem says "Start the sequence: 1, 1, 2, 3, 5, ..."
  // So the sequence includes two 1s at the start. If N >= 1, both 1s are <= N, so we print both.
  
  // But actually, let me think again. If N = 1, we print 1, 1.
  // If N = 0, we print nothing.
  
  console.log(b);
  
  let c = a + b;
  while (c <= N) {
    console.log(c);
    a = b;
    b = c;
    c = a + b;
  }
}