// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract HashReporter {
    
    enum Category {
        Scam,
        Phishing,
        Malware,
        Fraud,
        Other
    }

    
    struct Report {
        address reportedAddress; // The address being reported
        uint256 count;           
        Category category;       
    }

    Report[] public reports;

    mapping(address => uint256) public reportIndex;

    event ReportCreated(address indexed reportedAddress, uint256 count, Category category);

    function createReport(address _reportedAddress, Category _category) public {
        if (reportIndex[_reportedAddress] == 0 && reports.length == 0 || reports[reportIndex[_reportedAddress] - 1].reportedAddress != _reportedAddress) {
            reports.push(Report({
                reportedAddress: _reportedAddress,
                count: 1,
                category: _category
            }));

            reportIndex[_reportedAddress] = reports.length;

            emit ReportCreated(_reportedAddress, 1, _category);
        } else {
            uint256 index = reportIndex[_reportedAddress] - 1;
            reports[index].count++;
            emit ReportCreated(_reportedAddress, reports[index].count, _category);
        }
    }

    function getReport(uint256 index) public view returns (Report memory) {
        require(index < reports.length, "Index out of bounds");
        return reports[index];
    }

    function getReportByAddress(address _reportedAddress) public view returns (Report memory) {
        require(reportIndex[_reportedAddress] > 0, "No report found for this address");
        return reports[reportIndex[_reportedAddress] - 1];
    }

    function getCategoryString(Category _category) public pure returns (string memory) {
        if (_category == Category.Scam) return "Scam";
        if (_category == Category.Phishing) return "Phishing";
        if (_category == Category.Malware) return "Malware";
        if (_category == Category.Fraud) return "Fraud";
        if (_category == Category.Other) return "Other";
        return "";
    }
}
