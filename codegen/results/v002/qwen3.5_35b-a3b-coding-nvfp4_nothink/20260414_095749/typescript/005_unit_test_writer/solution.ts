function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTest(description: string, price: number, quantity: number, expected: number): void {
    const actual = calculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison
    const epsilon = 0.0001;
    if (Math.abs(actual - expected) < epsilon) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
    }
}

// Test cases covering boundary conditions
// Base price of 100 for easy calculation

// Boundary: quantity < 10 (quantity = 9)
runTest(
    "Quantity 9 (no discount)",
    100,
    9,
    100 * 9 // 900
);

// Boundary: quantity = 10 (10% discount)
runTest(
    "Quantity 10 (10% discount)",
    100,
    10,
    100 * 10 * 0.9 // 900
);

// Boundary: quantity < 50 (quantity = 49, 10% discount)
runTest(
    "Quantity 49 (10% discount)",
    100,
    49,
    100 * 49 * 0.9 // 4410
);

// Boundary: quantity >= 50 (25% discount)
runTest(
    "Quantity 50 (25% discount)",
    100,
    50,
    100 * 50 * 0.75 // 3750
);

// Additional test case: quantity > 50
runTest(
    "Quantity 100 (25% discount)",
    100,
    100,
    100 * 100 * 0.75 // 7500
);