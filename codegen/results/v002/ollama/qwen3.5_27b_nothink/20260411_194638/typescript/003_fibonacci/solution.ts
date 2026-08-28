/**
 * Prints all Fibonacci numbers up to and including the largest Fibonacci number
 * that does not exceed N.
 * 
 * Usage: tsx fibonacci.ts <N>
 */

function printFibonacciUpToN(limit: number): void {
  if (limit < 1) {
    return;
  }

  let a = 1;
  let b = 1;

  // Print the first number
  console.log(a);

  // If limit is 1, the next number (1) is still <= 1, but we handle the loop carefully.
  // The sequence is 1, 1, 2, 3, 5...
  // We print 'a', then calculate next, print if <= limit, update.
  
  // Since the sequence starts 1, 1, we need to handle the second '1' specifically 
  // or structure the loop to handle the pair (a, b) correctly.
  
  // Revised logic for strict sequence: 1, 1, 2, 3, 5
  // We start by printing the first 1.
  // Then we generate the next one.
  
  // Let's restart the loop structure to be cleaner:
  let current = 1;
  let next = 1;

  // Print the first term
  if (current <= limit) {
    console.log(current);
  }

  // Generate subsequent terms
  // We use a while loop to generate the next number.
  // If we stop here, we haven't printed the second '1' yet if N >= 1.
  
  // Actually, let's iterate:
  // F(1) = 1
  // F(2) = 1
  // F(3) = 2
  
  let f1 = 1;
  let f2 = 1;
  
  if (f1 <= limit) console.log(f1);
  if (f2 <= limit) console.log(f2);
  
  while (true) {
    const fNext = f1 + f2;
    if (fNext > limit) break;
    console.log(fNext);
    f1 = f2;
    f2 = fNext;
  }
}

const main = () => {
  // Get the argument from process.argv
  // process.argv[0] is tsx/node
  // process.argv[1] is the script name
  // process.argv[2] is the first argument
  const input = process.argv[2];

  if (input === undefined) {
    // No argument provided, do nothing or handle error?
    // The prompt implies we accept a single integer argument.
    // If missing, we do nothing.
    return;
  }

  const n = parseInt(input, 10);

  if (isNaN(n)) {
    // If the input is not a valid integer, do nothing.
    return;
  }

  printFibonacciUpToN(n);
};

main();