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
runTest("quantity = 9 (no discount)", 100, 9, 900);
runTest("quantity = 10 (10% discount)", 100, 10, 900);
runTest("quantity = 49 (10% discount)", 100, 49, 4410);
runTest("quantity = 50 (25% discount)", 100, 50, 3750);

// Additional test cases
runTest("quantity = 1 (no discount)", 100, 1, 100);
runTest("quantity = 100 (25% discount)", 100, 100, 7500);

// Program exit
process.exit(0);