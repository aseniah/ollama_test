const n = parseInt(process.argv[2]);

if (n === undefined || n === null || isNaN(n)) {
  // No valid argument provided, print nothing
} else if (n >= 1) {
  let a = 1;
  let b = 1;
  
  // Print first Fibonacci numbers as long as they are within the limit
  while (a <= n && b <= n) {
    console.log(a);
    console.log(b);
    let c = a + b;
    a = b;
    b = c;
  }
} else {
  // If n < 1, print nothing (do not output anything)
}