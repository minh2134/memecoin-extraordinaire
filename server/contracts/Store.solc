// server/contracts/Store.sol
pragma solidity ^0.8.0;
contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory _version) {  // Added "memory" keyword for string parameter
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
  
  function getItem(bytes32 key) external view returns (bytes32) {
    return items[key];
  }
}