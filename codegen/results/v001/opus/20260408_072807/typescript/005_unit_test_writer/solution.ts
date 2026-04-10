function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function assert(description: string, expected: number, actual: number): void {
    if (Math.abs(expected - actual) < 0.0001) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${actual})`);
    }
}

// quantity 9 — no discount
assert("quantity 9, price 10 (no discount)", 90, calculateDiscount(10, 9));

// quantity 10 — first discount tier boundary (10%)
assert("quantity 10, price 10 (10% discount)", 90, calculateDiscount(10, 10));

// quantity 49 — still in 10% tier
assert("quantity 49, price 10 (10% discount)", 441, calculateDiscount(10, 49));

// quantity 50 — second discount tier boundary (25%)
assert("quantity 50, price 10 (25% discount)", 375, calculateDiscount(10, 50));

// quantity 1 — well below first tier
assert("quantity 1, price 20 (no discount)", 20, calculateDiscount(20, 1));

// quantity 100 — well into 25% tier
assert("quantity 100, price 5 (25% discount)", 375, calculateDiscount(5, 100));
