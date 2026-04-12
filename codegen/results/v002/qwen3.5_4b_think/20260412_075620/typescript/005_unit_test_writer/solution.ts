function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases with boundary conditions
const testCases = [
    { price: 10, quantity: 9, expected: 90, description: "No discount (quantity = 9 < 10)" },
    { price: 10, quantity: 10, expected: 90, description: "10% discount (quantity = 10, 10 <= quantity < 50)" },
    { price: 10, quantity: 49, expected: 441, description: "10% discount (quantity = 49, 49 < 50)" },
    { price: 10, quantity: 50, expected: 375, description: "25% discount (quantity = 50, quantity >= 50)" }
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