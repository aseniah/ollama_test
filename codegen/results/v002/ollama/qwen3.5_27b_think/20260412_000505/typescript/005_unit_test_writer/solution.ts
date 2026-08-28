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
        description: "No discount for quantity 9 (less than 10)",
        price: 10,
        quantity: 9,
        expected: 10 * 9 // 90
    },
    {
        description: "10% discount for quantity 10 (boundary)",
        price: 10,
        quantity: 10,
        expected: 10 * 10 * 0.9 // 90
    },
    {
        description: "10% discount for quantity 49 (boundary before 50)",
        price: 10,
        quantity: 49,
        expected: 10 * 49 * 0.9 // 441
    },
    {
        description: "25% discount for quantity 50 (boundary)",
        price: 10,
        quantity: 50,
        expected: 10 * 50 * 0.75 // 375
    }
];

for (const testCase of testCases) {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    const expected = testCase.expected;
    
    if (result === expected) {
        console.log(`PASS: ${testCase.description}`);
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${expected}, got: ${result})`);
    }
}