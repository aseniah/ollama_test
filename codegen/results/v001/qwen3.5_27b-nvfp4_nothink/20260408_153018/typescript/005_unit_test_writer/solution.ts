// Copy of the function to be tested
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner logic
function runTest(
  description: string, 
  price: number, 
  quantity: number, 
  expected: number
): void {
  const actual = calculateDiscount(price, quantity);
  
  // Use a small epsilon for floating point comparison safety, though integers usually work fine here
  const epsilon = 0.0001;
  const passed = Math.abs(actual - expected) < epsilon;

  if (passed) {
    console.log(`PASS: ${description}`);
  } else {
    console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
  }
}

// Main execution block
function main(): void {
  // Test Case 1: Boundary condition (just below 10, no discount)
  // Price: 10, Qty: 9 -> Expected: 10 * 9 = 90
  runTest("Boundary: Quantity 9 (no discount)", 10, 9, 90);

  // Test Case 2: Boundary condition (at 10, 10% discount)
  // Price: 10, Qty: 10 -> Expected: 10 * 10 * 0.9 = 90
  runTest("Boundary: Quantity 10 (10% discount)", 10, 10, 90);

  // Test Case 3: Boundary condition (just below 50, 10% discount)
  // Price: 10, Qty: 49 -> Expected: 10 * 49 * 0.9 = 441
  runTest("Boundary: Quantity 49 (10% discount)", 10, 49, 441);

  // Test Case 4: Boundary condition (at 50, 25% discount)
  // Price: 10, Qty: 50 -> Expected: 10 * 50 * 0.75 = 375
  runTest("Boundary: Quantity 50 (25% discount)", 10, 50, 375);

  // Additional Test Case: Large quantity to verify logic flow
  // Price: 100, Qty: 100 -> Expected: 100 * 100 * 0.75 = 7500
  runTest("Large Quantity: Quantity 100 (25% discount)", 100, 100, 7500);
}

main();