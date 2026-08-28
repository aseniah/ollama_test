function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    description: string;
    price: number;
    quantity: number;
    expected: number;
}

function runTest(tc: TestCase): void {
    const got = calculateDiscount(tc.price, tc.quantity);
    if (Math.abs(got - tc.expected) < 0.0001) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${got})`);
    }
}

const tests: TestCase[] = [
    { description: "quantity 9 (below 10, no discount)", price: 10, quantity: 9, expected: 90 },
    { description: "quantity 10 (boundary, 10% discount)", price: 10, quantity: 10, expected: 90 },
    { description: "quantity 49 (below 50, 10% discount)", price: 10, quantity: 49, expected: 441 },
    { description: "quantity 50 (boundary, 25% discount)", price: 10, quantity: 50, expected: 375 },
    { description: "quantity 1 (single item, no discount)", price: 20, quantity: 1, expected: 20 },
    { description: "quantity 100 (large order, 25% discount)", price: 5, quantity: 100, expected: 375 },
];

for (const tc of tests) {
    runTest(tc);
}
