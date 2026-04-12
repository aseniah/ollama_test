// Copy of the provided function
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

function runTests(tests: TestCase[]): void {
    let passedCount = 0;
    let failedCount = 0;

    for (const test of tests) {
        const result = calculateDiscount(test.price, test.quantity);
        
        // Use a small epsilon for floating point comparison
        const epsilon = 0.0001;
        const isPass = Math.abs(result - test.expected) < epsilon;

        if (isPass) {
            console.log(`PASS: ${test.description}`);
            passedCount++;
        } else {
            console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
            failedCount++;
        }
    }

    console.log(`\nResults: ${passedCount} passed, ${failedCount} failed`);
}

// Define test cases covering boundary conditions and general logic
const testCases: TestCase[] = [
    {
        description: "Quantity below 10 (no discount)",
        price: 10,
        quantity: 9,
        expected: 90
    },
    {
        description: "Boundary at 10 (10% discount)",
        price: 10,
        quantity: 10,
        expected: 90
    },
    {
        description: "Boundary below 50 (10% discount)",
        price: 10,
        quantity: 49,
        expected: 441
    },
    {
        description: "Boundary at 50 (25% discount)",
        price: 10,
        quantity: 50,
        expected: 375
    },
    {
        description: "Large quantity (25% discount)",
        price: 20,
        quantity: 100,
        expected: 1500
    }
];

// Execute the tests
runTests(testCases);