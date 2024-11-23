// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test, console2} from "forge-std/Test.sol";
import {HashReporter} from "../src/HashReporter.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract HashReporterTest is Test {
    HashReporter public hashReporter;
    address public owner;
    address public user;
    address public reportedAddress;

    event ReportCreated(
        address indexed reportedAddress,
        uint256 count,
        uint256 category,
        uint256 date,
        string comments,
        string source
    );

    function setUp() public {
        owner = makeAddr("owner");
        user = makeAddr("user");
        reportedAddress = makeAddr("reported");

        vm.prank(owner);
        hashReporter = new HashReporter();
    }

    function test_InitialOwnership() public {
        assertEq(hashReporter.owner(), owner);
    }

    function test_CreateReport() public {
        vm.startPrank(owner);

        string memory comments = "Scam report";
        string memory source = "etherscan";
        uint256 date = block.timestamp;
        uint256 category = 1;

        vm.expectEmit(true, true, true, true);
        emit ReportCreated(
            reportedAddress,
            1,
            category,
            date,
            comments,
            source
        );

        hashReporter.createReport(
            reportedAddress,
            category,
            comments,
            source,
            date
        );

        (uint256 count, uint256 reportCategory) = hashReporter.getReportByAddress(reportedAddress);
        assertEq(count, 1);
        assertEq(reportCategory, category);

        vm.stopPrank();
    }

    function test_CreateMultipleReports() public {
        vm.startPrank(owner);

        // First report
        string memory firstComments = "First report";
        string memory firstSource = "etherscan";
        uint256 firstDate = block.timestamp;
        uint256 firstCategory = 1;

        hashReporter.createReport(
            reportedAddress,
            firstCategory,
            firstComments,
            firstSource,
            firstDate
        );

        // Second report
        string memory secondComments = "Second report";
        string memory secondSource = "blockscout";
        uint256 secondDate = block.timestamp;
        uint256 secondCategory = 2;

        vm.expectEmit(true, true, true, true);
        emit ReportCreated(
            reportedAddress,
            2,
            secondCategory,
            secondDate,
            secondComments,
            secondSource
        );

        hashReporter.createReport(
            reportedAddress,
            secondCategory,
            secondComments,
            secondSource,
            secondDate
        );

        (uint256 count, uint256 category) = hashReporter.getReportByAddress(reportedAddress);
        assertEq(count, 2);
        assertEq(category, secondCategory);

        vm.stopPrank();
    }

    function test_RevertWhenNonOwnerCreatesReport() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user)
        );

        hashReporter.createReport(
            reportedAddress,
            1,
            "Unauthorized report",
            "etherscan",
            block.timestamp
        );
    }

    function test_RevertWhenQueryingNonExistentReport() public {
        vm.expectRevert();
        hashReporter.getReportByAddress(makeAddr("nonexistent"));
    }

    function test_TransferOwnership() public {
        address newOwner = makeAddr("newOwner");

        vm.prank(owner);
        hashReporter.transferOwnership(newOwner);

        assertEq(hashReporter.owner(), newOwner);
    }

    function test_RevertWhenNonOwnerTransfersOwnership() public {
        address newOwner = makeAddr("newOwner");

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user)
        );

        hashReporter.transferOwnership(newOwner);
    }

    function test_RenounceOwnership() public {
        vm.prank(owner);
        hashReporter.renounceOwnership();

        assertEq(hashReporter.owner(), address(0));
    }

    function test_RevertWhenNonOwnerRenouncesOwnership() public {
        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user)
        );

        hashReporter.renounceOwnership();
    }

    function test_FuzzCreateReport(
        address _reportedAddress,
        uint256 _category,
        string memory _comments,
        string memory _source,
        uint256 _date
    ) public {
        vm.assume(_reportedAddress != address(0));
        vm.assume(_date > 0);

        vm.prank(owner);
        hashReporter.createReport(
            _reportedAddress,
            _category,
            _comments,
            _source,
            _date
        );

        (uint256 count, uint256 category) = hashReporter.getReportByAddress(_reportedAddress);
        assertEq(count, 1);
        assertEq(category, _category);
    }

    function test_CreateReportGasOptimized() public {
        vm.startPrank(owner);

        // Test first report
        uint256 gasBefore = gasleft();
        hashReporter.createReport(
            reportedAddress,
            1,
            "Test report",
            "etherscan",
            block.timestamp
        );
        uint256 gasAfter = gasleft();
        uint256 gasUsed = gasBefore - gasAfter;
        console2.log("Gas used for first report:", gasUsed);

        // Test subsequent report
        gasBefore = gasleft();
        hashReporter.createReport(
            reportedAddress,
            2,
            "Second report",
            "etherscan",
            block.timestamp
        );
        gasAfter = gasleft();
        gasUsed = gasBefore - gasAfter;
        console2.log("Gas used for subsequent report:", gasUsed);

        vm.stopPrank();
    }
}
