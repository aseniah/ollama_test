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
        description: "quantity 9 (no discount, below threshold)",
        price: 10,
        quantity: 9,
        expected: 10 * 9, // 90
    },
    {
        description: "quantity 10 (boundary: 10% discount applies)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9, // 90
    },
    {
        description: "quantity 49 (boundary: still 10% discount)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9, // 441
    },
    {
        description: "quantity 50 (boundary: 25% discount applies)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75, // 375
    },
    {
        description: "quantity 1 (no discount, single item)",
        price: 5,
        quantity: 1,
        expected: 5 * 1, // 5
    },
    {
        description: "quantity 100 (25% discount, above 50)",
        price: 20,
        quantity: 100,
        expected: 20 * 100 * 0.75, // 1500
    },
];

let passed = 0;
let failed = 0;

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    // Use a small epsilon for floating-point comparison
    const epsilon = 1e-9;
    if (Math.abs(result - tc.expected) < epsilon) {
        console.log(`PASS: ${tc.description}`);
        passed++;
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
        failed++;
    }
}

console.log(`\n--- Results: ${passed} passed, ${failed} failed, ${testCases.length} total ---`);