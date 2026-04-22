function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTests() {
    const tests = [
        {
            description: "Quantity 9 (no discount)",
            price: 10,
            quantity: 9,
            expected: 10 * 9
        },
        {
            description: "Quantity 10 (10% discount)",
            price: 10,
            quantity: 10,
            expected: 10 * 10 * 0.9
        },
        {
            description: "Quantity 49 (10% discount)",
            price: 10,
            quantity: 49,
            expected: 10 * 49 * 0.9
        },
        {
            description: "Quantity 50 (25% discount)",
            price: 10,
            quantity: 50,
            expected: 10 * 50 * 0.75
        }
    ];

    for (const test of tests) {
        const result = calculateDiscount(test.price, test.quantity);
        const expected = test.expected;
        // Use a small epsilon for floating point comparison
        const epsilon = 0.0001;
        if (Math.abs(result - expected) < epsilon) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${expected}, got: ${result})`);
        }
    }
}

runTests();