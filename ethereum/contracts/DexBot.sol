pragma solidity ^0.4.24;

contract DexBot {
    uint public number;

    constructor(uint initialNumber) public {
        number = initialNumber;
    }

    function setNumber(uint newNumber) public {
        number = newNumber;
    }
}
