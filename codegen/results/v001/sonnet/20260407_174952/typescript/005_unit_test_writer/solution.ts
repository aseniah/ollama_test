function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function assert(description: string, expected: number, got: number): void {
    if (Math.abs(expected - got) < 1e-9) {
        console.log(`PASS: ${description}`);
    } else {
        console.log(`FAIL: ${description} (expected: ${expected}, got: ${got})`);
    }
}

// quantity < 10: no discount
assert("quantity 9 — no discount", 10 * 9, calculateDiscount(10, 9));

// quantity >= 10: 10% discount (boundary)
assert("quantity 10 — 10% discount", 10 * 10 * 0.9, calculateDiscount(10, 10));

// quantity 49: still 10% discount (just below 50)
assert("quantity 49 — 10% discount", 10 * 49 * 0.9, calculateDiscount(10, 49));

// quantity >= 50: 25% discount (boundary)
assert("quantity 50 — 25% discount", 10 * 50 * 0.75, calculateDiscount(10, 50));

// quantity 1: well below first threshold
assert("quantity 1 — no discount", 10 * 1, calculateDiscount(10, 1));

// quantity 100: well above top threshold
assert("quantity 100 — 25% discount", 10 * 100 * 0.75, calculateDiscount(10, 100));
