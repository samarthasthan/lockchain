// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract SecretStorage {

    struct SecretStruct {
        uint256 id;
        string secret;
        uint256 createdAt;
    }

    mapping(address => SecretStruct[]) private secrets;

    // Store a new secret
    function setSecret(string memory _secret) public {
        uint256 newId = secrets[msg.sender].length + 1;
        uint256 currentTimestamp = block.timestamp;

        secrets[msg.sender].push(SecretStruct({
            id: newId,
            secret: _secret,
            createdAt: currentTimestamp
        }));
    }

    // Get all secrets for a user
    function getSecrets() public view returns (SecretStruct[] memory) {
        return secrets[msg.sender];
    }

    // Get a specific secret for a user by index
    function getSecretByIndex(uint256 _index) public view returns (SecretStruct memory) {
        require(_index < secrets[msg.sender].length, "Index out of bounds");
        return secrets[msg.sender][_index];
    }
}
