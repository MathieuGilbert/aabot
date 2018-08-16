# AABot

## Developer Setup

 * Run local Ethereum node: Ganache
 * truffle compile; truffle test


#### Generating .go adapter from contract .sol

`abigen -sol ethereum/contracts/DexBot.sol -pkg adapters -out adapters/dexbot.go`


#### Node Packages

Install dependencies: `npm install`

Live reloading: `npm run watch`
