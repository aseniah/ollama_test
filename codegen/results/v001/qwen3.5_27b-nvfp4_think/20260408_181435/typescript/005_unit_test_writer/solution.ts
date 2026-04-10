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
    const actual = calculateDiscount(testCase.price, testCase.quantity);
    // Use a small epsilon for floating point comparison
    const isMatch = Math.abs(actual - testCase.expected) < 0.001;

    if (isMatch) {
        console.log(`PASS: ${testCase.description}`);
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${actual})`);
    }
}

// Test cases including boundary conditions
const tests: TestCase[] = [
    // Boundary condition: quantity just below 10 (no discount)
    {
        description: "Quantity 9 - no discount",
        price: 100,
        quantity: 9,
        expected: 900
    },
    // Boundary condition: quantity exactly 10 (10% discount)
    {
        description: "Quantity 10 - 10% discount boundary",
        price: 100,
        quantity: 10,
        expected: 900
    },
    // Boundary condition: quantity just below 50 (10% discount)
    {
        description: "Quantity 49 - still 10% discount",
        price: 100,
        quantity: 49,
        expected: 4410
    },
    // Boundary condition: quantity exactly 50 (25% discount)
    {
        description: "Quantity 50 - 25% discount boundary",
        price: 100,
        quantity: 50,
        expected: 3750
    },
    // Additional test cases
    {
        description: "Quantity 1 - no discount",
        price: 50,
        quantity: 1,
        expected: 50
    },
    {
        description: "Quantity 25 - 10% discount mid-range",
        price: 20,
        quantity: 25,
        expected: 450
    },
    {
        description: "Quantity 100 - 25% discount large quantity",
        price: 8,
        quantity: 100,
        expected: 600
    }
];

// Run all tests
console.log("Running calculateDiscount tests...\n");
tests.forEach(runTest);
console.log("\nTests completed.");