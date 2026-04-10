function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Run tests
const tests = [
    {
        desc: "Quantity less than 10 (boundary)",
        price: 10,
        quantity: 9,
        expected: 90,
    },
    {
        desc: "Quantity at 10 (boundary)",
        price: 10,
        quantity: 10,
        expected: 90,
    },
    {
        desc: "Quantity at 49 (boundary)",
        price: 10,
        quantity: 49,
        expected: 441, // 49 * 10 * 0.9
    },
    {
        desc: "Quantity at 50 (boundary)",
        price: 10,
        quantity: 50,
        expected: 375, // 50 * 10 * 0.75
    },
];

let failed = false;

for (const test of tests) {
    const result = calculateDiscount(test.price, test.quantity);
    if (result === test.expected) {
        console.log(`PASS: ${test.desc}`);
    } else {
        console.log(`FAIL: ${test.desc} (expected: ${test.expected}, got: ${result})`);
        failed = true;
    }
}

if (failed) {
    process.exit(1);
}