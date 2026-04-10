// Define the calculateDiscount function
function calculateDiscount(price: number, quantity: number): number {
  if (quantity < 10) return price * quantity;
  else if (quantity < 50) return price * quantity * 0.9;
  else return price * quantity * 0.75;
}

// Test cases including boundary conditions
const testCases = [
  {
    description: "Test with quantity 9 (below 10)",
    inputs: { price: 100, quantity: 9 },
    expectedOutput: 900,
  },
  {
    description: "Test with quantity 10 (exactly 10)",
    inputs: { price: 100, quantity: 10 },
    expectedOutput: 1000,
  },
  {
    description: "Test with quantity 49 (just below 50)",
    inputs: { price: 100, quantity: 49 },
    expectedOutput: 4900,
  },
  {
    description: "Test with quantity 50 (exactly 50)",
    inputs: { price: 100, quantity: 50 },
    expectedOutput: 5000,
  },
];

// Function to run tests and print results
function runTests(testCases: Array<{ description: string; inputs: { price: number; quantity: number }; expectedOutput: number }>): void {
  for (const { description, inputs, expectedOutput } of testCases) {
    const result = calculateDiscount(inputs.price, inputs.quantity);
    if (result === expectedOutput) {
      console.log(`PASS: ${description}`);
    } else {
      console.log(`FAIL: ${description} (expected: ${expectedOutput}, got: ${result})`);
    }
  }
}

// Run the tests
runTests(testCases);