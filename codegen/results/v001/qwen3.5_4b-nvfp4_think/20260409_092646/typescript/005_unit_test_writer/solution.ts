function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    name: string;
    description: string;
    expected: number;
    actual?: number;
    inputPrice?: number;
    inputQuantity: number;
}

function runTest(test: TestCase) {
    const result = calculateDiscount(test.inputPrice ?? 10, test.inputQuantity);
    
    if (result === test.expected) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
    }
}

const tests: TestCase[] = [
    {
        name: "Quantity 9",
        description: "Single-digit quantity should use no discount",
        expected: 10 * 9,
        inputPrice: 10,
        inputQuantity: 9
    },
    {
        name: "Quantity 10",
        description: "Just above single digit should use 10% discount",
        expected: 10 * 10 * 0.9,
        inputPrice: 10,
        inputQuantity: 10
    },
    {
        name: "Quantity 49",
        description: "Just below 50 should use 10% discount",
        expected: 10 * 49 * 0.9,
        inputPrice: 10,
        inputQuantity: 49
    },
    {
        name: "Quantity 50",
        description: "Exactly 50 should use 25% discount",
        expected: 10 * 50 * 0.75,
        inputPrice: 10,
        inputQuantity: 50
    },
    {
        name: "Zero quantity",
        description: "Edge case - zero quantity should return 0",
        expected: 10 * 0,
        inputPrice: 10,
        inputQuantity: 0
    },
    {
        name: "Large quantity 99",
        description: "Large quantity beyond threshold should use 25% discount",
        expected: 10 * 99 * 0.75,
        inputPrice: 10,
        inputQuantity: 99
    }
];

// Run tests
console.log("Running Discount Calculation Tests\n");
console.log("=" .repeat(40));
for (const test of tests) {
    runTest(test);
}
console.log("=" .repeat(40));
console.log("Tests completed.");