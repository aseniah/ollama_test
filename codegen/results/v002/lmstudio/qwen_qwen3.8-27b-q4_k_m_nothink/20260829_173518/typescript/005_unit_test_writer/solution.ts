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
        expected: 10 * 9
    },
    {
        description: "quantity 10 (boundary, 10% discount)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9
    },
    {
        description: "quantity 49 (below 50, 10% discount)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9
    },
    {
        description: "quantity 50 (boundary, 25% discount)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75
    },
    {
        description: "quantity 100 (well above 50, 25% discount)",
        price: 5,
        quantity: 100,
        expected: 5 * 100 * 0.75
    }
];

function approximatelyEqual(a: number, b: number, epsilon: number = 1e-9): boolean {
    return Math.abs(a - b) < epsilon;
}

for (const tc of testCases) {
    const got = calculateDiscount(tc.price, tc.quantity);
    if (approximatelyEqual(got, tc.expected)) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${got})`);
    }
}