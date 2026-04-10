// Copy of the provided function
function calculateDiscount(price: number, quantity: number): number {
    if (quantity < 10) return price * quantity;
    else if (quantity < 50) return price * quantity * 0.9;
    else return price * quantity * 0.75;
}

// Test cases
const testCases = [
    {
        price: 10,
        quantity: 9,
        expected: 90,
        description: "No discount for quantity 9"
    },
    {
        price: 10,
        quantity: 10,
        expected: 90,
        description: "90% discount for quantity 10"
    },
    {
        price: 10,
        quantity: 49,
        expected: 441,
        description: "90% discount for quantity 49"
    },
    {
        price: 10,
        quantity: 50,
        expected: 375,
        description: "75% discount for quantity 50"
    },
    {
        price: 100,
        quantity: 5,
        expected: 500,
        description: "No discount for high price and quantity 5"
    }
];

// Run tests
const fs = require('fs');
const path = require('path');

// Write results to a file
const outputFile = path.join(__dirname, 'test_results.txt');
let fileContent = '';

for (const tc of testCases) {
    const result = calculateDiscount(tc.price, tc.quantity);
    const passed = result === tc.expected;
    
    let printMsg;
    if (passed) {
        printMsg = `PASS: ${tc.description}`;
    } else {
        printMsg = `FAIL: ${tc.description} (expected: ${tc.expected}, got: ${result})`;
    }
    
    fileContent += printMsg + '\n';
    console.log(printMsg);
}

// Write to file
fs.writeFileSync(outputFile, fileContent);