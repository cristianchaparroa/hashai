// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract Blacklist {
    using ECDSA for bytes32;
    uint256 public constant MAX_BATCH_SIZE = 100;

    error NoReportFound();
    error InvalidSignature();
    error InvalidNonce();
    error InvalidLength();

    struct Report {
        uint256 count;
        uint256 category;
        uint256 date;
        string comments;
        string source;
        bool exists;
    }

    struct UserOperation {
        address sender;
        address reportedAddress;
        uint256 category;
        string comments;
        string source;
        uint256 date;
        uint256 nonce;
        bytes signature;
    }

    mapping(address => Report) private reports;
    mapping(address => uint256) public nonces;

    event Blacklisted(
        address indexed reportedAddress,
        uint256 count,
        uint256 category,
        uint256 date,
        string comments,
        string source
    );

    function executeBatchOperations(UserOperation[] calldata userOps) external {
        require(userOps.length <= MAX_BATCH_SIZE, "Batch too large");

        // Use memory for temporary storage
        address sender = userOps[0].sender;
        uint256 currentNonce = nonces[sender];

        // Validate all operations first
        for(uint256 i = 0; i < userOps.length;) {
            require(userOps[i].sender == sender, "Mixed senders");
            require(userOps[i].nonce == currentNonce + i, "Invalid nonce sequence");

            bytes32 hash = getOperationHash(userOps[i]);
            address signer = hash.recover(userOps[i].signature);
            require(signer == sender, "Invalid signature");

            unchecked { ++i; }
        }

        // Process operations
        for(uint256 i = 0; i < userOps.length;) {
            _processOperation(userOps[i]);
            unchecked { ++i; }
        }

        // Update nonce once
        nonces[sender] += userOps.length;
    }

    function _processOperation(UserOperation calldata userOp) internal {
        Report storage report = reports[userOp.reportedAddress];

        if (!report.exists) {
            report.count = 1;
            report.category = userOp.category;
            report.comments = userOp.comments;
            report.source = userOp.source;
            report.date = userOp.date;
            report.exists = true;
        } else {
            unchecked {
                report.count++;
            }
            report.category = userOp.category;
        }

        emit Blacklisted(
            userOp.reportedAddress,
            report.count,
            userOp.category,
            userOp.date,
            userOp.comments,
            userOp.source
        );
    }

    function getOperationHash(UserOperation calldata userOp) public view returns (bytes32) {
        return keccak256(abi.encodePacked(
            address(this),
            userOp.sender,
            userOp.reportedAddress,
            userOp.category,
            userOp.comments,
            userOp.source,
            userOp.date,
            userOp.nonce
        ));
    }

    function getReportByAddress(address reportedAddress) external view returns (uint256 count, uint256 category) {
        Report storage report = reports[reportedAddress];
        if(!report.exists) revert NoReportFound();
        return (report.count, report.category);
    }
}
