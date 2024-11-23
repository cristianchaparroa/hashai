// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";

contract HashReporter is Ownable {
    error NoReportFound();
    error InvalidCategory();

    struct Report {
        uint256 count;      // number of cases reported for the same wallet
        uint256 category;   // category of the report, i.e: scam, phishing, etc..
        uint256 date;       // Date when was reported the address
        string comments;    // Comments related to the report.
        string source;      // Data source, i.e: etherescan, blockscout, an user;
        bool exists;
    }

    mapping(address => Report) private reports;
    event ReportCreated(
        address indexed reportedAddress,
        uint256 count,
        uint256 category,
        uint256 date,
        string comments,
        string source
    );

    constructor() Ownable(msg.sender) {}

    function createReport(
        address _reportedAddress,
        uint256 _category,
        string calldata _comments,
        string calldata _source,
        uint256 _date
    ) external onlyOwner {
        Report storage report = reports[_reportedAddress];

        if (!report.exists) {
            report.count = 1;
            report.category = _category;
            report.comments = _comments;
            report.source = _source;
            report.date = _date;
            report.exists = true;

            emit ReportCreated(_reportedAddress, report.count, _category, _date, _comments, _source);
            return;
        }

        unchecked {
            report.count++;  // Use unchecked since overflow is highly unlikely
        }
        report.category = _category;
        emit ReportCreated(_reportedAddress, report.count, _category, _date, _comments, _source);
    }

    function getReportByAddress(address _reportedAddress) external view returns (uint256 count, uint256 category) {
        Report storage report = reports[_reportedAddress];
        if(!report.exists) revert NoReportFound();
        return (report.count, report.category);
    }

    // Function to transfer ownership
    function transferOwnership(address newOwner) public override onlyOwner {
        super.transferOwnership(newOwner);
    }

    // Function to renounce ownership
    function renounceOwnership() public override onlyOwner {
        super.renounceOwnership();
    }
}
