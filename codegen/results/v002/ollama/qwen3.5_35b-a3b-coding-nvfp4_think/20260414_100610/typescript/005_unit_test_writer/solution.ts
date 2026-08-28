function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

type TestCase = {
    description: string;
    price: number;
    quantity: number;
    expected: number;
};

const testCases: TestCase[] = [
    {
        description: "Quantity 9 (no discount)",
        price: 100,
        quantity: 9,
        expected: 100 * 9
    },
    {
        description: "Quantity 10 (10% discount)",
        price: 100,
        quantity: 10,
        expected: 100 * 10 * 0.9
    },
    {
        description: "Quantity 49 (10% discount)",
        price: 100,
        quantity: 49,
        expected: 100 * 49 * 0.9
    },
    {
        description: "Quantity 50 (25% discount)",
        price: 100,
        quantity: 50,
        expected: 100 * 50 * 0.75
    },
    {
        description: "Quantity 100 (25% discount)",
        price: 50,
        quantity: 100,
        expected: 50 * 100 * 0.75
    }
];

function runTests(): void {
    let passed = 0;
    let failed = 0;

    for (const testCase of testCases) {
        const result = calculateDiscount(testCase.price, testCase.quantity);
        const expected = testCase.expected;
        
        // Use a small epsilon for floating point comparison
        const epsilon = 0.0001;
        const isClose = Math.abs(result - expected) < epsilon;

        if (isClose) {
            console.log(`PASS: ${testCase.description}`);
            passed++;
        } else {
            console.log(`FAIL: ${testCase.description} (expected: ${expected}, got: ${result})`);
            failed++;
        }
    }

    console.log(`\nResults: ${passed} passed, ${failed} failed`);
}

runTests();