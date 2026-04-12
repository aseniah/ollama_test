function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTests(): void {
    const testCases = [
        { description: "Quantity less than 10", price: 10, quantity: 9, expected: 90 },
        { description: "Boundary condition: Quantity equals 10", price: 10, quantity: 10, expected: 100 },
        { description: "Quantity between 10 and 50", price: 10, quantity: 49, expected: 441 },
        { description: "Boundary condition: Quantity equals 50", price: 10, quantity: 50, expected: 375 }
    ];

    testCases.forEach(testCase => {
        const result = calculateDiscount(testCase.price, testCase.quantity);
        if (result === testCase.expected) {
            console.log(`PASS: ${testCase.description}`);
        } else {
            console.log(`FAIL: ${testCase.description} (expected: ${testCase.expected}, got: ${result})`);
        }
    });
}

runTests();