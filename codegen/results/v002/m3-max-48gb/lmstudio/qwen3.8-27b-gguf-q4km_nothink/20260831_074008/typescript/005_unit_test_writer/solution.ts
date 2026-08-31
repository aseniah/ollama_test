function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test helper
function test(description: string, price: number, quantity: number, expected: number): void {
    const result = calculateDiscount(price, quantity);
    // Use a small epsilon for floating point comparison
    const epsilon = 1e-9;
    if (Math.abs(result - expected) < epsilon) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${result})`);
    }
}

// Test case 1: quantity 9 (below 10, no discount)
test("quantity 9 (no discount)", 100, 9, 900);

// Test case 2: quantity 10 (boundary, 10% discount)
test("quantity 10 (10% discount)", 100, 10, 900);

// Test case 3: quantity 49 (below 50, 10% discount)
test("quantity 49 (10% discount)", 100, 49, 4410);

// Test case 4: quantity 50 (boundary, 25% discount)
test("quantity 50 (25% discount)", 100, 50, 3750);