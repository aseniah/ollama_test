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
        description: "quantity=9 (below 10, no discount)",
        price: 100,
        quantity: 9,
        expected: 900,
    },
    {
        description: "quantity=10 (boundary: 10% discount applies)",
        price: 100,
        quantity: 10,
        expected: 900,
    },
    {
        description: "quantity=49 (below 50, 10% discount)",
        price: 100,
        quantity: 49,
        expected: 4410,
    },
    {
        description: "quantity=50 (boundary: 25% discount applies)",
        price: 100,
        quantity: 50,
        expected: 3750,
    },
    {
        description: "quantity=1 (minimum, no discount)",
        price: 25.5,
        quantity: 1,
        expected: 25.5,
    },
    {
        description: "quantity=100 (well above 50, 25% discount)",
        price: 20,
        quantity: 100,
        expected: 1500,
    },
];

let passed = 0;
let failed = 0;

for (const tc of testCases) {
    const got = calculateDiscount(tc.price, tc.quantity);

    // Use a small epsilon to handle floating-point comparison
    const epsilon = 1e-9;
    if (Math.abs(got - tc.expected) < epsilon) {
        console.log(`PASS: ${tc.description}`);
        passed++;
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${got})`);
        failed++;
    }
}

console.log(`\n${passed} passed, ${failed} failed out of ${testCases.length} test cases.`);