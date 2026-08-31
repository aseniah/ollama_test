function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function test(description: string, price: number, quantity: number, expected: number): void {
    const got = calculateDiscount(price, quantity);
    const tolerance = 1e-9;
    if (Math.abs(got - expected) < tolerance) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${got})`);
    }
}

// Boundary and key test cases
test("quantity 9 (below 10, no discount)", 10, 9, 10 * 9);
test("quantity 10 (boundary, 10% discount)", 10, 10, 10 * 10 * 0.9);
test("quantity 49 (below 50, 10% discount)", 10, 49, 10 * 49 * 0.9);
test("quantity 50 (boundary, 25% discount)", 10, 50, 10 * 50 * 0.75);
test("quantity 1 (minimum, no discount)", 5, 1, 5 * 1);
test("quantity 100 (above 50, 25% discount)", 3, 100, 3 * 100 * 0.75);