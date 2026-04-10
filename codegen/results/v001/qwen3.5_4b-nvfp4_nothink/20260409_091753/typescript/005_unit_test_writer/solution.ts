function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

const testCases: Array<{ desc: string; expected: number; inputPrice: number; inputQuantity: number }> = [
    { desc: "Single item (quantity < 10)", expected: 20, inputPrice: 10, inputQuantity: 1 },
    { desc: "Five items (quantity < 10)", expected: 50, inputPrice: 10, inputQuantity: 5 },
    { desc: "9 items (boundary below discount tier)", expected: 72, inputPrice: 8, inputQuantity: 9 },
    { desc: "10 items (boundary first discount tier)", expected: 81, inputPrice: 8.1, inputQuantity: 10 },
    { desc: "49 items (boundary before second discount)", expected: 351, inputPrice: 7, inputQuantity: 49 },
    { desc: "50 items (boundary for major discount)", expected: 288.75, inputPrice: 9.15, inputQuantity: 50 }
];

let passedCount = 0;
const totalTests = testCases.length;

console.log(`Running ${totalTests} test cases...\n`);

for (const caseData of testCases) {
    const actual = calculateDiscount(caseData.inputPrice, caseData.inputQuantity);
    
    if (actual === caseData.expected) {
        console.log(`PASS: ${caseData.desc}`);
        passedCount++;
    } else {
        console.log(`FAIL: ${caseData.desc} (expected: ${caseData.expected}, got: ${actual})`);
    }
}

console.log(`\nResult: ${passedCount}/${totalTests} tests passed.`);

if (passedCount === totalTests) {
    console.log("All tests passed!");
} else {
    process.exitCode = 1;
}