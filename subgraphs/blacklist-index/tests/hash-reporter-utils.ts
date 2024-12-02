import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import {
  OwnershipTransferred,
  ReportCreated
} from "../generated/HashReporter/HashReporter"

export function createOwnershipTransferredEvent(
  previousOwner: Address,
  newOwner: Address
): OwnershipTransferred {
  let ownershipTransferredEvent = changetype<OwnershipTransferred>(
    newMockEvent()
  )

  ownershipTransferredEvent.parameters = new Array()

  ownershipTransferredEvent.parameters.push(
    new ethereum.EventParam(
      "previousOwner",
      ethereum.Value.fromAddress(previousOwner)
    )
  )
  ownershipTransferredEvent.parameters.push(
    new ethereum.EventParam("newOwner", ethereum.Value.fromAddress(newOwner))
  )

  return ownershipTransferredEvent
}

export function createReportCreatedEvent(
  reportedAddress: Address,
  count: BigInt,
  category: BigInt,
  date: BigInt,
  comments: string,
  source: string
): ReportCreated {
  let reportCreatedEvent = changetype<ReportCreated>(newMockEvent())

  reportCreatedEvent.parameters = new Array()

  reportCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "reportedAddress",
      ethereum.Value.fromAddress(reportedAddress)
    )
  )
  reportCreatedEvent.parameters.push(
    new ethereum.EventParam("count", ethereum.Value.fromUnsignedBigInt(count))
  )
  reportCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "category",
      ethereum.Value.fromUnsignedBigInt(category)
    )
  )
  reportCreatedEvent.parameters.push(
    new ethereum.EventParam("date", ethereum.Value.fromUnsignedBigInt(date))
  )
  reportCreatedEvent.parameters.push(
    new ethereum.EventParam("comments", ethereum.Value.fromString(comments))
  )
  reportCreatedEvent.parameters.push(
    new ethereum.EventParam("source", ethereum.Value.fromString(source))
  )

  return reportCreatedEvent
}
