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
        price: 100,
        quantity: 9,
        expected: 900
    },
    {
        description: "Quantity 10 (10% discount)",
        price: 100,
        quantity: 10,
        expected: 900
    },
    {
        description: "Quantity 49 (10% discount)",
        price: 100,
        quantity: 49,
        expected: 4410
    },
    {
        description: "Quantity 50 (25% discount)",
        price: 100,
        quantity: 50,
        expected: 3750
    }
];

let passCount = 0;
let failCount = 0;

console.log("Running calculateDiscount tests...\n");

for (const test of testCases) {
    const result = calculateDiscount(test.price, test.quantity);
    if (result === test.expected) {
        console.log(`PASS: ${test.description}`);
        passCount++;
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
        failCount++;
    }
}

console.log(`\nSummary: ${passCount} passed, ${failCount} failed`);

process.exit(failCount > 0 ? 1 : 0);