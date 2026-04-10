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
        price: 10,
        quantity: 9,
        expected: 90
    },
    {
        description: "Quantity 10 (10% discount)",
        price: 10,
        quantity: 10,
        expected: 90 // 10 * 10 * 0.9
    },
    {
        description: "Quantity 49 (10% discount)",
        price: 10,
        quantity: 49,
        expected: 441 // 10 * 49 * 0.9
    },
    {
        description: "Quantity 50 (25% discount)",
        price: 10,
        quantity: 50,
        expected: 375 // 10 * 50 * 0.75
    }
];

function runTests(tests: TestCase[]): void {
    tests.forEach((test) => {
        const result = calculateDiscount(test.price, test.quantity);
        // Using a small epsilon for floating point comparison safety
        const epsilon = 0.0001;
        const passed = Math.abs(result - test.expected) < epsilon;

        if (passed) {
            console.log(`PASS: ${test.description}`);
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
        }
    });
}

runTests(testCases);