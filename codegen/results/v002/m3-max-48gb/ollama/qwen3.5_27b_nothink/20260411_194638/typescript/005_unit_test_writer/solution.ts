// Copy of the function to test
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test runner function
function runTest(
  description: string,
  price: number,
  quantity: number,
  expected: number
): void {
  const got = calculateDiscount(price, quantity);
  
  // Use a small epsilon for floating point comparison
  const epsilon = 0.0001;
  const passed = Math.abs(got - expected) < epsilon;

  if (passed) {
    console.log(`PASS: ${description}`);
  } else {
    console.log(`FAIL: ${description} (expected: ${expected}, got: ${got})`);
  }
}

// Main execution
function main(): void {
  // Test Case 1: Boundary condition (quantity < 10)
  // 10 units * $5 = $50 (No discount)
  runTest("Quantity 9 (No Discount)", 10, 9, 90);

  // Test Case 2: Boundary condition (quantity >= 10, < 50)
  // 10 units * $5 = $50 -> 10% off = $45
  runTest("Quantity 10 (10% Discount)", 10, 10, 90);

  // Test Case 3: Boundary condition (quantity < 50)
  // 49 units * $10 = $490 -> 10% off = $441
  runTest("Quantity 49 (10% Discount)", 10, 49, 441);

  // Test Case 4: Boundary condition (quantity >= 50)
  // 50 units * $10 = $500 -> 25% off = $375
  runTest("Quantity 50 (25% Discount)", 10, 50, 375);
}

main();