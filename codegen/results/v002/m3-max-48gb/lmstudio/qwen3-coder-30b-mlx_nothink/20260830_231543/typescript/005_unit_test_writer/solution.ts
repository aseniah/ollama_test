function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTest(description: string, price: number, quantity: number, expected: number): void {
    const result = calculateDiscount(price, quantity);
    const tolerance = 0.0001; // For floating point comparison
    
    if (Math.abs(result - expected) < tolerance) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${result})`);
    }
}

// Test cases
runTest("quantity 9 (no discount)", 10, 9, 90);
runTest("quantity 10 (10% discount)", 10, 10, 90);
runTest("quantity 49 (10% discount)", 10, 49, 441);
runTest("quantity 50 (25% discount)", 10, 50, 375);
runTest("quantity 5 (no discount)", 5, 5, 25);
runTest("quantity 15 (10% discount)", 20, 15, 270);
runTest("quantity 100 (25% discount)", 8, 100, 600);