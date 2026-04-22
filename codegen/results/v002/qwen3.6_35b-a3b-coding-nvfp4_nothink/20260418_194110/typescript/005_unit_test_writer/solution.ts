function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTest(price: number, quantity: number, expected: number, description: string): void {
    const result = calculateDiscount(price, quantity);
    const tolerance = 0.001;
    const pass = Math.abs(result - expected) < tolerance;
    if (pass) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${result})`);
    }
}

// Test cases covering boundary conditions and various scenarios
runTest(10, 9, 90, "Quantity 9 (below 10, no discount)");
runTest(10, 10, 90, "Quantity 10 (first discount tier, 10% off)");
runTest(10, 49, 441, "Quantity 49 (second discount tier, 10% off)");
runTest(10, 50, 375, "Quantity 50 (third discount tier, 25% off)");
runTest(10, 100, 750, "Quantity 100 (above 50, 25% off)");
runTest(5, 0, 0, "Quantity 0 (edge case, zero quantity)");