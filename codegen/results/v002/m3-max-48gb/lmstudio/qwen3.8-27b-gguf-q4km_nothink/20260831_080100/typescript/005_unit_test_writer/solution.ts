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
    const actual = calculateDiscount(test.price, test.quantity);
    // Use a small epsilon for floating-point comparison
    const epsilon = 1e-9;
    if (Math.abs(actual - test.expected) < epsilon) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})`);
    }
}

// Test cases covering boundary conditions
const testCases: TestCase[] = [
    {
        description: "Quantity below 10 (no discount)",
        price: 10,
        quantity: 9,
        expected: 10 * 9 // 90
    },
    {
        description: "Quantity at 10 (10% discount)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9 // 90
    },
    {
        description: "Quantity below 50 but >= 10 (10% discount)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9 // 441
    },
    {
        description: "Quantity at 50 (25% discount)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75 // 375
    },
    {
        description: "Quantity above 50 (25% discount)",
        price: 5,
        quantity: 100,
        expected: 5 * 100 * 0.75 // 375
    }
];

for (const test of testCases) {
    runTest(test);
}