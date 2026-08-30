function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTest(description: string, price: number, quantity: number, expected: number): void {
    const actual = calculateDiscount(price, quantity);
    // Using a small epsilon for floating point comparison
    const epsilon = 0.0001;
    const passed = Math.abs(actual - expected) < epsilon;

    if (passed) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
    }
}

// Test cases
runTest("Quantity 9 (below 10 threshold)", 100, 9, 900);
runTest("Quantity 10 (exactly at 10 threshold)", 100, 10, 900);
runTest("Quantity 49 (below 50 threshold)", 100, 49, 4410);
runTest("Quantity 50 (exactly at 50 threshold)", 100, 50, 3750);

// Additional boundary check for quantity < 1
runTest("Quantity 0 (edge case)", 100, 0, 0);