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

const tests: TestCase[] = [
    { description: "quantity 9 (no discount)", price: 10, quantity: 9, expected: 90 },
    { description: "quantity 10 (10% discount boundary)", price: 10, quantity: 10, expected: 90 },
    { description: "quantity 49 (10% discount upper boundary)", price: 10, quantity: 49, expected: 441 },
    { description: "quantity 50 (25% discount boundary)", price: 10, quantity: 50, expected: 375 },
    { description: "quantity 1 (no discount, single item)", price: 20, quantity: 1, expected: 20 },
    { description: "quantity 100 (25% discount)", price: 5, quantity: 100, expected: 375 },
];

for (const t of tests) {
    const got = calculateDiscount(t.price, t.quantity);
    if (Math.abs(got - t.expected) < 0.0001) {
        console.log(`PASS: ${t.description}`);
    } else {
        console.log(`FAIL: ${t.description} (expected: ${t.expected}, got: ${got})`);
    }
}
