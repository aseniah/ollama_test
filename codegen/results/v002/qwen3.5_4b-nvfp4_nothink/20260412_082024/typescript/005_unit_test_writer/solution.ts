function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

function runTests() {
    const tests = [
        { desc: "Below threshold: quantity 9", expectedPrice: 10, expectedQuantity: 9, expectedResult: 90 },
        { desc: "At first boundary: quantity 10", expectedPrice: 10, expectedQuantity: 10, expectedResult: 90 },
        { desc: "Between boundaries: quantity 49", expectedPrice: 10, expectedQuantity: 49, expectedResult: 44.1 },
        { desc: "At second boundary: quantity 50", expectedPrice: 10, expectedQuantity: 50, expectedResult: 37.5 },
    ];

    tests.forEach((test) => {
        const result = calculateDiscount(test.expectedPrice, test.expectedQuantity);
        
        const isPass = Math.abs(result - test.expectedResult) < 1e-6; // floating point tolerance
        
        if (isPass) {
            console.log(`PASS: ${test.desc}`);
        } else {
            console.log(`FAIL: ${test.desc} (expected: ${test.expectedResult}, got: ${result})`);
        }
    });
}

runTests();