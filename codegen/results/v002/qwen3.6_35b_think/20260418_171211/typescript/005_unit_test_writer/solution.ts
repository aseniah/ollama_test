function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function test(
    description: string,
    price: number,
    quantity: number,
    expected: number
): void {
    const result = calculateDiscount(price, quantity);
    // Use a small epsilon for reliable floating-point comparison
    const pass = Math.abs(result - expected) < 1e-9;

    if (pass) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${result})`);
    }
}

// Test cases covering the specified boundary conditions
test("Quantity 9 (below first threshold)", 100, 9, 900);
test("Quantity 10 (at 10% discount threshold)", 100, 10, 900);
test("Quantity 49 (just under 25% discount threshold)", 100, 49, 4410);
test("Quantity 50 (at 25% discount threshold)", 100, 50, 3750);