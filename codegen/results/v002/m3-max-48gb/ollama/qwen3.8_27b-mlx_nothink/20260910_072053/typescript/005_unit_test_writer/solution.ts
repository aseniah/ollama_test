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
        description: "Quantity 9 (no discount)",
        price: 10,
        quantity: 9,
        expected: 10 * 9, // 90
    },
    {
        description: "Quantity 10 (10% discount)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9, // 90
    },
    {
        description: "Quantity 49 (10% discount)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9, // 441
    },
    {
        description: "Quantity 50 (25% discount)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75, // 375
    },
    {
        description: "Quantity 1 (no discount, single item)",
        price: 5,
        quantity: 1,
        expected: 5 * 1, // 5
    },
    {
        description: "Quantity 100 (25% discount, large quantity)",
        price: 20,
        quantity: 100,
        expected: 20 * 100 * 0.75, // 1500
    },
];

for (const test of testCases) {
    const result = calculateDiscount(test.price, test.quantity);
    // Use a small epsilon for floating point comparison
    if (Math.abs(result - test.expected) < 1e-9) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
    }
}