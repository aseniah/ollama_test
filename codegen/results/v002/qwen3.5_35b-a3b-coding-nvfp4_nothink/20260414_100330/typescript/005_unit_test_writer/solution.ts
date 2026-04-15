// Function to test
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Helper to run a test case
function runTest(
    description: string,
    price: number,
    quantity: number,
    expected: number
): void {
    const actual = calculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison
    const epsilon = 0.0001;
    const diff = Math.abs(actual - expected);
    
    if (diff <= epsilon) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
    }
}

// Main test runner
function main(): void {
    // Test case 1: quantity < 10 (no discount)
    // Boundary: quantity = 9
    runTest(
        "No discount for quantity 9",
        100, // price
        9,   // quantity
        900  // 100 * 9
    );

    // Test case 2: quantity >= 10 and < 50 (10% discount)
    // Boundary: quantity = 10
    runTest(
        "10% discount for quantity 10",
        100,
        10,
        900  // 100 * 10 * 0.9
    );

    // Test case 3: quantity >= 10 and < 50 (10% discount)
    // Boundary: quantity = 49
    runTest(
        "10% discount for quantity 49",
        100,
        49,
        4410 // 100 * 49 * 0.9
    );

    // Test case 4: quantity >= 50 (25% discount)
    // Boundary: quantity = 50
    runTest(
        "25% discount for quantity 50",
        100,
        50,
        3750 // 100 * 50 * 0.75
    );
}

main();