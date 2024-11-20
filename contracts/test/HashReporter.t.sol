// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "forge-std/Test.sol";
import "../src/HashReporter.sol";

contract HashReporterTest is Test {
    HashReporter public reporter;
    address public constant TEST_ADDRESS = address(0x1);

    event ReportCreated(address indexed reportedAddress, uint256 count, uint256 category);

    function setUp() public {
        reporter = new HashReporter();
    }

    function testCreateFirstReport() public {
        vm.expectEmit(true, false, false, true);
        emit ReportCreated(TEST_ADDRESS, 1, 0);

        reporter.createReport(TEST_ADDRESS, 0);

        (uint256 count, uint256 category) = reporter.getReportByAddress(TEST_ADDRESS);
        assertEq(count, 1);
        assertEq(category, 0);
    }

    function testIncrementExistingReport() public {
        reporter.createReport(TEST_ADDRESS, 1);

        vm.expectEmit(true, false, false, true);
        emit ReportCreated(TEST_ADDRESS, 2, 1);

        reporter.createReport(TEST_ADDRESS, 1);

        (uint256 count, uint256 category) = reporter.getReportByAddress(TEST_ADDRESS);
        assertEq(count, 2);
        assertEq(category, 1);
    }

    function testGetReportByAddressReverts() public {
        vm.expectRevert(NoReportFound.selector);
        reporter.getReportByAddress(TEST_ADDRESS);
    }

    function testGetCategoryString() public {
        assertEq(reporter.getCategoryString(0), "Scam");
        assertEq(reporter.getCategoryString(1), "Phishing");
        assertEq(reporter.getCategoryString(2), "Malware");
        assertEq(reporter.getCategoryString(3), "Fraud");
        assertEq(reporter.getCategoryString(4), "Other");
        assertEq(reporter.getCategoryString(5), "");
    }

    function testFuzzCreateReport(address reportedAddress, uint256 category) public {
        vm.assume(category <= 4);
        reporter.createReport(reportedAddress, category);

        (uint256 count, uint256 returnedCategory) = reporter.getReportByAddress(reportedAddress);
        assertEq(count, 1);
        assertEq(returnedCategory, category);
    }
}
