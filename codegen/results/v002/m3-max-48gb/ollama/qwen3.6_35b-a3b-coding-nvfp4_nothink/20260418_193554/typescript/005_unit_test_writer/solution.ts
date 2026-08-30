function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
const testCases = [
    {
        description: "Quantity 9 (no discount)",
        price: 100,
        quantity: 9,
        expected: 900
    },
    {
        description: "Quantity 10 (10% discount)",
        price: 100,
        quantity: 10,
        expected: 900
    },
    {
        description: "Quantity 49 (10% discount)",
        price: 100,
        quantity: 49,
        expected: 4410
    },
    {
        description: "Quantity 50 (25% discount)",
        price: 100,
        quantity: 50,
        expected: 3750
    }
];

for (const tc of testCases) {
    const got = calculateDiscount(tc.price, tc.quantity);
    if (got === tc.expected) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${got})`);
    }
}