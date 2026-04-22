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
        description: "quantity 9 (no discount, just under threshold)",
        price: 100,
        quantity: 9,
        expected: 900,
    },
    {
        description: "quantity 10 (first tier discount applied)",
        price: 100,
        quantity: 10,
        expected: 900,
    },
    {
        description: "quantity 49 (mid-tier discount applied)",
        price: 100,
        quantity: 49,
        expected: 4410,
    },
    {
        description: "quantity 50 (highest discount applied)",
        price: 100,
        quantity: 50,
        expected: 3750,
    },
];

function areEqual(a: number, b: number): boolean {
    return Math.abs(a - b) < 1e-9;
}

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    if (areEqual(result, tc.expected)) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
    }
}