# Smart Contract Documentation

This directory contains the smart contract for the project, including the contract source code, build artifacts, and migration scripts.

## Project Structure

- `build/`: Contains the compiled contract artifacts and ABI files.
  - `contracts/`: Compiled contract files in JSON format.
- `contracts/`: Contains the source code of the smart contract.
  - `secret.sol`: The Solidity smart contract source code.
- `migrations/`: Migration scripts for deploying the smart contract to a blockchain.
- `test/`: Contains test scripts for the smart contract.
- `truffle-config.js`: Configuration file for Truffle.

## Getting Started

### Prerequisites

1. **Node.js**: Make sure you have Node.js installed. You can download it from [nodejs.org](https://nodejs.org/).
2. **Truffle**: This project uses Truffle for managing smart contracts. Install Truffle globally by running:
   ```sh
   npm install -g truffle
   ```
3. **jq**: This project uses `jq` to process JSON files. Install `jq` using the package manager for your operating system:
   - **Ubuntu/Debian**:
     ```sh
     sudo apt-get install jq
     ```
   - **macOS**:
     ```sh
     brew install jq
     ```
   - **Windows**: You can download `jq` from the [jq releases page](https://github.com/stedolan/jq/releases) and follow the installation instructions.

### Installation

1. **Install Dependencies**: Navigate to the `smart-contract` directory and install the necessary dependencies:
   ```sh
   npm install
   ```

### Build the Contract

1. **Compile the Contract**: To compile the smart contract, use the following command:
   ```sh
   truffle compile
   ```
   This will generate the compiled contract artifacts in the `build/contracts/` directory.

### Migrate the Contract

1. **Deploy the Contract**: Ensure you have a blockchain environment set up (e.g., Ganache, an Ethereum testnet). Then deploy the contract using:
   ```sh
   truffle migrate
   ```
   This command runs the migration scripts located in the `migrations/` directory to deploy the contract.

### Test the Contract

1. **Run Tests**: To run the test scripts, use the following command:
   ```sh
   truffle test
   ```
   This will execute the test scripts located in the `test/` directory.

### Configuration

The `truffle-config.js` file contains the configuration settings for Truffle, including network configurations and compiler settings. Modify this file to fit your deployment environment.

## Additional Notes

- Ensure that you have set up a `.env` file with the necessary environment variables (e.g., private keys, Infura/Alchemy project IDs) if required by your configuration.
- For more details on Truffle, refer to the [Truffle documentation](https://www.trufflesuite.com/docs/truffle/overview).

## Troubleshooting

- **Compilation Errors**: Ensure that your Solidity code is correct and compatible with the Solidity compiler version specified in `truffle-config.js`.
- **Migration Issues**: Check the migration scripts and ensure that your blockchain environment is properly set up.

## License

This project is licensed under the MIT License. See the [LICENSE](../LICENSE) file for details.

```

### Key Points Added:

- **jq Installation Instructions**: Added details on how to install `jq` for different operating systems.

This additional information ensures that users have all the tools they need to work with the smart contract and handle JSON data appropriately. Adjust the installation commands based on the package managers and platforms relevant to your project.


ganache-cli --account_keys_path ./accounts.json --networkId 1337
truffle migrate --network development
```
