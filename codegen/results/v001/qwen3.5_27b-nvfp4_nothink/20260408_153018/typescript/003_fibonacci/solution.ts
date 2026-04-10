// Read the second argument from process.argv (index 2) as the command line input
const input = process.argv[2];

if (!input) {
  // No argument provided, exit silently or with error as needed
  process.exit(0);
}

const N = parseInt(input, 10);

// If N is not a number or is less than 1, print nothing and exit
if (isNaN(N) || N < 1) {
  process.exit(0);
}

// Start the sequence: 1, 1, ...
let a = 1;
let b = 1;

// Function to print current number
const printNum = (val: number) => console.log(val);

// Handle the first two numbers (1, 1) if they are <= N
if (a <= N) {
  printNum(a);
}
if (b <= N) {
  printNum(b);
}

// Generate subsequent Fibonacci numbers
let next;
while (true) {
  next = a + b;
  
  if (next > N) {
    break;
  }

  printNum(next);

  // Shift the window: a becomes b, b becomes next
  a = b;
  b = next;
}