var DexBot = artifacts.require("DexBot");

module.exports = function(deployer) {
    deployer.deploy(DexBot, 42).then(() => console.log(DexBot.address));
};
