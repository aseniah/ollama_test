function solve() {
  const arg = process.argv[2];

  // If no argument is provided, exit without printing anything
  if (!arg) {
    return;
  }

  const N = parseInt(arg, 10);

  // If N < 1 (or if it's not a valid number), print nothing and exit.
  if (isNaN(N) || N < 1) {
    return;
  }

  let prev = 0;
  let curr = 1;

  // Generate and print Fibonacci numbers while they do not exceed N
  while (curr <= N) {
    console.log(curr);
    
    // Calculate the next number in the sequence
    const next = prev + curr;
    prev = curr;
    curr = next;
  }
}

solve();