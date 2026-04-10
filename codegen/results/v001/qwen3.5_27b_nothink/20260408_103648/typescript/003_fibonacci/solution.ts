function main(): void {
  // Parse the integer argument from the command line
  const input = process.argv[2];
  
  if (!input) {
    return;
  }

  const N = parseInt(input, 10);

  // If N < 1, print nothing
  if (isNaN(N) || N < 1) {
    return;
  }

  // Start the sequence: 1, 1, 2, 3, 5, ...
  let a = 1;
  let b = 1;

  // Print the first Fibonacci number if it doesn't exceed N
  if (a <= N) {
    console.log(a);
  }

  // If N is 1, the loop below won't print anything else (b=1, a becomes 2)
  // But we need to handle the sequence correctly.
  // The sequence is 1, 1, 2... so we print the first 1.
  // Then we calculate next numbers.
  
  while (b <= N) {
    console.log(b);
    const next = a + b;
    a = b;
    b = next;
  }
}

main();