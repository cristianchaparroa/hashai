// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";

    error NoReportFound();
    error InvalidCategory();

contract HashReporter is Ownable {
    struct Report {
        uint256 count;
        uint256 category;
        bool exists;
    }

    mapping(address => Report) private reports;
    event ReportCreated(address indexed reportedAddress, uint256 count, uint256 category);

    constructor() Ownable(msg.sender) {}

    function createReport(address _reportedAddress, uint256 _category) external onlyOwner {
        Report storage report = reports[_reportedAddress];
        if (!report.exists) {
            reports[_reportedAddress] = Report({
                count: 1,
                category: _category,
                exists: true
            });
            emit ReportCreated(_reportedAddress, 1, _category);
            return;
        }

        report.count++;
        report.category = _category;
        emit ReportCreated(_reportedAddress, report.count, _category);
    }

    function getReportByAddress(address _reportedAddress) external view returns (uint256 count, uint256 category) {
        Report storage report = reports[_reportedAddress];
        if(!report.exists) revert NoReportFound();
        return (report.count, report.category);
    }

    function getCategoryString(uint256 _category) external pure returns (string memory) {
        if (_category == 0) return "Scam";
        else if (_category == 1) return "Phishing";
        else if (_category == 2) return "Malware";
        else if (_category == 3) return "Fraud";
        else if (_category == 4) return "Other";
        return "";
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
