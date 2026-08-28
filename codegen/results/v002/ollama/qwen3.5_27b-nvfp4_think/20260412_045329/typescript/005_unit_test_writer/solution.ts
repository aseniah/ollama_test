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
        expected: 900  // 100 * 9 = 900 (no discount)
    },
    {
        description: "Quantity 10 (10% discount boundary)",
        price: 100,
        quantity: 10,
        expected: 900  // 100 * 10 * 0.9 = 900
    },
    {
        description: "Quantity 49 (10% discount)",
        price: 100,
        quantity: 49,
        expected: 4410 // 100 * 49 * 0.9 = 4410
    },
    {
        description: "Quantity 50 (25% discount boundary)",
        price: 100,
        quantity: 50,
        expected: 3750 // 100 * 50 * 0.75 = 3750
    },
    {
        description: "Quantity 100 (25% discount)",
        price: 100,
        quantity: 100,
        expected: 7500 // 100 * 100 * 0.75 = 7500
    }
];

function runTest(testCase: TestCase): void {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    
    // Use a small epsilon for floating point comparison
    const epsilon = 0.0001;
    const passed = Math.abs(result - testCase.expected) < epsilon;

    if (passed) {
        console.log(`PASS: ${testCase.description}`);
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
    }
}

// Run all test cases
testCases.forEach(runTest);