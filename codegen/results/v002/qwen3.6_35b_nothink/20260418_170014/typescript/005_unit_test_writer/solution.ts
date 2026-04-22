function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function testCase(description: string, price: number, quantity: number, expected: number): void {
    const actual = calculateDiscount(price, quantity);
    if (Math.abs(actual - expected) < 1e-9) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
    }
}

// Test case 1: quantity = 9 (below first threshold, no discount)
testCase("Quantity 9 - no discount (price=100)", 100, 9, 100 * 9);

// Test case 2: quantity = 10 (first threshold, 10% discount)
testCase("Quantity 10 - 10% discount (price=100)", 100, 10, 100 * 10 * 0.9);

// Test case 3: quantity = 49 (second threshold, 10% discount)
testCase("Quantity 49 - 10% discount (price=100)", 100, 49, 100 * 49 * 0.9);

// Test case 4: quantity = 50 (third tier, 25% discount)
testCase("Quantity 50 - 25% discount (price=100)", 100, 50, 100 * 50 * 0.75);