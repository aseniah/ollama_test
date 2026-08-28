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

const testCases: TestCase[] = [
    {
        description: "quantity below 10 (no discount)",
        price: 10,
        quantity: 9,
        expected: 10 * 9,
    },
    {
        description: "quantity exactly 10 (10% discount)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9,
    },
    {
        description: "quantity below 50 (10% discount)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9,
    },
    {
        description: "quantity exactly 50 (25% discount)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75,
    },
];

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    if (result === tc.expected) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
    }
}