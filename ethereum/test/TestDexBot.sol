pragma solidity ^0.4.24;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/DexBot.sol";

contract TestDexBot {
    DexBot dexBot = DexBot(DeployedAddresses.DexBot());

    function testConstructor() public {
        Assert.equal(dexBot.number(), uint(42), "It assigned the number.");
    }
}
