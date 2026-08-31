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

function runTest(test: TestCase): void {
    const result = calculateDiscount(test.price, test.quantity);
    // Use a small epsilon for floating point comparison
    const epsilon = 1e-9;
    if (Math.abs(result - test.expected) < epsilon) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
    }
}

const testCases: TestCase[] = [
    {
        description: "Quantity 9 (below 10, no discount)",
        price: 10,
        quantity: 9,
        expected: 10 * 9,
    },
    {
        description: "Quantity 10 (boundary, 10% discount)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9,
    },
    {
        description: "Quantity 49 (below 50, 10% discount)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9,
    },
    {
        description: "Quantity 50 (boundary, 25% discount)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75,
    },
    {
        description: "Quantity 0 (no items, no discount)",
        price: 10,
        quantity: 0,
        expected: 0,
    },
    {
        description: "Quantity 100 (well above 50, 25% discount)",
        price: 10,
        quantity: 100,
        expected: 10 * 100 * 0.75,
    },
];

for (const test of testCases) {
    runTest(test);
}