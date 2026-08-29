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
        description: "quantity 9 has no discount",
        price: 100,
        quantity: 9,
        expected: 900
    },
    {
        description: "quantity 10 gets 10% discount",
        price: 100,
        quantity: 10,
        expected: 900
    },
    {
        description: "quantity 49 gets 10% discount",
        price: 100,
        quantity: 49,
        expected: 4410
    },
    {
        description: "quantity 50 gets 25% discount",
        price: 100,
        quantity: 50,
        expected: 3750
    }
];

const EPSILON = 1e-9;

for (const test of testCases) {
    const got = calculateDiscount(test.price, test.quantity);

    if (Math.abs(got - test.expected) <= EPSILON) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${got})`);
    }
}