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

function runTest(testCase: TestCase): void {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    
    // Use a small epsilon for floating point comparison
    const epsilon = 0.001;
    if (Math.abs(result - testCase.expected) < epsilon) {
        console.log(`PASS: ${testCase.description}`);
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
    }
}

// Test cases including boundary conditions (9, 10, 49, 50)
const testCases: TestCase[] = [
    { description: "Quantity 9 (no discount)", price: 10, quantity: 9, expected: 90 },
    { description: "Quantity 10 (10% discount boundary)", price: 10, quantity: 10, expected: 90 },
    { description: "Quantity 49 (10% discount)", price: 10, quantity: 49, expected: 441 },
    { description: "Quantity 50 (25% discount boundary)", price: 10, quantity: 50, expected: 375 },
    { description: "Quantity 100 (25% discount)", price: 8, quantity: 100, expected: 600 },
];

console.log("Running tests...\n");

testCases.forEach(runTest);