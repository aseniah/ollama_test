function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
const testCases = [
    {
        description: "Quantity 9 (no discount)",
        price: 10,
        quantity: 9,
        expected: 90
    },
    {
        description: "Quantity 10 (90% discount)",
        price: 10,
        quantity: 10,
        expected: 90
    },
    {
        description: "Quantity 49 (90% discount)",
        price: 10,
        quantity: 49,
        expected: 441
    },
    {
        description: "Quantity 50 (75% discount)",
        price: 10,
        quantity: 50,
        expected: 375
    }
];

// Run tests
for (const test of testCases) {
    const result = calculateDiscount(test.price, test.quantity);
    if (result === test.expected) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
    }
}