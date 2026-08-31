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
    const got = calculateDiscount(price, quantity);
    const tolerance = 1e-9;

    if (Math.abs(got - expected) < tolerance) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${got})`);
    }
}

// Boundary: just below 10 (no discount)
runTest("price=10, quantity=9 (no discount)", 10, 9, 10 * 9);

// Boundary: exactly 10 (10% discount kicks in)
runTest("price=10, quantity=10 (10% discount)", 10, 10, 10 * 10 * 0.9);

// Boundary: just below 50 (still 10% discount)
runTest("price=10, quantity=49 (10% discount)", 10, 49, 10 * 49 * 0.9);

// Boundary: exactly 50 (25% discount kicks in)
runTest("price=10, quantity=50 (25% discount)", 10, 50, 10 * 50 * 0.75);

// Extra: mid-range with different price
runTest("price=25, quantity=20 (10% discount)", 25, 20, 25 * 20 * 0.9);

// Extra: large quantity with different price
runTest("price=5, quantity=100 (25% discount)", 5, 100, 5 * 100 * 0.75);