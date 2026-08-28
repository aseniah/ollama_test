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
        description: "quantity 9 (below 10, no discount)",
        price: 10,
        quantity: 9,
        expected: 10 * 9, // 90
    },
    {
        description: "quantity 10 (boundary, 10% discount)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9, // 90
    },
    {
        description: "quantity 49 (just below 50, 10% discount)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9, // 441
    },
    {
        description: "quantity 50 (boundary, 25% discount)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75, // 375
    },
    {
        description: "quantity 1, price 5 (simple case)",
        price: 5,
        quantity: 1,
        expected: 5 * 1, // 5
    },
    {
        description: "quantity 0 (zero quantity)",
        price: 10,
        quantity: 0,
        expected: 10 * 0, // 0
    },
];

let passed = 0;
let failed = 0;

for (const tc of testCases) {
    const got = calculateDiscount(tc.price, tc.quantity);
    if (Math.abs(got - tc.expected) < 1e-9) {
        console.log(`PASS: ${tc.description}`);
        passed++;
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${got})`);
        failed++;
    }
}

console.log(`\n${passed} passed, ${failed} failed, ${testCases.length} total`);