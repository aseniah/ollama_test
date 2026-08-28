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

function runTest(test: TestCase): void {
    const result = calculateDiscount(test.price, test.quantity);
    
    // Use a small epsilon for floating point comparison
    const epsilon = 0.0001;
    const passed = Math.abs(result - test.expected) < epsilon;
    
    if (passed) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
    }
}

// Test cases covering boundary conditions
const tests: TestCase[] = [
    {
        description: "Quantity 9 (no discount)",
        price: 100,
        quantity: 9,
        expected: 900  // 100 * 9 = 900
    },
    {
        description: "Quantity 10 (10% discount starts)",
        price: 100,
        quantity: 10,
        expected: 900  // 100 * 10 * 0.9 = 900
    },
    {
        description: "Quantity 49 (10% discount, before 25%)",
        price: 100,
        quantity: 49,
        expected: 4410 // 100 * 49 * 0.9 = 4410
    },
    {
        description: "Quantity 50 (25% discount starts)",
        price: 100,
        quantity: 50,
        expected: 3750 // 100 * 50 * 0.75 = 3750
    }
];

// Run all tests
console.log("Running calculateDiscount tests...\n");
tests.forEach(runTest);