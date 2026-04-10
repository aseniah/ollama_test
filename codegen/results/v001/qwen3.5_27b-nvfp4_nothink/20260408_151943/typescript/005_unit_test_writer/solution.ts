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

const tests: TestCase[] = [
    { 
        description: "Quantity below 10 (no discount)", 
        price: 10, 
        quantity: 9, 
        expected: 90 // 10 * 9
    },
    { 
        description: "Boundary condition: Quantity exactly 10 (10% discount)", 
        price: 10, 
        quantity: 10, 
        expected: 90 // 10 * 10 * 0.9
    },
    { 
        description: "Quantity between 10 and 50 (boundary below 50)", 
        price: 20, 
        quantity: 49, 
        expected: 882 // 20 * 49 * 0.9
    },
    { 
        description: "Boundary condition: Quantity exactly 50 (25% discount)", 
        price: 20, 
        quantity: 50, 
        expected: 750 // 20 * 50 * 0.75
    },
    {
        description: "Quantity above 50",
        price: 100,
        quantity: 60,
        expected: 4500 // 100 * 60 * 0.75
    }
];

tests.forEach((test) => {
    const result = calculateDiscount(test.price, test.quantity);
    // Use a small epsilon for floating point comparison safety, though these specific values are clean integers/decimals
    const isPass = Math.abs(result - test.expected) < 0.00001;

    if (isPass) {
        console.log(`PASS: ${test.description}`);
    } else {
        console.log(`FAIL: ${test.description} (expected: ${test.expected}, got: ${result})`);
    }
});