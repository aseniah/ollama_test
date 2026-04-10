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
        description: "No discount for quantity 9 (boundary)",
        price: 10,
        quantity: 9,
        expected: 90 // 10 * 9
    },
    {
        description: "10% discount for quantity 10 (boundary)",
        price: 10,
        quantity: 10,
        expected: 90 // 10 * 10 * 0.9
    },
    {
        description: "10% discount for quantity 49 (boundary)",
        price: 10,
        quantity: 49,
        expected: 441 // 10 * 49 * 0.9
    },
    {
        description: "25% discount for quantity 50 (boundary)",
        price: 10,
        quantity: 50,
        expected: 375 // 10 * 50 * 0.75
    }
];

function runTests(tests: TestCase[]): void {
    let passCount = 0;
    let failCount = 0;

    for (const test of tests) {
        const actual = calculateDiscount(test.price, test.quantity);
        // Using a small epsilon for floating point comparison safety, though these are clean numbers
        const epsilon = 0.0001;
        const isPass = Math.abs(actual - test.expected) < epsilon;

        if (isPass) {
            console.log(`PASS: ${test.description}`);
            passCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${actual})`);
            failCount++;
        }
    }

    console.log(`\nSummary: ${passCount} passed, ${failCount} failed`);
}

// Run the tests
runTests(testCases);