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
        description: "quantity 9 (no discount)",
        price: 10,
        quantity: 9,
        expected: 90
    },
    {
        description: "quantity 10 (10% discount boundary)",
        price: 10,
        quantity: 10,
        expected: 90
    },
    {
        description: "quantity 49 (10% discount boundary)",
        price: 10,
        quantity: 49,
        expected: 441
    },
    {
        description: "quantity 50 (25% discount boundary)",
        price: 10,
        quantity: 50,
        expected: 375
    },
    {
        description: "quantity 1 (no discount, single item)",
        price: 5,
        quantity: 1,
        expected: 5
    },
    {
        description: "quantity 100 (25% discount, large quantity)",
        price: 20,
        quantity: 100,
        expected: 1500
    }
];

function compareWithTolerance(expected: number, got: number): boolean {
    const tolerance = 1e-9;
    return Math.abs(expected - got) < tolerance;
}

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    if (compareWithTolerance(tc.expected, result)) {
        console.log(`PASS: ${tc.description}`);
    } else {
        console.log(`FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`);
    }
}