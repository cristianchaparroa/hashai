import {
  OwnershipTransferred as OwnershipTransferredEvent,
  ReportCreated as ReportCreatedEvent,
} from "../generated/HashReporter/HashReporter"
import { OwnershipTransferred, ReportCreated } from "../generated/schema"

export function handleOwnershipTransferred(
  event: OwnershipTransferredEvent,
): void {
  let entity = new OwnershipTransferred(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.previousOwner = event.params.previousOwner
  entity.newOwner = event.params.newOwner

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleReportCreated(event: ReportCreatedEvent): void {
  let entity = new ReportCreated(
    event.transaction.hash.concatI32(event.logIndex.toI32()),
  )
  entity.reportedAddress = event.params.reportedAddress
  entity.count = event.params.count
  entity.category = event.params.category
  entity.date = event.params.date
  entity.comments = event.params.comments
  entity.source = event.params.source

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
