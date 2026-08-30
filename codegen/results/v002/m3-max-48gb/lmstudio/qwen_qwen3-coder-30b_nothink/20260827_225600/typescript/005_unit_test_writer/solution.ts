function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function testCalculateDiscount(): void {
    const testCases = [
        {
            description: "quantity 9 (below threshold)",
            price: 10,
            quantity: 9,
            expected: 90
        },
        {
            description: "quantity 10 (first threshold)",
            price: 10,
            quantity: 10,
            expected: 90
        },
        {
            description: "quantity 49 (below second threshold)",
            price: 10,
            quantity: 49,
            expected: 441
        },
        {
            description: "quantity 50 (second threshold)",
            price: 10,
            quantity: 50,
            expected: 375
        },
        {
            description: "quantity 100 (above second threshold)",
            price: 5,
            quantity: 100,
            expected: 375
        }
    ];

    testCases.forEach(testCase => {
        const result = calculateDiscount(testCase.price, testCase.quantity);
        const passed = Math.abs(result - testCase.expected) < 0.0001; // Handle floating point precision
        
        if (passed) {
            console.log(`PASS: ${testCase.description}`);
        } else {
            console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
        }
    });
}

testCalculateDiscount();