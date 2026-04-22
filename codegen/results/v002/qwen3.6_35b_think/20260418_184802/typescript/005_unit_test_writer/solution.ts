/**
 * Calculates a bulk discount based on quantity.
 * - quantity < 10: No discount
 * - 10 <= quantity < 50: 10% discount (multiply by 0.9)
 * - quantity >= 50: 25% discount (multiply by 0.75)
 */
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

/**
 * Helper function to run a single test case and print the result.
 * @param description - A brief description of the test case.
 * @param actual - The value returned by the function being tested.
 * @param expected - The expected value for this test case.
 */
function runTest(description: string, actual: number, expected: number): void {
    // Using a small epsilon for floating point comparison to handle precision issues
    if (Math.abs(actual - expected) < 1e-9) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
    }
}

/**
 * Main execution: runs the test suite.
 */
function main(): void {
    const basePrice = 100;

    // Test 1: Quantity below threshold (quantity < 10)
    // Expected: 100 * 5 = 500
    runTest(
        "Calculate full price for quantity 5",
        calculateDiscount(basePrice, 5),
        basePrice * 5
    );

    // Test 2: Boundary condition just below first discount tier (quantity = 9)
    // Expected: 100 * 9 = 900
    runTest(
        "Calculate full price for quantity 9 (boundary)",
        calculateDiscount(basePrice, 9),
        basePrice * 9
    );

    // Test 3: Boundary condition for first discount tier (quantity = 10)
    // Expected: 100 * 10 * 0.9 = 900
    runTest(
        "Calculate 10% discount for quantity 10 (boundary)",
        calculateDiscount(basePrice, 10),
        basePrice * 10 * 0.9
    );

    // Test 4: Boundary condition within second tier (quantity = 49)
    // Expected: 100 * 49 * 0.9 = 4410
    runTest(
        "Calculate 10% discount for quantity 49 (boundary)",
        calculateDiscount(basePrice, 49),
        basePrice * 49 * 0.9
    );

    // Test 5: Boundary condition for second discount tier (quantity = 50)
    // Expected: 100 * 50 * 0.75 = 3750
    runTest(
        "Calculate 25% discount for quantity 50 (boundary)",
        calculateDiscount(basePrice, 50),
        basePrice * 50 * 0.75
    );

    // Test 6: Quantity above second tier (quantity > 50)
    // Expected: 100 * 100 * 0.75 = 7500
    runTest(
        "Calculate 25% discount for quantity 100",
        calculateDiscount(basePrice, 100),
        basePrice * 100 * 0.75
    );
}

// Execute the main function
main();