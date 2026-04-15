function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTest(
    description: string,
    price: number,
    quantity: number,
    expected: number
): void {
    const actual = calculateDiscount(price, quantity);
    const tolerance = 0.0001;
    const passed = Math.abs(actual - expected) < tolerance;

    if (passed) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
    }
}

// Test cases covering boundary conditions
console.log("Running calculateDiscount tests...\n");

// Test case 1: quantity = 9 (should use no discount)
runTest(
    "quantity 9 (no discount)",
    100,
    9,
    100 * 9
);

// Test case 2: quantity = 10 (should use 10% discount)
runTest(
    "quantity 10 (10% discount)",
    100,
    10,
    100 * 10 * 0.9
);

// Test case 3: quantity = 49 (should use 10% discount)
runTest(
    "quantity 49 (10% discount)",
    100,
    49,
    100 * 49 * 0.9
);

// Test case 4: quantity = 50 (should use 25% discount)
runTest(
    "quantity 50 (25% discount)",
    100,
    50,
    100 * 50 * 0.75
);

// Additional test case: quantity = 0 (edge case)
runTest(
    "quantity 0 (zero quantity)",
    100,
    0,
    100 * 0
);

// Additional test case: quantity = 100 (well into bulk discount)
runTest(
    "quantity 100 (25% discount)",
    100,
    100,
    100 * 100 * 0.75
);

console.log("\nTests completed.");