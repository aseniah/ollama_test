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
        price: 100,
        quantity: 9,
        expected: 100 * 9, // 900
    },
    {
        description: "quantity 10 (boundary, 10% discount)",
        price: 100,
        quantity: 10,
        expected: 100 * 10 * 0.9, // 900
    },
    {
        description: "quantity 49 (below 50, 10% discount)",
        price: 100,
        quantity: 49,
        expected: 100 * 49 * 0.9, // 4410
    },
    {
        description: "quantity 50 (boundary, 25% discount)",
        price: 100,
        quantity: 50,
        expected: 100 * 50 * 0.75, // 3750
    },
    {
        description: "quantity 1 (minimum, no discount)",
        price: 25,
        quantity: 1,
        expected: 25 * 1, // 25
    },
    {
        description: "quantity 100 (well above 50, 25% discount)",
        price: 20,
        quantity: 100,
        expected: 20 * 100 * 0.75, // 1500
    },
];

let passed = 0;
let failed = 0;

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    const isPass = Math.abs(result - tc.expected) < 1e-10;

    if (isPass) {
        console.log(`PASS: ${tc.description}`);
        passed++;
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
        failed++;
    }
}

console.log(`\n${passed + failed} tests: ${passed} passed, ${failed} failed`);

if (failed > 0) {
    process.exit(1);
}