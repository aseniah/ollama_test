function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTest(description: string, price: number, quantity: number, expected: number): void {
    const got = calculateDiscount(price, quantity);
    if (Math.abs(got - expected) < 1e-9) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${got})`);
    }
}

// Test cases including boundary conditions
runTest("Quantity 9 (just below first threshold)", 10, 9, 10 * 9);
runTest("Quantity 10 (at first threshold, 10% off)", 10, 10, 10 * 10 * 0.9);
runTest("Quantity 49 (just below second threshold)", 10, 49, 10 * 49 * 0.9);
runTest("Quantity 50 (at second threshold, 25% off)", 10, 50, 10 * 50 * 0.75);