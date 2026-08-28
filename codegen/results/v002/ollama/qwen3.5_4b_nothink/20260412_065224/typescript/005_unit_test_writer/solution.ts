function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
function runTests() {
    const cases = [
        { price: 10, quantity: 9, expected: 90, desc: "No discount, quantity less than 10" },
        { price: 10, quantity: 10, expected: 90, desc: "10% discount, quantity 10" },
        { price: 10, quantity: 49, expected: 441, desc: "10% discount, quantity 49" },
        { price: 10, quantity: 50, expected: 375, desc: "25% discount, quantity 50" },
    ];

    for (const { price, quantity, expected, desc } of cases) {
        const result = calculateDiscount(price, quantity);
        if (result === expected) {
            console.log(`PASS: ${desc}`);
        } else {
            console.log(`FAIL: ${desc} (expected: ${expected}, got: ${result})`);
        }
    }
}

runTests();