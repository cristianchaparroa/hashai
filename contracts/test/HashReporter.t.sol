// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "forge-std/Test.sol";
import "../src/HashReporter.sol";

error OwnableUnauthorizedAccount(address account);

contract HashReporterTest is Test {
    HashReporter public reporter;
    address public owner;
    address public randomUser;
    address public reportedAddress;

    address public constant TEST_ADDRESS = address(0x1);

    event ReportCreated(address indexed reportedAddress, uint256 count, uint256 category);

    function setUp() public {
        owner = address(this);
        randomUser = address(0x1);
        reportedAddress = address(0x2);
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

    function testFuzzCreateReport(address _reportedAddress, uint256 category) public {
        vm.assume(category <= 4);
        reporter.createReport(_reportedAddress, category);

        (uint256 count, uint256 returnedCategory) = reporter.getReportByAddress(_reportedAddress);
        assertEq(count, 1);
        assertEq(returnedCategory, category);
    }

    function testOwnerCanCreateReport() public {
        // Owner creates a report
        reporter.createReport(reportedAddress, 0); // Category 0 = "Scam"

        // Verify the report was created
        (uint256 count, uint256 category) = reporter.getReportByAddress(reportedAddress);
        assertEq(count, 1, "Report count should be 1");
        assertEq(category, 0, "Category should be 0 (Scam)");
    }

    function testRandomUserCannotCreateReport() public {
        // Switch to random user context
        vm.startPrank(randomUser);

        // Expect the transaction to revert with OwnableUnauthorizedAccount error
        vm.expectRevert(
            abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, randomUser)
        );
        reporter.createReport(reportedAddress, 0);

        vm.stopPrank();
    }

    function testOwnerCanTransferOwnership() public {
        address newOwner = address(0x3);

        // Transfer ownership to new address
        reporter.transferOwnership(newOwner);

        // Verify new owner
        assertEq(reporter.owner(), newOwner, "Ownership should be transferred to new owner");
    }

    function testRandomUserCannotTransferOwnership() public {
        address newOwner = address(0x3);

        // Switch to random user context
        vm.startPrank(randomUser);

        // Expect the transaction to revert with OwnableUnauthorizedAccount error
        vm.expectRevert(
            abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, randomUser)
        );
        reporter.transferOwnership(newOwner);

        vm.stopPrank();
    }

    function testOwnershipTransferCompleteCycle() public {
        address newOwner = address(0x3);

        // Initial owner creates a report
        reporter.createReport(reportedAddress, 0);

        // Transfer ownership
        reporter.transferOwnership(newOwner);

        // Try to create report with old owner (should fail)
        vm.expectRevert(
            abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, address(this))
        );
        reporter.createReport(reportedAddress, 1);

        // Switch to new owner context
        vm.startPrank(newOwner);

        // New owner should be able to create report
        reporter.createReport(reportedAddress, 1);

        // Verify the report was updated by new owner
        (uint256 count, uint256 category) = reporter.getReportByAddress(reportedAddress);
        assertEq(count, 2, "Report count should be 2");
        assertEq(category, 1, "Category should be 1");

        vm.stopPrank();
    }
}
