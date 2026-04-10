function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

interface TestCase {
    price: number;
    quantity: number;
    expected: number;
    description: string;
}

const testCases: TestCase[] = [
    {
        price: 10,
        quantity: 9,
        expected: 90,
        description: "Quantity 9 (below 10, no discount)"
    },
    {
        price: 10,
        quantity: 10,
        expected: 90,
        description: "Quantity 10 (boundary, 10% discount)"
    },
    {
        price: 10,
        quantity: 49,
        expected: 441,
        description: "Quantity 49 (below 50, 10% discount)"
    },
    {
        price: 10,
        quantity: 50,
        expected: 375,
        description: "Quantity 50 (boundary, 25% discount)"
    }
];

let passCount = 0;
let failCount = 0;

testCases.forEach((testCase) => {
    const result = calculateDiscount(testCase.price, testCase.quantity);
    // Using a small epsilon for floating point comparison safety, though these specific inputs result in clean decimals
    const epsilon = 0.000001;
    const isPass = Math.abs(result - testCase.expected) < epsilon;

    if (isPass) {
        console.log(`PASS: ${testCase.description}`);
        passCount++;
    } else {
        console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
        failCount++;
    }
});

console.log(`\nTotal: ${passCount} passed, ${failCount} failed`);