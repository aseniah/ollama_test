function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTests() {
    const testCases = [
        { desc: "Quantity 9 (no discount)", price: 10, quantity: 9, expected: 90 },
        { desc: "Quantity 10 (10% discount)", price: 10, quantity: 10, expected: 90 },
        { desc: "Quantity 49 (10% discount)", price: 10, quantity: 49, expected: 441 },
        { desc: "Quantity 50 (25% discount)", price: 10, quantity: 50, expected: 375 },
        { desc: "Quantity 100 (25% discount)", price: 5, quantity: 100, expected: 375 }
    ];

    testCases.forEach(tc => {
        const got = calculateDiscount(tc.price, tc.quantity);
        // Using a small epsilon for floating point comparison
        const isPass = Math.abs(got - tc.expected) < 1e-9;
        if (isPass) {
            console.log(`PASS: ${tc.desc}`);
        } else {
            console.log(`FAIL: ${tc.desc} (expected: ${tc.expected}, got: ${got})`);
        }
    });
}

runTests();