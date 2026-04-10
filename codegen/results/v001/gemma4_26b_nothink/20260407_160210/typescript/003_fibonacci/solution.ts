/**
 * This script reads a single integer N from the command line arguments.
 * It prints the Fibonacci sequence starting from 1, 1, 2, ...
 * up to the largest Fibonacci number that does not exceed N.
 * 
 * Usage: tsx script.ts <N>
 */

function main(): void {
  // Retrieve command line arguments
  const args = process.argv.slice(2);

  // Validate that an argument was provided
  if (args.length === 0) {
    return;
  }

  // Parse the argument as an integer
  const n = parseInt(args[0], 10);

  // If N is not a valid number or N < 1, print nothing per requirements
  if (isNaN(n) || n < 1) {
    return;
  }

  // Fibonacci sequence generation logic
  let prev = 1;
  let curr = 1;

  // Edge case: If N is 1, the loop logic needs to handle the first 1.
  // Since the sequence starts 1, 1, 2... 
  // We print the first 1 manually or via loop logic.
  
  // We'll use a simple iterative approach.
  // We start by printing the first '1'.
  console.log(prev);

  // If the second '1' is also <= N, we print it.
  if (curr <= n) {
    // Check if we are looking at the first 1 or a subsequent number.
    // Since the sequence starts 1, 1... we check the second term.
    // However, to strictly follow the 1, 1, 2... pattern:
    
    // Let's refine the loop:
    // We track the sequence. The first two terms are 1 and 1.
    
    // Start with the second 1.
    let a = 1; // term i-2
    let b = 1; // term i-1
    
    // We already printed the first '1'.
    // Now handle the second '1'.
    if (b <= n) {
        // We need to be careful not to double print if N=1 and we logic incorrectly.
        // But the requirement says "Print all... up to... largest... not exceeding N".
        // For N=1, the sequence is 1, 1.
        
        // If the first 1 was printed, let's check if we can print a second 1.
        // Actually, let's use a cleaner loop.
    }
  }
}

/**
 * Revised implementation for clarity and correctness.
 */
function runFibonacci(n: number): void {
  if (n < 1) return;

  let a = 1;
  let b = 1;

  // Handle the first term
  console.log(a);

  // If N is at least 1, the second term is also 1.
  // We check if the next term (1) is <= N.
  if (n >= 1) {
    // To avoid infinite loop or logical errors, 
    // we handle the first two 1s explicitly or via a controlled loop.
    // Sequence: 1, 1, 2, 3, 5...
    
    // We already printed the first 1.
    // Let's check if we can print the second 1.
    // Since N >= 1, the second 1 is always valid.
    // But wait, if N=1, the largest Fib <= 1 is 1. 
    // Does "all Fibonacci numbers up to..." include both 1s?
    // Standard Fibonacci sequence definition: F1=1, F2=1.
    // The prompt says: "Start the sequence: 1, 1, 2, 3, 5, ..."
    // This implies both 1s should be printed if N >= 1.
    
    console.log(b);
    
    // Now calculate subsequent numbers: 2, 3, 5...
    let next = a + b;
    while (next <= n) {
      console.log(next);
      a = b;
      b = next;
      next = a + b;
    }
  }
}

// Re-implementing main to be robust
function execute(): void {
  const args = process.argv.slice(2);
  if (args.length === 0) return;

  const n = parseInt(args[0], 10);
  if (isNaN(n) || n < 1) return;

  // The Fibonacci sequence: F(1)=1, F(2)=1, F(3)=2, F(4)=3...
  // We print F(i) such that F(i) <= N.

  let sequence: number[] = [];
  let f1 = 1;
  let f2 = 1;

  if (f1 <= n) sequence.push(f1);
  if (f2 <= n) sequence.push(f2);

  let next = f1 + f2;
  while (next <= n) {
    sequence.push(next);
    f1 = f2;
    f2 = next;
    next = f1 + f2;
  }

  // Special case: If N=1, the loop above handles it, but we need to ensure 
  // we don't print two 1s if the logic is flawed. 
  // If N=1, sequence is [1, 1].
  // If N=2, sequence is [1, 1, 2].
  
  // The only edge case is if N=1, f1 and f2 are both 1.
  // The logic above correctly produces [1, 1] for N=1.
  
  for (const num of sequence) {
    console.log(num);
  }
}

execute();