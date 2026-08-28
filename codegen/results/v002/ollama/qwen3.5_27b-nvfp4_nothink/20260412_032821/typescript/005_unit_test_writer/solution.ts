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

const tests: TestCase[] = [
    {
        description: "Quantity 9 (No discount)",
        price: 10,
        quantity: 9,
        expected: 90.0
    },
    {
        description: "Quantity 10 (10% discount boundary)",
        price: 10,
        quantity: 10,
        expected: 90.0 // 10 * 10 * 0.9
    },
    {
        description: "Quantity 49 (10% discount upper boundary)",
        price: 20,
        quantity: 49,
        expected: 882.0 // 20 * 49 * 0.9
    },
    {
        description: "Quantity 50 (25% discount boundary)",
        price: 20,
        quantity: 50,
        expected: 750.0 // 20 * 50 * 0.75
    }
];

console.log("Running tests for calculateDiscount...\n");

tests.forEach(runTest);

console.log("\nTests complete.");