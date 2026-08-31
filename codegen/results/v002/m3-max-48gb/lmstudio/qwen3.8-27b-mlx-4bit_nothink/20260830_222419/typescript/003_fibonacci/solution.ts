function main(): void {
  const n = parseInt(process.argv[2] ?? "", 10);

  if (isNaN(n) || n < 1) {
    return;
  }

  let a = 1;
  let b = 1;

  // Print first 1
  console.log(a);

  // If N is exactly 1, we should only print one 1? 
  // The sequence is 1, 1, 2, 3, 5...
  // For N=1, we print 1. Should we print the second 1? 
  // "up to and including the largest Fibonacci number that does not exceed N"
  // Fibonacci numbers: 1, 1, 2, 3, 5, 8, ...
  // If N=1, the largest fib <= 1 is 1. 
  // The sequence starts 1, 1, 2...
  // Usually for such problems, if N=1, we print 1. 
  // If N=2, we print 1, 1, 2.
  // If N=3, we print 1, 1, 2, 3.
  
  // So we need to print all Fibonacci numbers F_k where F_k <= N.
  // F_1 = 1, F_2 = 1, F_3 = 2, F_4 = 3, ...
  
  // We already printed F_1.
  if (n >= 1) {
    console.log(b); // Print F_2 which is 1
  }

  a = 1;
  b = 1;
  
  // Generate subsequent Fibonacci numbers
  while (true) {
    const c = a + b;
    if (c > n) {
      break;
    }
    console.log(c);
    a = b;
    b = c;
  }
}

main();